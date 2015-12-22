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
	MostViewed []Popular
}

func IndexData() Index {
	var index Index
	index.Title = conf["title"]
	
	index.MostViewed = make([]Popular, 5)
	
	rows := dbconn.Query("SELECT post.pid, title FROM post, hits where post.pid = hits.pid order by hits.hits desc limit 5")
	for i := 0; rows.Next(); i++ {
		var pid, title string
		rows.Scan(&pid, &title)
		index.MostViewed[i].Link = fmt.Sprintf("/a/%s", pid)
		index.MostViewed[i].Name = title
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
