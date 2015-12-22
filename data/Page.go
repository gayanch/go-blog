package data

import (
	"github.com/gayanch/go-blog/dbconn"
	"fmt"
	"strconv"
)	

var articlesPerPage, _ = strconv.Atoi(conf["articles_per_page"])

type entry struct {
	Pid string
	Title string
	Body string
}

type Page struct {
	Entry []entry
}

func PageByNumber(p int) Page {
//	sql := "select count(pid) from post"
//	rows := dbconn.Query(sql)
//	var count int
//	if (rows.Next()) {
//		rows.Scan(&count)
//	}
//	rows.Close()
//	
//	var page Page
//	if (p * articlesPerPage > count) {
//		
//	} else {
//		page.Title = make([]string, articlesPerPage)
//		page.Body = make([]string, articlesPerPage)
//		
//		rows = dbconn.Query("SELECT title, entry FROM post limit 4")
//		for i := 0; rows.Next(); i++ {
//			rows.Scan(&page.Title[i], &page.Body[i])
//		}
//		rows.Close()
//		//return 
//	}
	
	var page Page
	page.Entry = make([]entry, 5)
	fmt.Println(articlesPerPage)
	rows := dbconn.Query("SELECT pid, title, body FROM post limit 5")
	for i := 0; rows.Next(); i++ {
		rows.Scan(&page.Entry[i].Pid, &page.Entry[i].Title, &page.Entry[i].Body)
	}
	rows.Close()
	return page
}
