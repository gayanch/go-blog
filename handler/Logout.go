package handler

import (
	"net/http"
	"time"
	"fmt"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout")
	expires := time.Now()
	cookie := &http.Cookie{Name: "type", Expires: expires, Value: "none"}
	http.SetCookie(w, cookie)
//	if cookie, err := r.Cookie("type"); err != nil {
////		cookie.MaxAge = 0
////		cookie.Expires = time.Now()
////		cookie.Value = "none"
////		r.AddCookie(cookie)
//		expires := time.Now()
//		cookie := &http.Cookie{Name: "type", Expires: expires, Value: "none"}
//		http.SetCookie(w, cookie)
//	}
	http.Redirect(w, r, "/", 302)
}
