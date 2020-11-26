package stream

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// http 使用websocket协议处理
var websockUpgrader = websocket.Upgrader{
	HandshakeTimeout: 60 * time.Second,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 设置websocket信息
type WebSocketMessage struct {
	MessageType int
	Date        []byte
}

// 设置websocket连接
type WebsocketConn struct {
	// 底层的Socket
	Wsocket *websocket.Conn
	// 读取的队列
	inChan chan *WebSocketMessage
	// 发送的队列
	outChan chan *WebSocketMessage
	// 避免重复关闭锁
	mutex sync.Mutex
	// 是否已经关闭
	isClose bool
	//关闭通知
	closeChan chan byte
}

// 执行读取的携程
func (wsConn *WebsocketConn) webSocketReadLoop() {
	var (
		messageType int
		data        []byte
		msg         *WebSocketMessage
		err         error
	)
	for {
		// 读取一个message
		if messageType, data, err = wsConn.Wsocket.ReadMessage(); err != nil {
			goto ERR
		}
		// 构造消息对象
		msg = &WebSocketMessage{
			MessageType: messageType,
			Date:        data,
		}
		// 存入请求的队列
		select {
		case wsConn.inChan <- msg:
		case <-wsConn.closeChan:
			goto CLOSE
		}
	}
ERR:
	wsConn.WsClose()
CLOSE:
}

// 执行发送消息携程
func (wsConn *WebsocketConn) websocketWriteLoop() {
	var (
		msg *WebSocketMessage
		err error
	)
	for {
		select {
		// 从管道取出一个信息
		case msg = <-wsConn.outChan:
			// 复写给websocket
			if err = wsConn.Wsocket.WriteMessage(msg.MessageType, msg.Date); err != nil {
				goto ERR
			}
		case <-wsConn.closeChan:
			goto CLOSE
		}
	}
ERR:
	wsConn.WsClose()
CLOSE:
}

// 并发安全API

// 初始化Websocket信息
func InitWebsocket(w http.ResponseWriter, r *http.Request) (wsConn *WebsocketConn, err error) {
	var (
		wsSocket *websocket.Conn
	)
	// 回答客户端告知升级连接为websocket
	if wsSocket, err = websockUpgrader.Upgrade(w, r, nil); err != nil {
		return nil, err
	}
	// 构造websocket连接对象
	wsConn = &WebsocketConn{
		Wsocket:   wsSocket,
		inChan:    make(chan *WebSocketMessage, 1000),
		outChan:   make(chan *WebSocketMessage, 1000),
		isClose:   false,
		closeChan: make(chan byte),
	}

	// 开启读携程
	go wsConn.webSocketReadLoop()
	// 开启写携程
	go wsConn.websocketWriteLoop()
	return
}

// 发送消息
func (wsConn *WebsocketConn) WsSend(messageType int, data []byte) (err error) {
	select {
	case wsConn.outChan <- &WebSocketMessage{
		MessageType: messageType,
		Date:        data,
	}:
	case <-wsConn.closeChan:
		err = errors.New("Websocket is Closed")
	}
	return nil
}

// 读取消息
func (wsConn *WebsocketConn) WsRead() (msg *WebSocketMessage, err error) {
	select {
	case msg = <-wsConn.outChan:
		fmt.Println("读消息", msg)
		return
	case <-wsConn.closeChan:
		err = errors.New("WebSocket is Closed")
	}
	return
}

// 关闭连接
func (wsConn *WebsocketConn) WsClose() {
	// 关闭websocket连接
	wsConn.Wsocket.Close()
	// 加锁处理
	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	// 判断是否已经关闭了closechan
	if !wsConn.isClose {
		wsConn.isClose = true
		close(wsConn.closeChan)
	}
}
