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

func NewSessionManager() SessionManager {
	var sm SessionManager
	sm.Session = make(map[string]session)
	return sm
}

func(sm SessionManager) Get(r *http.Request) (session, bool) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return session{}, false
	}
	
	username := cookie.Value
	if value, ok := sm.Session[username]; ok {
		return value, true
	}
	return session{}, false
}

func (sm *SessionManager) Set(w http.ResponseWriter, username, userType string) {
	sm.Session[username] = session{userType}
	cookie := http.Cookie{Name: "session", Value: username, Expires: time.Now().Add(time.Hour)}
	http.SetCookie(w, &cookie)
}

func (sm *SessionManager) Delete(username string) {
	delete (sm.Session, username)
}
