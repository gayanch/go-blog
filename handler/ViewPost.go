package handler

import (
    "net/http"
    "text/template"
    "fmt"

    "github.com/gayanch/go-blog/db/orm"
)

func ViewPost(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, r.URL.Path)

    pid := r.URL.Path[len("/view/"):]

    if post, ok := orm.Post(pid); ok {
        t, _ := template.ParseFiles("template/view.html")
        t.Execute(w, post)
    } else {
        http.NotFound(w, r)
    }

}
