package main

import (
  "net/http"

  "github.com/gayanch/go-blog/handler"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", handler.Home)
    http.HandleFunc("/page", handler.Page)
    http.HandleFunc("/login", handler.Login)
    http.HandleFunc("/logout", handler.Logout)
    http.HandleFunc("/new", handler.NewPost)
    http.HandleFunc("/view/", handler.ViewPost)

    http.ListenAndServe(":8080", nil)
}
