package data

import (
	"github.com/gayanch/go-blog/dbconn"
	"fmt"
	"time"
)	

type Page struct {
	Entry []Article
}

func PageByNumber(p int) Page {
	var page Page
	page.Entry = make([]Article, 0)
	sql := fmt.Sprintf("SELECT pid, title, body, time FROM post limit %s", conf["articles_per_page"])
	rows := dbconn.Query(sql)
	
	var (
		pid, title, body string
		time time.Time
	)
	for i := 0; rows.Next(); i++ {
		rows.Scan(&pid, &title, &body, &time)
		a := Article{ pid, title, body, fmt.Sprintf("%s", time) }
		page.Entry = append(page.Entry, a)
	}
	rows.Close()
	return page
}
