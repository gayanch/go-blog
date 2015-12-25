package db

import (
    "github.com/gayanch/go-blog/config"
)

var dbconf map[string]string

func init() {
    if conf, err := config.ReadDbConfig(); err == nil {
        dbconf = make(map[string]string)
        for _, xml := range conf.Configs {
          dbconf[xml.Key] = xml.Value
        }
    }
}
