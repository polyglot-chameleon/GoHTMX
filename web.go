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

func getData(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	var posts = controller.Controller.All()
	var dataHtml = ""

	for i := 0; i < len(posts); i++ {
		dataHtml += fmt.Sprintf("<li>%s</li>", posts[i].Title)
	}

	io.WriteString(rw, dataHtml)
}

func postData(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	controller.Controller.Create(controller.PostResource{Title: req.Form.Get("title"), Body: req.Form.Get("body")})
}

func addRoutes() {
	http.HandleFunc("GET /", serveFile("index.html"))
	http.HandleFunc("GET /form/create", serveFile("htmx/form.html"))

	http.HandleFunc("GET /data", getData)
	http.HandleFunc("POST /data", postData)

	http.HandleFunc("GET /htmx.min.js", serveFile("dist/htmx.min.js"))
}

func main() {
	addRoutes()
	http.ListenAndServe(":8080", nil)
}
