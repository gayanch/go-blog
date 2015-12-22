package main

import (
	"net/http"
	"github.com/gayanch/go-blog/handler"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/a/", handler.Article)
	http.HandleFunc("/p/", handler.Page)
	http.HandleFunc("/login", handler.Login)
	
	http.ListenAndServe(":8080", nil)
}
