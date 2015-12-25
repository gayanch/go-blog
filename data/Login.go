package data

import (
    "fmt"

    "github.com/gayanch/go-blog/db"
)

type login struct {
  Title string
}

func Login() login {
  return login{"Go Blog - Login"}
}

func Auth(username, password string) (string, bool) {
    sql := fmt.Sprintf("SELECT type FROM login, loginlevel WHERE levelid = id AND username='%s' AND password=sha2('%s', 0)", username, password)
    rows := db.Query(sql)
    defer rows.Close()

    if rows.Next() {
        var userType string
        rows.Scan(&userType)
        return userType, true
    }
    return "none", false
}
