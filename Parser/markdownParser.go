package Parser

import (
	"io/ioutil"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func ParserMarkdown(c *gin.Context) {

	articlePath := c.Param("title")
	JsonFile, err := os.Open("./src/md/" + articlePath + ".md")

	Bytes, err := ioutil.ReadAll(JsonFile)

	if err != nil {
		slog.Error("Failed: %s", err)
	}
	// slog.Info("Markdown", Bytes)
	markdownText := string(mdToHTML(Bytes))
	defer JsonFile.Close()
	c.String(http.StatusOK, markdownText)
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
