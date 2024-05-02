package Parser

import (
	"io/ioutil"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Article struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	PagePotho string `json:"pagePotho"`
	Time      string `json:"time"`
	Intro     string `json:"intro"`
	Tag       []string
}

func ListResponseJSON(c *gin.Context) {
	JsonFile, err := os.Open("/src/article.json")
	defer JsonFile.Close()
	Bytes, _ := ioutil.ReadAll(JsonFile)
	if err != nil {
		slog.Error("Failed: %s", err)
	}

	// jsonData, err := json.MarshalIndent(string(Bytes), "", "    ")
	// if err != nil {
	// 	slog.Error("Failed: %s", err)
	// }
	c.JSON(http.StatusOK, string(Bytes))
	return
}
