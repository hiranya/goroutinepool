# goroutinepool
A simple goroutine pool for Golang.

goroutinepool uses channels as a semaphore to implement this.

## Documentation
https://godoc.org/github.com/hiranya/goroutinepool

## Example
See example/example1.go for a working example
```go
package main

import (
	"fmt"
	"github.com/hiranya/goroutinepool"
)

func main() {
	cmd := &CounterCmd{}

	pool := goroutinepool.Pool{}

	// param1: number of iterations
	// param2: level of concurrency required
	// param3: the command/job to execute
	pool.Run(30, 10, cmd)
}

type CounterCmd struct {
}

// Provide an Exec() function to implement the Command interface provided by goroutinepool library
func (cmd *CounterCmd) Exec() {
	// -----
	// your time consuming job
	// -----
}

```