package queue

// Worker structure
type Worker struct {
	ID   int
	wq   <-chan Job // Unidirectional Work Queue
	stop chan bool
}

// NewWorker creates new Worker struct
func NewWorker(id int, workQueue chan Job) *Worker {
	return &Worker{
		ID:   id,
		wq:   workQueue,
		stop: make(chan bool, 1),
	}
}

// Run starts worker
func (w Worker) Run() {
	go func() {
		for {
			select {
			case work := <-w.wq:
				work.Perform()
			case <-w.stop:
				return
			}

		}
	}()
}

// Stop sends signal to worker to process left Job's and stop
func (w Worker) Stop() {
	go func() {
		w.stop <- true
	}()
}
