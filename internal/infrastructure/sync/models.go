package sync

import "time"

type Request struct {
	ID         string
	chResponse chan []byte
	chError    chan error
	createdAt  time.Time
}
