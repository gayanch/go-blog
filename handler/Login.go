package handler

import (
  "net/http"
  "html/template"
  "fmt"

  "github.com/gayanch/go-blog/data"
)

func Login(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.Method, ": Login")
  if r.Method == "GET" {
      t, _ := template.ParseFiles("template/login.html")
      t.ExecuteTemplate(w, t.Name(), data.Login())
  } else {
      r.ParseForm()
      username := r.Form["username"][0]
      password := r.Form["password"][0]

      if userType, ok := data.Auth(username, password); ok {
          //save session
          sm.Set(w, username, userType)

          http.Redirect(w, r, "/", 302)
      } else {
          fmt.Fprintf(w, "Invalid login")
      }
  }

}
