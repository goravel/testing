package testing

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"goravel/bootstrap"
	"goravel/testing/resources/events"

	contractevents "github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventTestSuite struct {
	suite.Suite
}

func TestEventTestSuite(t *testing.T) {
	file.Remove("./storage")
	file.Remove("./app")

	bootstrap.Boot()

	suite.Run(t, new(EventTestSuite))
}

func (s *EventTestSuite) SetupTest() {

}

func (s *EventTestSuite) TestMakeEvent() {
	t := s.T()
	Equal(t, "make:event OrderShipped", "Event created successfully")
	assert.True(t, file.Exists("./app/events/order_shipped.go"))
}

func (s *EventTestSuite) TestMakeListener() {
	t := s.T()
	Equal(t, "make:listener SendShipmentNotification", "Listener created successfully")
	assert.True(t, file.Exists("./app/listeners/send_shipment_notification.go"))
}

func (s *EventTestSuite) TestEvent() {
	t := s.T()
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	go func(ctx context.Context) {
		if err := facades.Queue.Worker(nil).Run(); err != nil {
			facades.Log.Errorf("Queue run error: %v", err)
		}

		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	assert.Nil(t, facades.Event.Job(&events.TestEvent{}, []contractevents.Arg{
		{Type: "string", Value: "Goravel"},
		{Type: "int", Value: 1},
	}).Dispatch())

	log := fmt.Sprintf("storage/logs/goravel-%s.log", time.Now().Format("2006-01-02"))
	assert.True(t, file.Exists(log))
	time.Sleep(3 * time.Second)
	data, err := ioutil.ReadFile(log)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(string(data), "test_sync_listener: Goravel, 1"))
	assert.True(t, strings.Contains(string(data), "test_async_listener: Goravel, 1"))
}

func (s *EventTestSuite) TestCancelEvent() {
	t := s.T()
	assert.EqualError(t, facades.Event.Job(&events.TestCancelEvent{}, []contractevents.Arg{
		{Type: "string", Value: "Goravel"},
		{Type: "int", Value: 1},
	}).Dispatch(), "cancel")

	log := fmt.Sprintf("storage/logs/goravel-%s.log", time.Now().Format("2006-01-02"))
	assert.True(t, file.Exists(log))
	data, err := ioutil.ReadFile(log)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(string(data), "test_cancel_listener: Goravel, 1"))
	assert.False(t, strings.Contains(string(data), "test_sync_listener: Goravel, 1"))
}
