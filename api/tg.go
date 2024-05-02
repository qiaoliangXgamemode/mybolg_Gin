package api

import (
	"Myblog/Parser"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

var apiUrl = "https://api.telegram.org"

func init() {
	router = gin.Default()
	router.Any("/", Parser.ListResponseJSON)
}

func Listen(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
