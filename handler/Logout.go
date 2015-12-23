package handler

import (
	"net/http"
	"time"
	"fmt"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("type")	
	fmt.Println(cookie, ": ", err)
	//if (cookie != nil) {
		cookie.Expires = time.Now()
		cookie.Value = "none"
		http.SetCookie(w, cookie)
	//}
	http.Redirect(w, r, "/", 301)
}
