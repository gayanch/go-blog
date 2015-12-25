package handler

import (
    "github.com/gayanch/go-blog/session"
)

var sm session.SessionManager

func init() {
    //starting session manager
    sm = session.NewSessionManager()
}
