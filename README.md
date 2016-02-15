# goroutinepool
A simple goroutine pool for Golang.

goroutinepool uses channels as a semaphore to implement this.


## Example
See example/example1.go for a working code example
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

func (cmd *CounterCmd) Exec() {
	// -----
	// your time consuming job
	// -----
}

```