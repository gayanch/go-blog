package orm

import (
    "github.com/gayanch/go-blog/db/entity"
)

func Home() entity.Home {
    sql := "SELECT post.pid, title FROM post, hits WHERE hits.pid = post.pid ORDER BY hits DESC LIMIT 5 "
    p := make([]entity.Post, 0)

    db := connect()
    defer db.Close()

    rows, _ := db.Query(sql)
    defer rows.Close()

    for rows.Next() {
        var tmp entity.Post
        rows.Scan(&tmp.Pid, &tmp.Title)
        p = append(p, tmp)
    }
    h := entity.Home{Title: blogconf["title"], Popular:p}
    return h
}
