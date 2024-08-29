package controller

import (
	"GoHTMX/util"
	"database/sql"
	"fmt"
	"os"

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
}

func (mc *controller) Connect() error {
	var err error
	Controller.db, err = sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_URL"))
	util.Check(err)
	return err
}

func (mc *controller) Create(newPost PostResource) (sql.Result, error) {
	result, err := mc.db.Exec(fmt.Sprintf("INSERT INTO posts(title, body) VALUES ('%s', '%s')", newPost.Title, newPost.Body))
	return result, err
}

func (mc *controller) Read(postID int64) (PostResource, error) {
	rows, err := mc.db.Query(fmt.Sprintf("SELECT title, body FROM posts WHERE id = %v", postID))
	util.Check(err)

	defer rows.Close()

	post := PostResource{Title: "", Body: ""}

	for rows.Next() {
		rows.Scan(&post.Title, &post.Body)
	}

	return post, err
}

func (mc *controller) All() ([]PostResource, error) {
	rows, err := mc.db.Query("SELECT title, body FROM posts;")
	util.Check(err)

	defer rows.Close()

	var storedPosts []PostResource
	post := PostResource{Title: "", Body: ""}

	for rows.Next() {
		rows.Scan(&post.Title, &post.Body)
		storedPosts = append(storedPosts, post)
	}

	return storedPosts, err
}

func (mc *controller) Delete(postId int64) (sql.Result, error) {
	result, err := mc.db.Exec(fmt.Sprintf("DELETE FROM posts WHERE id = %v;", postId))
	util.Check(err)
	return result, err
}
