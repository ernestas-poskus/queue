package queue

// Job queue interface
type Job interface {
	Perform()
}
