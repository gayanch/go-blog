package orm

import (
    "github.com/gayanch/go-blog/db/entity"
)

//get post by pid
func Post(id string) (entity.Post, bool) {
    db := connect()
    defer db.Close()

    stmt, _ := db.Prepare("SELECT pid, time, title, body FROM post WHERE pid = ?")
    defer stmt.Close()

    rows, _ := stmt.Query(id)
    defer rows.Close()

    if rows.Next() {
        var post entity.Post
        rows.Scan(&post.Pid, &post.Time, &post.Title, &post.Body)
        post.Title = post.Title + " - " + blogconf["title"]
        return post, true
    } else {
        return entity.Post{}, false
    }
}
