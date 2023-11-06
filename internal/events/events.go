package events

import (
	"github.com/labstack/gommon/log"
)

var LibraryItemCh = make(chan LibraryItemUpsertedEvent)

type LibraryItemUpsertedEvent struct {
	Message string
}

type LibraryItemEventListener struct {
	ID int
}

func (l *LibraryItemEventListener) Listen(events <-chan LibraryItemUpsertedEvent) {
	for e := range events {
		log.Infof("Listener %d received event: %s\n", l.ID, e.Message)
	}
}
