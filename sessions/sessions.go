package sessions

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"sync"
	"time"
)

// Create a structure to represent a sessions and its associated data.
type Session struct {
	ID        string
	Data      map[string]any
	ExpiresAt time.Time
}

var SessionStore = struct {
	sync.RWMutex
	Sessions map[string]*Session
}{Sessions: make(map[string]*Session)}

func generateSessionId() string {
	b := make([]byte, 16) // 16 bytes = 128 bits

	// fill the byte slice with cryptographically random bytes.
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}

func CreateSession(w http.ResponseWriter) string {
	sessionId := generateSessionId()

	session := &Session{
		ID:        sessionId,
		Data:      make(map[string]any),
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	SessionStore.Lock()
	SessionStore.Sessions[sessionId] = session
	SessionStore.Unlock()

	// Set session id in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
	})

	return sessionId
}

func GetSession(r *http.Request) *Session {
	cookie, err := r.Cookie("session_id")

	if err != nil {
		log.Println("Unable to find cookie, session_id")
		return nil
	}

	SessionStore.Lock()
	session, exists := SessionStore.Sessions[cookie.Value]
	SessionStore.Unlock()

	if !exists || session.ExpiresAt.Before(time.Now()) {
		return nil
	}

	return session
}

func DestroySession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		SessionStore.Lock()
		delete(SessionStore.Sessions, cookie.Value)
		SessionStore.Unlock()
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1, // expires immediately
	})
}
