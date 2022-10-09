package events

import "github.com/goravel/framework/contracts/event"

type TestCancelEvent struct {
}

func (receiver *TestCancelEvent) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}
