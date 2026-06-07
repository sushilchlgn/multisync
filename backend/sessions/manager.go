package sessions

import (
	"fmt"
	"time"
)

func CreateSession(name string, device string) *Session {

	id := fmt.Sprintf("%d", time.Now().UnixNano())

	session := &Session{
		ID:     id,
		Name:   name,
		Device: device,
		URL:    "",
		Status: "created",
	}

	sessionStore[id] = session

	return session
}

func GetSessions() []*Session {

	list := []*Session{}

	for _, s := range sessionStore {
		list = append(list, s)
	}

	return list
}

func DeleteSession(id string) {

	delete(sessionStore, id)
}