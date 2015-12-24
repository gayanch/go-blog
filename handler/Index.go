package handler

import (
	"net/http"
	"html/template"
	"fmt"

	"github.com/gayanch/go-blog/data"
	"github.com/gayanch/go-blog/session"
)

var sm session.SessionManager

func init() {
	sm = session.NewSessionManager()
}

func Index(w http.ResponseWriter, r *http.Request) {
//	value := "none"
//	if cookie, err := r.Cookie("type"); err == nil && cookie != nil {
//		fmt.Println("Index: ", cookie.Value)
//		value = cookie.Value
//	}
		fmt.Println("Index: ")
		var value string = "none"
		if s, ok := sm.Get(r); ok {
			value = s.UserType
		}

	t, _ := template.ParseFiles("template/index.html");
	t.ExecuteTemplate(w, t.Name(), data.IndexData(value))
}
