package data

import (
  "github.com/gayanch/go-blog/dbconn"

  "fmt"
)

func SaveAricle(title, body string) {
  sql := fmt.Sprintf("INSERT INTO post(time, title, body) VALUES (now(), '%s', '%s')", title, body)
  id := dbconn.Insert(sql)
  fmt.Println(id)
  // sql = fmt.Sprintf("INSERT INTO hit values(%d, 0)", id)
  // dbconn.Query(sql).Close()
}
