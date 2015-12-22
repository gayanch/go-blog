package handler

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/gayanch/go-blog/data"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	
	expire := time.Now().Add(time.Hour)
	cookie := http.Cookie{Name: "username", Value: username, Expires: expire}
	http.SetCookie(w, &cookie)
	
	if (data.Auth(username, password)) {
		fmt.Fprintf(w, "%v", true)
	} else {
		fmt.Fprintf(w, "%v", false)
	}
}
