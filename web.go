package main

import (
	"GoHTMX/controller"
	"fmt"
	"io"
	"net/http"
)

func serveFile(fpath string) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) { http.ServeFile(rw, req, fpath) }
}

func handleData(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")

		var posts = controller.Controller.All()
		var dataHtml = ""

		for i := 0; i < len(posts); i++ {
			dataHtml += fmt.Sprintf("<li>%s</li>", posts[i].Title)
		}

		io.WriteString(rw, dataHtml)

	} else if req.Method == "POST" {
		req.ParseForm()

		controller.Controller.Create(controller.PostResource{Title: req.Form.Get("title"), Body: req.Form.Get("body")})
	}
}

func main() {
	http.HandleFunc("/", serveFile("index.html"))
	http.HandleFunc("/data", handleData)
	http.HandleFunc("/form/create", serveFile("htmx/form.html"))

	http.HandleFunc("/htmx.min.js", serveFile("dist/htmx.min.js"))
	http.ListenAndServe(":8080", nil)
}
