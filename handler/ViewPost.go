package handler

import (
    "net/http"
    "text/template"
    "fmt"

    "github.com/gayanch/go-blog/data"
)

func ViewPost(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, r.URL.Path)

    pid := r.URL.Path[len("/view/"):]

    t, _ := template.ParseFiles("template/view.html")
    t.Execute(w, data.ViewPost(pid))
}
