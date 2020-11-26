package httpLog

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"

	"net/http"
)

var DeFaultPreFix = "http-log-prefix"

// HttpLog 一个打印请求信息的插件
type HttpLog struct {
	Name   string
	Prefix string
}

//定义一些参数，可以通过启动micro 的时候传参
func (l *HttpLog) Flags() []cli.Flag {
	return []cli.Flag{cli.StringFlag{
		Name:   "httpLog",
		Usage:  "需要从命令行传过来的参数",
		EnvVar: "HTTPLOG",
	}}
}

func (l *HttpLog) Commands() []cli.Command {
	return nil
}

//处理程序 会在每次请求的时候调用
func (l *HttpLog) Handler() plugin.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("== %s: %s ==", l.Prefix, r.URL.Path)
			handler.ServeHTTP(w, r)
		})
	}
}

//初始化数据,程序启动的时候会调用这个方法，可以在这里初始化一些参数
func (l *HttpLog) Init(ctx *cli.Context) error {
	prefix := ctx.String("httplog")
	if prefix == "" {
		l.Prefix = DeFaultPreFix
	} else {
		l.Prefix = prefix
	}
	return nil
}
func (l *HttpLog) String() string {
	return l.Name
}

//调用这个方法 new 插件，也可以在启动的时候这样子写
func NewPlugin() plugin.Plugin {
	return &HttpLog{
		Name: "httpLog",
	}
}
