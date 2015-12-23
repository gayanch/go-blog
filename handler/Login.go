package handler

import (
	//"fmt"
	"net/http"
	"time"
	
	"github.com/gayanch/go-blog/data"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	
	ok, level := data.Auth(username, password)
	//fmt.Println("login ", username, password)
	if (ok) {
		expire := time.Now().Add(time.Hour)
		cookie := http.Cookie{Name: "type", Value: level, Expires: expire}		
		http.SetCookie(w, &cookie)		
	}
	http.Redirect(w, r, "/", 301)
}
