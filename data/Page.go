package data

import (
    "fmt"
    "time"

    "github.com/gayanch/go-blog/db"
)

type pagePost struct {
    Pid string
    Time time.Time
    Title string
    Body string
}

type page []pagePost

func Page(pno string) page {
    p := make(page, 0)
    sql := fmt.Sprintf("SELECT pid, time, title, body FROM post ORDER BY pid DESC")
    rows := db.Query(sql)
    defer rows.Close()

    for rows.Next() {
        var tmp pagePost
        rows.Scan(&tmp.Pid, &tmp.Time, &tmp.Title, &tmp.Body)

        if (len(tmp.Body) > 100) {
            //we are not showing whole body in page view
            tmp.Body = tmp.Body[:100]
        }
        p = append(p, tmp)
    }
    return p
}
