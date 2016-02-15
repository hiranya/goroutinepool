package main

import (
	"fmt"
	"github.com/hiranya/goroutinepool"
	"time"
)

func main() {
	cmd := &CounterCmd{}

	pool := goroutinepool.Pool{}
	pool.Run(30, 5, cmd)
}

type CounterCmd struct {
}

// Provide an Exec() function to implement the Command interface provided by goroutinepool library
func (cmd *CounterCmd) Exec() {
	fmt.Println("Counter:", time.Now())
	time.Sleep(time.Millisecond * 2000)
}
