package orm

import (
    "database/sql"
    "fmt"

    _ "github.com/ziutek/mymysql/godrv"
    "github.com/gayanch/go-blog/config"
)

var (
    blogconf map[string]string
    dbconf map[string]string
    conStr string
)

//load xml xonfiguration at init
func init() {
    if c, err := config.ReadBlogConfig(); err == nil {
        blogconf = make(map[string]string)
        for _, xml := range c.Configs {
            blogconf[xml.Key] = xml.Value
        }
    }

    if c, err := config.ReadDbConfig(); err == nil {
        dbconf = make(map[string]string)
        for _, xml := range c.Configs {
            dbconf[xml.Key] = xml.Value
        }
    }

    conStr = fmt.Sprintf("tcp:%s*%s/%s/%s", dbconf["db-host"], dbconf["db-name"], dbconf["db-user"], dbconf["db-passwd"])
}

func connect() *sql.DB {
    db, err := sql.Open("mymysql", conStr)
    if err != nil {
        fmt.Println("Connection error: ", err)
    }
    return db
}
