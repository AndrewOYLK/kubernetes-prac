package main

import (
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"

	"test/httpLog"
)

func init() {
	plugin.Register(httpLog.NewPlugin())

	// plugin.Register(plugin.NewPlugin(
	// 	plugin.WithName("httpLog"),
	// 	plugin.WithHandler(func(handler http.Handler) http.Handler {
	// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 			log.Printf("%s: %s", "myprofix", r.URL.Path)
	// 			handler.ServeHTTP(w, r)
	// 		})
	// 	}),
	// 	plugin.WithFlag(
	// 		cli.StringFlag{
	// 			Name:   "httpLog",
	// 			Usage:  "需要从命令行传过来的参数",
	// 			EnvVar: "HTTP_LOG",
	// 		}),
	// 	plugin.WithInit(func(context *cli.Context) error {
	// 		// TODO
	// 		return nil
	// 	}),
	// ))
}

func main() {
	cmd.Init()
}
