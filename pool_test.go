package goroutinepool

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPoolRun(t *testing.T) {
	const numOfIterations int = 100
	cmd := &CounterCmd{}
	cmd.DelayStream = make(chan time.Duration, numOfIterations)

	pool := Pool{}

	// Pass the parameters to the channel
	for i := 0; i < numOfIterations; i++ {
		cmd.DelayStream <- time.Duration(rand.Intn(5))
	}

	// Run pool
	pool.Run(numOfIterations, 10, cmd)
}

type CounterCmd struct {
	DelayStream chan time.Duration
}

func (cmd *CounterCmd) Exec() {
	fmt.Println("Counter:", time.Now())
	time.Sleep(time.Millisecond * 1000 * <-cmd.DelayStream)
}
