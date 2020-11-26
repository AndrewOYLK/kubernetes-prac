package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

func main() {
	e := gin.Default()
	e.GET("/ws/pod/:namespace/:pod/:container", wsPodHandler)
	e.GET("/ws/node/:ip", wsHostHandler)
	e.GET("/ws/local", wsLocalHandler)
	e.Run(":8080")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsPodMsg 搜索
type WsPodMsg struct {
	NameSpace string `uri:"namespace"`
	Pod       string `uri:"pod"`
	Container string `uri:"container"`
}

func wsPodHandler(ctx *gin.Context) {
	var sm = new(WsPodMsg)
	ctx.ShouldBindUri(sm)

	// 从http协议升级到websocket协议
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		log.Println("upgrader.Upgrade")
		return
	}
	defer conn.Close()
}

// HostInfo 主机信息
type HostInfo struct {
	IP       string
	User     string
	Password string
	Port     string
}

// 主机信息
var hosts []*HostInfo = []*HostInfo{&HostInfo{
	IP:       "10.1.0.13",
	User:     "root",
	Password: "123qwe",
	Port:     "222",
}, &HostInfo{
	IP:       "10.200.10.77",
	User:     "root",
	Password: "123qwe",
	Port:     "222",
}}

type wsBufferWriter struct {
	buffer bytes.Buffer
	mu     sync.Mutex // 锁
}

// 缓存写
func (w *wsBufferWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}

// SshConn ssh连接
type SshConn struct {
	StdinPipe   io.WriteCloser
	ComboOutput *wsBufferWriter
	Session     *ssh.Session
}

func wsHostHandler(ctx *gin.Context) {
	// 获取主机信息
	var host = new(HostInfo)
	hostIP := ctx.Param("ip")
	for _, v := range hosts {
		if hostIP == v.IP {
			host = v
		}
	}

	// 定义ssh客户端
	config := &ssh.ClientConfig{
		Timeout:         5 * time.Second,
		User:            host.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password(host.Password)},
	}
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host.IP, host.Port), config)
	if err != nil {
		panic(err)
	}

	// 新建session会话
	sshSession, _ := sshClient.NewSession()

	stdinP, _ := sshSession.StdinPipe()
	comboWriter := new(wsBufferWriter)
	sshSession.Stdout = comboWriter
	sshSession.Stderr = comboWriter
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	cols := 120
	rows := 32
	// 新建一个伪终端
	if err := sshSession.RequestPty("xterm", rows, cols, modes); err != nil {
		panic(err)
	}

	// 启动
	if err := sshSession.Shell(); err != nil {
		panic(err)
	}

	ssConn := &SshConn{StdinPipe: stdinP, ComboOutput: comboWriter, Session: sshSession}

	// websocket
	wsConn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer wsConn.Close()

	quitChan := make(chan bool, 3)
	var logBuff = new(bytes.Buffer)

	// most messages are ssh output, not webSocket input
	go ssConn.ReceiveWsMsg(wsConn, logBuff, quitChan)
	go ssConn.SendComboOutput(wsConn, quitChan)
	go ssConn.SessionWait(quitChan)
	<-quitChan
}

func wsLocalHandler(ctx *gin.Context) {

}

const (
	wsMsgCmd    = "cmd"
	wsMsgResize = "resize"
)

type wsMsg struct {
	Type string `json:"type"`
	Cmd  string `json:"cmd"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
}

//ReceiveWsMsg  receive websocket msg do some handling then write into ssh.session.stdin
func (ssConn *SshConn) ReceiveWsMsg(wsConn *websocket.Conn, logBuff *bytes.Buffer, exitCh chan bool) {
	//tells other go routine quit
	defer setQuit(exitCh)
	for {
		select {
		case <-exitCh:
			return
		default:
			//read websocket msg
			_, wsData, err := wsConn.ReadMessage()
			if err != nil {
				logrus.WithError(err).Error("reading webSocket message failed")
				return
			}
			//unmashal bytes into struct
			msgObj := wsMsg{}
			if err := json.Unmarshal(wsData, &msgObj); err != nil {
				logrus.WithError(err).WithField("wsData", string(wsData)).Error("unmarshal websocket message failed")
			}
			switch msgObj.Type {
			case wsMsgResize:
				//handle xterm.js size change
				if msgObj.Cols > 0 && msgObj.Rows > 0 {
					if err := ssConn.Session.WindowChange(msgObj.Rows, msgObj.Cols); err != nil {
						logrus.WithError(err).Error("ssh pty change windows size failed")
					}
				}
			case wsMsgCmd:
				//handle xterm.js stdin
				decodeBytes, err := base64.StdEncoding.DecodeString(msgObj.Cmd)
				if err != nil {
					logrus.WithError(err).Error("websock cmd string base64 decoding failed")
				}
				if _, err := ssConn.StdinPipe.Write(decodeBytes); err != nil {
					logrus.WithError(err).Error("ws cmd bytes write to ssh.stdin pipe failed")
				}
				//write input cmd to log buffer
				if _, err := logBuff.Write(decodeBytes); err != nil {
					logrus.WithError(err).Error("write received cmd into log buffer failed")
				}
			}
		}
	}
}
func (ssConn *SshConn) SendComboOutput(wsConn *websocket.Conn, exitCh chan bool) {
	//tells other go routine quit
	defer setQuit(exitCh)

	//every 120ms write combine output bytes into websocket response
	tick := time.NewTicker(time.Millisecond * time.Duration(120))
	//for range time.Tick(120 * time.Millisecond){}
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			//write combine output bytes into websocket response
			if err := flushComboOutput(ssConn.ComboOutput, wsConn); err != nil {
				logrus.WithError(err).Error("ssh sending combo output to webSocket failed")
				return
			}
		case <-exitCh:
			return
		}
	}
}

func (ssConn *SshConn) SessionWait(quitChan chan bool) {
	if err := ssConn.Session.Wait(); err != nil {
		logrus.WithError(err).Error("ssh session wait failed")
		setQuit(quitChan)
	}
}

func setQuit(ch chan bool) {
	ch <- true
}

//flushComboOutput flush ssh.session combine output into websocket response
func flushComboOutput(w *wsBufferWriter, wsConn *websocket.Conn) error {
	if w.buffer.Len() != 0 {
		err := wsConn.WriteMessage(websocket.TextMessage, w.buffer.Bytes())
		if err != nil {
			return err
		}
		w.buffer.Reset()
	}
	return nil
}
