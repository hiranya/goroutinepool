package goroutinepool

import (
	"fmt"
	"testing"
	"time"
)

func TestPoolRun(t *testing.T) {
	cmd := &CounterCmd{}

	pool := Pool{}
	pool.Run(30, 10, cmd)
}

type CounterCmd struct {
}

func (cmd *CounterCmd) Exec() {
	fmt.Println("Counter:", time.Now())
	time.Sleep(time.Millisecond * 1000)
}
