package data

import (
    "fmt"

    "github.com/gayanch/go-blog/db"
)

type popular struct {
  Id int
  Title string
}

type home struct {
  Title string
  Popular []popular
  Admin bool
}

func Home() home {
    sql := fmt.Sprintf("SELECT post.pid, title FROM post, hits WHERE hits.pid = post.pid ORDER BY hits DESC LIMIT 5 ")
    p := make([]popular, 0)

    rows := db.Query(sql)
    defer rows.Close()
    for rows.Next() {
        var tmp popular
        rows.Scan(&tmp.Id, &tmp.Title)
        p = append(p, tmp)
    }
    h := home{Title: blogconf["title"], Popular:p}
    return h
}
