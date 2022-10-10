package providers

import (
	contractevent "github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/facades"

	"goravel/testing/resources/events"
	"goravel/testing/resources/listeners"
)

type EventServiceProvider struct {
}

func (receiver *EventServiceProvider) Register() {

}

func (receiver *EventServiceProvider) Boot() {
	facades.Event.Register(receiver.listen())
}

func (receiver *EventServiceProvider) listen() map[contractevent.Event][]contractevent.Listener {
	return map[contractevent.Event][]contractevent.Listener{
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
