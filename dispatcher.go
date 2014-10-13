package queue

// Dispatcher queue wrapper
type Dispatcher struct {
	wc int
	wq chan Job
}

// NewDispatcher initialize new queue
func NewDispatcher(workerCount int, queueSize int) *Dispatcher {
	return &Dispatcher{
		wc: workerCount,
		wq: make(chan Job, queueSize),
	}
}

func (d Dispatcher) initWorkers() {
	go func() {
		for i := 1; i <= d.wc; i++ {
			w := NewWorker(i, d.wq)
			w.Run()
		}
	}()
}

// Start for job requests
func (d Dispatcher) Start() {
	d.initWorkers()
}

// Collect add job requests to work queue
func (d Dispatcher) Collect(j Job) {
	d.wq <- j
}
