package data

import (
	"fmt"

	"github.com/gayanch/go-blog/dbconn"
)

func Auth(username, password string) bool {
	sql := fmt.Sprintf("select username from login where username = '%s' and password = sha2('%s', 0)", username, password)
	rows := dbconn.Query(sql)
	defer rows.Close()
	if (rows.Next()) {
	 return true
	}
	return false
}
