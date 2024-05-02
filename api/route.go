package api

import (
	"Myblog/Parser"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

}

var (
	r *gin.Engine
)

// var ListenPort = flag.Int("port", 8080, "Listen Port")

func init() {
	// TODO
	r := gin.Default()
	// r.GET("/", func(context *gin.Context) {
	// 	context.String(http.StatusOK, "hellow")
	// })
	r.Any("/api/route", Parser.ListResponseJSON)
	// r.GET("/api/route", Parser.ListResponseJSON)
	// r.GET("/", Parser.ListResponseJSON)
	// r.GET("/ArticleList", Parser.ListResponseJSON)

	// r.GET("/articleContext/:title", Parser.ParserMarkdown)
	// flag.Parse()

	// slog.Info("HTTP/TCP", "Listen", *ListenPort)

}

func Handler(w http.ResponseWriter, rw *http.Request) {
	r.ServeHTTP(w, rw)
}
