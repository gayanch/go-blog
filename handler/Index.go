package handler

import (
	"net/http"
	"html/template"
	//"fmt"
	"github.com/gayanch/go-blog/data"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
		
	t, _ := template.ParseFiles("template/index.html");
	t.ExecuteTemplate(w, t.Name(), data.IndexData(cookie.Value))
}
