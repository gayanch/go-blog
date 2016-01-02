package orm

import (
    "github.com/gayanch/go-blog/db/entity"
)

func Editor() entity.Editor {
    var editor entity.Editor

    editor.Title = "Edit Posts - " + blogconf["title"]

    db := connect()
    defer db.Close()

    editor.Posts = make([]entity.Post, 0)
    rows, _ := db.Query("SELECT pid, time, title FROM post ORDER BY pid DESC")
    defer rows.Close()

    for rows.Next() {
        var post entity.Post
        rows.Scan(&post.Pid, &post.Time, &post.Title)
        editor.Posts = append(editor.Posts, post)
    }

    return editor
}
