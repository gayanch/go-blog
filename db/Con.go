package db

import (
  "database/sql"
  "fmt"
  _ "github.com/ziutek/mymysql/godrv"
)

func connect() *sql.DB {
    host := dbconf["db-host"]
    dbname := dbconf["db-name"]
    user := dbconf["db-user"]
    passwd := dbconf["db-passwd"]

    conStr := fmt.Sprintf("tcp:%s*%s/%s/%s", host, dbname, user, passwd)
    dbc, err := sql.Open("mymysql", conStr)
    if err != nil {
        panic(err)
    }
    return dbc
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
