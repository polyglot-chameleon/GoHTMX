package main

import (
	"fmt"
	"net/http"

	"GoHTMX/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.StaticFile("/htmx.min.js", "./dist/htmx.min.js")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/data", func(c *gin.Context) {
		var posts = controller.Controller.All()
		var dataHtml = ""

		for i := 0; i < len(posts); i++ {
			dataHtml += fmt.Sprintf("<li>%s</li>", posts[i].Title)
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(dataHtml))
	})

	r.GET("/form/create", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<form hx-post="/data" hx-target="this" hx-swap="outerHTML"><div><label>Title</label><input type="text" name="title"></div><div><label>Body</label><input type="text" name="body"></div><button class="btn">Submit</button><button class="btn" hx-get="/contact/1">Cancel</button></form>`))
	})

	r.POST("/data", func(c *gin.Context) {
		controller.Controller.Create(controller.PostResource{Title: c.PostForm("title"), Body: c.PostForm("body")})

	})

	r.Run()
}
