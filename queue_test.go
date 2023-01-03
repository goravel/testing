package testing

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"goravel/bootstrap"
	"goravel/testing/resources/jobs"

	"github.com/goravel/framework/contracts/queue"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type QueueTestSuite struct {
	suite.Suite
}

func TestQueueTestSuite(t *testing.T) {
	file.Remove("./storage")
	bootstrap.Boot()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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
	go func(ctx context.Context) {
		if err := facades.Queue.Worker(&queue.Args{
			Connection: "test",
			Queue:      "test1",
			Concurrent: 2,
		}).Run(); err != nil {
			facades.Log.Errorf("Queue run error: %v", err)
		}

		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		if err := facades.Queue.Worker(&queue.Args{
			Connection: "redis",
			Queue:      "test1",
			Concurrent: 2,
		}).Run(); err != nil {
			facades.Log.Errorf("Queue run error: %v", err)
		}

		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	suite.Run(t, new(QueueTestSuite))
}

func (s *QueueTestSuite) SetupTest() {
}

func (s *QueueTestSuite) TestMakeJob() {
	t := s.T()
	Equal(t, "make:job TestJob", "Job created successfully")
	assert.True(t, file.Exists("./app/jobs/test_job.go"))
	assert.True(t, file.Remove("./app"))
}

func (s *QueueTestSuite) TestSyncQueue() {
	t := s.T()
	assert.Nil(t, facades.Queue.Job(&jobs.TestSyncJob{}, []queue.Arg{
		{Type: "string", Value: "TestSyncQueue"},
		{Type: "int", Value: 1},
	}).DispatchSync())

	log := fmt.Sprintf("storage/logs/goravel-%s.log", time.Now().Format("2006-01-02"))
	assert.True(t, file.Exists(log))
	data, err := ioutil.ReadFile(log)
	assert.Nil(t, err)
	assert.True(t, strings.Contains(string(data), "test_sync_job: TestSyncQueue, 1"))
}

func (s *QueueTestSuite) TestDefaultAsyncQueue() {
	s.Nil(facades.Queue.Job(&jobs.TestAsyncJob{}, []queue.Arg{
		{Type: "string", Value: "TestDefaultAsyncQueue"},
		{Type: "int", Value: 1},
	}).Dispatch())

	time.Sleep(3 * time.Second)
	log := fmt.Sprintf("storage/logs/goravel-%s.log", time.Now().Format("2006-01-02"))
	s.True(file.Exists(log))
	data, err := ioutil.ReadFile(log)
	s.Nil(err)
	s.True(strings.Contains(string(data), "test_async_job: TestDefaultAsyncQueue, 1"))
}

func (s *QueueTestSuite) TestCustomAsyncQueue() {
	s.Nil(facades.Queue.Job(&jobs.TestAsyncJob{}, []queue.Arg{
		{Type: "string", Value: "TestCustomAsyncQueue"},
		{Type: "int", Value: 1},
	}).OnConnection("test").OnQueue("test1").Dispatch())
	time.Sleep(2 * time.Second)
	log := fmt.Sprintf("storage/logs/goravel-%s.log", time.Now().Format("2006-01-02"))
	s.True(file.Exists(log))
	data, err := ioutil.ReadFile(log)
	s.Nil(err)
	s.True(strings.Contains(string(data), "test_async_job: TestCustomAsyncQueue, 1"))
}

func (s *QueueTestSuite) TestErrorAsyncQueue() {
	s.Nil(facades.Queue.Job(&jobs.TestAsyncJob{}, []queue.Arg{
		{Type: "string", Value: "TestErrorAsyncQueue"},
		{Type: "int", Value: 1},
	}).OnConnection("redis").OnQueue("test2").Dispatch())
	time.Sleep(2 * time.Second)
	log := fmt.Sprintf("storage/logs/goravel-%s.log", time.Now().Format("2006-01-02"))
	if file.Exists(log) {
		data, err := ioutil.ReadFile(log)
		s.Nil(err)
		s.False(strings.Contains(string(data), "test_async_job: TestErrorAsyncQueue, 1"))
	}
}

func (s *QueueTestSuite) TestChainAsyncQueue() {
	s.Nil(facades.Queue.Chain([]queue.Jobs{
		{
			Job: &jobs.TestAsyncJob{},
			Args: []queue.Arg{
				{Type: "string", Value: "TestChainAsyncQueue"},
				{Type: "int", Value: 1},
			},
		},
		{
			Job: &jobs.TestSyncJob{},
			Args: []queue.Arg{
				{Type: "string", Value: "TestChainSyncQueue"},
				{Type: "int", Value: 1},
			},
		},
	}).Dispatch())
	time.Sleep(2 * time.Second)
	log := fmt.Sprintf("storage/logs/goravel-%s.log", time.Now().Format("2006-01-02"))
	s.True(file.Exists(log))
	data, err := ioutil.ReadFile(log)
	s.Nil(err)
	s.True(strings.Contains(string(data), "test_sync_job: TestChainSyncQueue, 1"))
	s.True(strings.Contains(string(data), "test_async_job: TestChainAsyncQueue, 1"))
}
