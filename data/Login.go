package data

import (
	"fmt"
	
	"github.com/gayanch/go-blog/dbconn"
)

func Auth(username, password string) (bool, string) {
	sql := fmt.Sprintf("select type from login, loginlevel where login.levelid = loginlevel.id AND username = '%s' and password = sha2('%s', 0)", username, password)
	rows := dbconn.Query(sql)
	defer rows.Close()
	//fmt.Println(rows)
	if (rows.Next()) {
		var level string	
		rows.Scan(&level)
		fmt.Println(level)
		return true, level
	}
	return false, ""
}
