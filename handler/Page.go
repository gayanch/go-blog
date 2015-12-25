package handler

import (
  "net/http"
  "text/template"
  "fmt"

  "github.com/gayanch/go-blog/data"
)

func Page(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, r.URL)

    t, _ := template.ParseFiles("template/page.html")
    t.Execute(w, data.Page("1"))
}
