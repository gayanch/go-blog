package handler

import (
	"net/http"
	"html/template"
	
	"github.com/gayanch/go-blog/data"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html");
	t.ExecuteTemplate(w, t.Name(), data.IndexData())
}
