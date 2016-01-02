package handler

import (
    "net/http"
    "html/template"
    "fmt"

    "github.com/gayanch/go-blog/db/orm"
)

func Edit(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, ": /edit")

    editor := orm.Editor()
    if s, ok := sm.Get(r); ok && s.UserType == "admin" {
        t, _ := template.ParseFiles("template/edit.html")
        t.ExecuteTemplate(w, t.Name(), editor)
    } else {
        http.Redirect(w, r, "/", 302)
    }

}
