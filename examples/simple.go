package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/ernestas-poskus/queue"
)

// Implements Job interface
/////////////////////////////////////////

// JobSimple struct for Worker to process
type JobSimple struct{}

// Perform - specific task or one work unit
func (j JobSimple) Perform() {
	fmt.Println("Performing job")
}

/////////////////////////////////////////

func main() {

	// NumCPU - spawn number of workers
	// 100 - Buffered asynchronous job queue
	d := queue.NewDispatcher(runtime.NumCPU(), 100)
	d.Start() // Initializes workers

	for {
		fmt.Println("Wait")
		time.Sleep(130 * time.Millisecond)
		d.Collect(&JobSimple{}) // Collects new job requests
	}
}
