package session

import (
    "net/http"
    "time"
)

type session struct {
    UserType string
}

type SessionManager struct {
    Session map[string]session
}

func NewSessionManager() SessionManager{
    var sm SessionManager
    sm.Session = make(map[string]session)
    return sm
}

func (s *SessionManager) Set(w http.ResponseWriter, username, userType string) {
    expires := time.Now().Add(time.Hour)
    cookie := http.Cookie{Name: "session", Value: username, Expires: expires}
    http.SetCookie(w, &cookie)
    s.Session[username] = session{userType}
}

func (s SessionManager) Get(r *http.Request) (session, bool) {
    if cookie, err := r.Cookie("session"); err == nil && cookie != nil {
        if s, ok := s.Session[cookie.Value]; ok {
            return s, true
        } else {
            return session{}, false
        }
    }
    return session{}, false
}

func (s *SessionManager) Delete(r *http.Request) {
    if cookie, err := r.Cookie("session"); err == nil && cookie != nil {
        delete(s.Session, cookie.Value)
    }
}
