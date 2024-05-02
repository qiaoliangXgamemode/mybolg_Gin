package main

import (
	"Myblog/Parser"
	"flag"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func init() {

}

var ListenPort = flag.Int("port", 8080, "Listen Port")

func main() {
	// TODO
	r := gin.Default()
	// r.GET("/", func(context *gin.Context) {
	// 	context.String(http.StatusOK, "hellow")
	// })
	r.GET("/ArticleList", Parser.ListResponseJSON)

	r.GET("/articleContext/:title", Parser.ParserMarkdown)
	flag.Parse()

	slog.Info("HTTP/TCP", "Listen", *ListenPort)
	Handler(r)
}

func Handler(r *gin.Engine) {
	r.Run(fmt.Sprintf(":%v", *ListenPort))
}
