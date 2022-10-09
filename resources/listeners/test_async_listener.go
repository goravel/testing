package listeners

import (
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/facades"
)

type TestAsyncListener struct {
}

func (receiver *TestAsyncListener) Signature() string {
	return "test_async_listener"
}

func (receiver *TestAsyncListener) Queue(args ...interface{}) event.Queue {
	return event.Queue{
		Enable:     true,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestAsyncListener) Handle(args ...interface{}) error {
	facades.Log.Infof("test_async_listener: %s, %d", args[0], args[1])

	return nil
}
