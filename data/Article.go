package data

import (
	"fmt"
	"time"
	
	"github.com/gayanch/go-blog/dbconn"
)

type Article struct {
	Pid string
	Title string
	Body string
	Time string
}

func ArticleById(id string) Article {
	sql := fmt.Sprintf("SELECT pid, title, body, time FROM post WHERE pid = %s", id)
	rows := dbconn.Query(sql)
	rows.Next()
	var (
		article Article
		time time.Time
	)
	rows.Scan(&article.Pid, &article.Title, &article.Body, &time)
	rows.Close()
	article.Title = fmt.Sprintf("%s - %s", article.Title, conf["title"])
	article.Time = fmt.Sprintf("%s", time)
	fmt.Println(article.Time, time)
	return article
}
