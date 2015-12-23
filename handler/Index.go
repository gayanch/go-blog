package handler

import (
	"net/http"
	"html/template"
	"fmt"
	"github.com/gayanch/go-blog/data"
)

func Index(w http.ResponseWriter, r *http.Request) {
	value := "none"
	if cookie, err := r.Cookie("type"); err == nil && cookie != nil {
		fmt.Println("Index: ", cookie.Value)
		value = cookie.Value
	}
		
	t, _ := template.ParseFiles("template/index.html");
	t.ExecuteTemplate(w, t.Name(), data.IndexData(value))
}
