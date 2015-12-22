package handler

import (
	"strconv"
	"html/template"	
	"net/http"
	
	"github.com/gayanch/go-blog/data"
)

func Page(w http.ResponseWriter, r *http.Request) {
	pid, _ := strconv.Atoi( r.URL.Path[ len("/p/"): ] )
	//fmt.Println(pid)
	page := data.PageByNumber(pid)
	
	t, _ := template.ParseFiles("template/page.html")
	t.ExecuteTemplate(w, t.Name(), page)
}
