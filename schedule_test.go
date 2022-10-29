package testing

import (
	"context"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
	"time"

	"goravel/bootstrap"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/file"
	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	bootstrap.Boot()

	second, _ := strconv.Atoi(time.Now().Format("05"))
	// Make sure run 3 times
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(120+6+60-second)*time.Second)
	go func(ctx context.Context) {
		facades.Schedule.Run()

		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	time.Sleep(time.Duration(120+5+60-second) * time.Second)
	log := fmt.Sprintf("storage/logs/goravel-%s.log", time.Now().Format("2006-01-02"))
	assert.True(t, file.Exist(log))
	data, err := ioutil.ReadFile(log)
	assert.Nil(t, err)
	assert.Equal(t, 3, strings.Count(string(data), "schedule closure immediately"))
	assert.Equal(t, 3, strings.Count(string(data), "Run test command success, argument_0: argument0, argument_1: argument1, option_name: Goravel, option_age: 18, arguments: argument0,argument1"))
	assert.Equal(t, 2, strings.Count(string(data), "schedule closure DelayIfStillRunning"))
	assert.Equal(t, 1, strings.Count(string(data), "schedule closure SkipIfStillRunning"))
	assert.True(t, file.Remove("./storage"))
}
