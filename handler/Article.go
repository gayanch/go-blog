package handler

import (
	"net/http"
	"html/template"
	"fmt"	//debug import
	"github.com/gayanch/go-blog/data"
)

func Article(w http.ResponseWriter, r *http.Request) {
	aId := r.URL.Path[ len("/a/"): ]
	t, _ := template.ParseFiles("template/article.html")
	article := data.ArticleById(aId)
	if article.Title == "" && article.Body == "" {
		fmt.Fprintf(w, "<h1>404: Article Not Found</h1>")
	}	else {
		t.ExecuteTemplate(w, t.Name(), article)
	}
}
