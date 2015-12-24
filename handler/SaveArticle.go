package handler

import (
  //"fmt"
  "net/http"

  "github.com/gayanch/go-blog/data"
)

func SaveArticle(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  title := r.Form["title"][0]
  body := r.Form["body"][0]

  data.SaveAricle(title, body)
  http.Redirect(w, r, "/", 301)
}
