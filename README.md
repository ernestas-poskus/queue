queue
=====

Golang channel queue-worker implementation

- Dispatcher: initializes queue & workers
- Collector: distributes work requests to Dispatcher
- Worker: performs unit of work
- Job: interface for implementation
