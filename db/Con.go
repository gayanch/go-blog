package db

import (
  "database/sql"
  "fmt"
  _ "github.com/ziutek/mymysql/godrv"
)

const (
    db = "goblog"
    user = "root"
    passwd = "123"
)

func connect() *sql.DB {
    conStr := fmt.Sprintf("tcp:localhost:3306*%s/%s/%s", db, user, passwd)
    db, err := sql.Open("mymysql", conStr)
    if err != nil {
        panic(err)
    }
    return db
}

func Query(sql string) *sql.Rows {
    rows, err := connect().Query(sql)
    if err != nil {
        panic(err)
    }
    return rows
}

func Insert(sql string) int64 {
    stmt, _ := connect().Prepare(sql)
    defer stmt.Close()

    res, _ := stmt.Exec()
    c, _ := res.RowsAffected()
    return c
}

func InsertAndGetId(sql string) int64 {
    stmt, _ := connect().Prepare(sql)
    defer stmt.Close()

    res, _ := stmt.Exec()
    id, _ := res.LastInsertId()
    return id
}
