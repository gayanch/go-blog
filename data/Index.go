package data

import (
	"fmt"
	
	"github.com/gayanch/go-blog/dbconn"
	"github.com/gayanch/go-blog/config"
)

var conf map[string]string

type Popular struct {
	Link string
	Name string
}

type Index struct {
	Title string
	Admin bool
	MostViewed []Popular
}

func IndexData(user string) Index {
	var index Index
	index.Title = conf["title"]
	
	fmt.Println("Index data: ", user)
	if user == "admin" {
		index.Admin = true
	} else {
		index.Admin = false
	}
	
	index.MostViewed = make([]Popular, 0)
	
	rows := dbconn.Query("SELECT post.pid, title FROM post, hits where post.pid = hits.pid order by hits.hits desc limit 5")
	for i := 0; rows.Next(); i++ {
		var pid, title string
		rows.Scan(&pid, &title)
//		index.MostViewed[i].Link = fmt.Sprintf("/a/%s", pid)
//		index.MostViewed[i].Name = title
		p := Popular{ fmt.Sprintf("/a/%s", pid), title }
		index.MostViewed = append(index.MostViewed, p)
	}
	rows.Close()
	return index
}

//Init blog configuration
func init() {
	conf = make(map[string]string)
	c, _ := config.ReadBlogConfig()
	for _, v := range c.Configs {
		conf[v.Key] = v.Value
	}
}
