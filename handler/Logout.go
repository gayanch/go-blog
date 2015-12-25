package handler

import (
    "net/http"
    "fmt"
)

func Logout(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, ": Logout")
    sm.Delete(r)
    http.Redirect(w, r, "/", 302)
}
