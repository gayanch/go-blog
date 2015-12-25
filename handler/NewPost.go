package handler

import (
  "net/http"
  "html/template"
  "fmt"

  "github.com/gayanch/go-blog/data"
)

func NewPost(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, ": new")

    if s, ok := sm.Get(r); ok && s.UserType == "admin" {
        if r.Method == "GET" {
            t, _ := template.ParseFiles("template/new.html")
            t.ExecuteTemplate(w, t.Name(), data.NewPost())
        } else {
            r.ParseForm()
            title := r.Form["title"][0]
            body := r.Form["body"][0]
            data.SavePost(title, body)
            http.Redirect(w, r, "/", 302)
        }
    } else {
        http.Redirect(w, r, "/", 302)
    }

}
