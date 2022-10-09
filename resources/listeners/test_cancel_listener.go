package listeners

import (
	"errors"

	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/facades"
)

type TestCancelListener struct {
}

func (receiver *TestCancelListener) Signature() string {
	return "test_cancel_listener"
}

func (receiver *TestCancelListener) Queue(args ...interface{}) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *TestCancelListener) Handle(args ...interface{}) error {
	facades.Log.Infof("test_cancel_listener: %s, %d\n", args[0], args[1])

	return errors.New("cancel")
}
