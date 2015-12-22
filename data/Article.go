package data

import (
	"fmt"
	
	"github.com/gayanch/go-blog/dbconn"
)

type Article struct {
	Title string
	Body string
}

func ArticleById(id string) Article {
	sql := fmt.Sprintf("SELECT title, body FROM post WHERE pid = %s", id)
	rows := dbconn.Query(sql)
	rows.Next()
	var article Article
	rows.Scan(&article.Title, &article.Body)
	rows.Close()
	article.Title = fmt.Sprintf("%s - %s", article.Title, conf["title"])
	return article
}
