package handler

import (
    "github.com/gayanch/go-blog/session"
)

var sm session.SessionManager

func init() {
    sm = session.NewSessionManager()
}
