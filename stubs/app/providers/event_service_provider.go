package providers

import (
	"goravel/testing/resources/events"
	"goravel/testing/resources/listeners"

	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/facades"
)

type EventServiceProvider struct {
}

func (receiver *EventServiceProvider) Register() {

}

func (receiver *EventServiceProvider) Boot() {
	facades.Event.Register(receiver.listen())
}

func (receiver *EventServiceProvider) listen() map[event.Event][]event.Listener {
	return map[event.Event][]event.Listener{
		&events.TestEvent{}: {
			&listeners.TestSyncListener{},
			&listeners.TestAsyncListener{},
		},
		&events.TestCancelEvent{}: {
			&listeners.TestCancelListener{},
			&listeners.TestSyncListener{},
		},
	}
}
