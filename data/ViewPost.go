package data

import (
    "time"
    "fmt"

    "github.com/gayanch/go-blog/db"
)

type post struct {
    Time time.Time
    Title string
    Body string
}

func ViewPost(pid string) post {
    sql := fmt.Sprintf("SELECT time, title, body FROM post WHERE pid = %s", pid)
    rows := db.Query(sql)
    defer rows.Close()

    if rows.Next() {
        var p post
        rows.Scan(&p.Time, &p.Title, &p.Body)
        p.Title = p.Title + " - " + blogconf["title"]

        //update hit counters
        sql = fmt.Sprintf("UPDATE hits SET hits = hits + 1 WHERE pid = %s", pid)
        db.Insert(sql)

        return p
    }
    return post{Title:"Post Not Found"}
}
