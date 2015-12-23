package main

import (
	"net/http"
	"github.com/gayanch/go-blog/handler"
)

func main() {	
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/login", handler.Login)	
	http.HandleFunc("/a/", handler.Article)
	http.HandleFunc("/p/", handler.Page)
	http.HandleFunc("/", handler.Index)
		
	http.ListenAndServe(":8080", nil)
}
