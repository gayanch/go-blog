package handler

import (
  "net/http"
  "html/template"
  "fmt"

  "github.com/gayanch/go-blog/db/orm"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, r.URL)

    data := orm.Home()
    if s, ok := sm.Get(r); ok && s.UserType == "admin" {
        data.Admin = true
    } else {
        data.Admin = false
    }
    t, _ := template.ParseFiles("template/index.html")
    t.ExecuteTemplate(w, t.Name(), data)
}
