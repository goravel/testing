package listeners

import (
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/facades"
)

type TestSyncListener struct {
}

func (receiver *TestSyncListener) Signature() string {
	return "test_sync_listener"
}

func (receiver *TestSyncListener) Queue(args ...interface{}) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestSyncListener) Handle(args ...interface{}) error {
	facades.Log.Infof("test_sync_listener: %s, %d", args[0], args[1])

	return nil
}
