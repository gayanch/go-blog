package handler

import (
  "net/http"
  "html/template"
  "fmt"

  "github.com/gayanch/go-blog/data"
)

func NewArticle(w http.ResponseWriter, r *http.Request) {
  fmt.Println("NewArticle: ")
  if s, ok := sm.Get(r); ok && s.UserType == "admin" {
    t, _ := template.ParseFiles("template/new.html")
    t.ExecuteTemplate(w, t.Name(), data.NewArticleData())
  } else {
    http.Redirect(w, r, "/", 301)
  }

}
