package data

import (
    "fmt"

    "github.com/gayanch/go-blog/db"
)

type newPost struct {
    Title string
}

func NewPost() newPost {
    return newPost{blogconf["title"] + " - New Post"}
}

func SavePost(title, body string) {
    sql := fmt.Sprintf("INSERT INTO post(time, title, body) VALUES(now(), '%s', '%s')", title, body)
    id := db.InsertAndGetId(sql)
    fmt.Println("New Article Id:", id)

    sql = fmt.Sprintf("INSERT INTO hits VALUES(%d, 0)", id)
    ra := db.Insert(sql)
    fmt.Println("Rows Affected:", ra)
}
