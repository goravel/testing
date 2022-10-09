package events

import "github.com/goravel/framework/contracts/event"

type TestEvent struct {
}

func (receiver *TestEvent) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}
