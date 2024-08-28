package controller

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type PostResource struct {
	Title string
	Body  string
}

type controller struct {
	db *sql.DB
}

var Controller *controller

func init() {
	Controller = &controller{}
	Controller.db, _ = sql.Open("sqlite3", "db/gohtmx.db")
}

func (mc *controller) Create(newPost PostResource) {
	_, err := mc.db.Exec(fmt.Sprintf("INSERT INTO posts(title, body) VALUES ('%s', '%s')", newPost.Title, newPost.Body))
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *controller) All() []PostResource {
	rows, _ := mc.db.Query("SELECT title, body FROM posts;")
	defer rows.Close()

	var storedPosts []PostResource
	post := PostResource{Title: "", Body: ""}

	for rows.Next() {
		rows.Scan(&post.Title, &post.Body)
		storedPosts = append(storedPosts, post)
	}

	return storedPosts
}
