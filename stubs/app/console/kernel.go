package console

import (
	"time"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
	"github.com/goravel/framework/facades"

	"goravel/testing/resources/commands"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	return []schedule.Event{
		facades.Schedule.Call(func() {
			facades.Log.Info("schedule closure immediately")
		}).EveryMinute(),
		facades.Schedule.Call(func() {
			time.Sleep(61 * time.Second)
			facades.Log.Info("schedule closure DelayIfStillRunning")
		}).EveryMinute().DelayIfStillRunning(),
		facades.Schedule.Call(func() {
			time.Sleep(61 * time.Second)
			facades.Log.Info("schedule closure SkipIfStillRunning")
		}).EveryMinute().SkipIfStillRunning(),
		facades.Schedule.Command("test --name Goravel argument0 argument1").EveryMinute(),
	}
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{
		&commands.Test{},
	}
}
