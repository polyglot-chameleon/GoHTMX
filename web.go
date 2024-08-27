package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var data = [4]string{"Article 1", "Article 2", "Article 3", "Article 4"}

	r := gin.Default()
	r.LoadHTMLFiles("index.htmx")
	r.StaticFile("/htmx.min.js", "./dist/htmx.min.js")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.htmx", gin.H{})
	})

	r.GET("/data", func(c *gin.Context) {
		var dataHtml = ""

		for i := 0; i < len(data); i++ {
			dataHtml += fmt.Sprintf("<li>%s</li>", data[i])
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(dataHtml))
	})

	r.Run()
}
