package poller

import (
	"fmt"
	"testing"
	"time"
)

func TestLooper(t *testing.T) {
	count := 0
	poller := NewPoller(func() {
		count++
	}, time.Second)
	defer poller.Stop()

	go poller.Run()

	time.Sleep(5 * time.Second)
	fmt.Println("Loop:", count)
}
