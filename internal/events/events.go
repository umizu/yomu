package events

import (
	"reflect"

	"github.com/labstack/gommon/log"
	"github.com/umizu/yomu/internal/data"
)

type event interface {
	handle()
}

func Listen(messages <-chan interface{}) {
	for msg := range messages {
		switch v := msg.(type) {
		case event:
			v.handle()
		default:
			log.Errorf("unknown event: %v", reflect.TypeOf(v))
		}
	}
}

type LibraryItemUpsertedEvent struct {
	Message string
	Store   data.LibraryItemStore
}

func (e LibraryItemUpsertedEvent) handle() {
	log.Infof("event message: %s", e.Message)
}
