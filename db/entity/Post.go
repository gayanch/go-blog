package entity

import "time"

type Post struct {
    Pid int
    Time time.Time
    Title string
    Body string
}
