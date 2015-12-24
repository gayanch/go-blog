package main

import (
	"net/http"
	"github.com/gayanch/go-blog/handler"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/a/", handler.Article)
	http.HandleFunc("/p/", handler.Page)
	http.HandleFunc("/new", handler.NewArticle)
	http.HandleFunc("/save", handler.SaveArticle)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)

	http.ListenAndServe(":8080", nil)
}
