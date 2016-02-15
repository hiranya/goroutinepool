// A simple goroutine pool for Golang
package goroutinepool

// Command: Implement the Command interface by providing an Exec() function in your structs.
// This Exec() function encapsulates the logic that you want to execute in goroutines.
type Command interface {
	Exec()
}

// A helper struct to manage the pool
type Pool struct{}

// The Run() function takes
// numOfIterations: number of times this command should execute
// concurrency: The pool size or the maximum number of goroutines that should execute at any given time
// cmd: The implementation of the goroutinepool.Command interface
func (pool *Pool) Run(numOfIterations int, concurrency int, cmd Command) {

	c := make(chan int, concurrency)

	cnt := make(chan bool)
	counter := 0

	for i := 0; i < numOfIterations; i++ {
		c <- 1
		go func(cmd Command) {
			cmd.Exec()
			<-c
			cnt <- true
		}(cmd)
	}

	for {
		if <-cnt {
			counter++
			if counter == numOfIterations {
				break
			}
		}
	}
}
