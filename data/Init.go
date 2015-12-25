package data

import (
    "github.com/gayanch/go-blog/config"
)


var blogconf map[string]string

func init() {
    //reading blog configuration from ConfigManager
    if c, err := config.ReadBlogConfig(); err == nil {
        blogconf = make(map[string]string)
        for _, xml := range c.Configs {
            blogconf[xml.Key] = xml.Value
        }
    }
}
