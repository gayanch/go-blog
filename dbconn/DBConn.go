package dbconn

import (
	_ "github.com/ziutek/mymysql/godrv"
	"github.com/gayanch/go-blog/config"
	"fmt"
	"database/sql"
)

var dbConfig map[string]string

func init() {
	dbConfig = make(map[string]string)
	conf, err := config.ReadDbConfig();

	if (err != nil) {
		panic(err)
	}

	for _, v := range conf.Configs {
		dbConfig[v.Key] = v.Value
	}
}

func connect() *sql.DB{
	conStr := fmt.Sprintf("tcp:%s*%s/%s/%s", dbConfig["db-host"], dbConfig["db-name"], dbConfig["db-user"], dbConfig["db-passwd"])
	db, err  := sql.Open("mymysql", conStr)
	if(err != nil) {
		panic(err)
	}
	return db
}

func Query(sql string) *sql.Rows {
	db := connect()
	rows, _ := db.Query(sql)
	return rows
}

func Insert(sql string) int64 {
	db := connect()
	defer db.Close()

	stmt, _ := db.Prepare(sql)
	defer stmt.Close()

	res, _ := stmt.Exec()

	id, _ := res.LastInsertId()
	return id
}
