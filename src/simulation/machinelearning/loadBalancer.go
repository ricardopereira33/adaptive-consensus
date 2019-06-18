package machinelearning

import (
	"container/heap"
)

// Request represent a request package
type Request struct {
	data 			[][][]float32
	responseChannel chan [][][]float32
}

// Balancer represents
type Balancer struct {
	id       int
	pool     Pool
	requests chan Request
	done 	 chan *Worker
}

// NewBalancer creates a new balancer
func NewBalancer(id, nWorkers, nRequester int) *Balancer {
	done := make(chan *Worker, nWorkers)
	requests := make(chan Request)

	balancer := &Balancer{id, make(Pool, 0, nWorkers), requests, done}

	for i := 0; i < nWorkers; i++ {
		worker := newWorker(make(chan Request, nRequester), nRequester)

		heap.Push(&balancer.pool, worker)

		go worker.doWork(balancer.done)
	}

	go balancer.balance()

	return balancer
}

// CreateRequest submit a new request to a worker
func (balancer *Balancer) CreateRequest(data [][][]float32) [][][]float32 {
	responseChannel := make(chan [][][]float32)

	balancer.requests <- Request{data, responseChannel}

	result := <- responseChannel

	return result
}

func (balancer *Balancer) balance() {
	for {
		select {
		case request := <- balancer.requests:
			balancer.dispatch(request)

		case worker := <- balancer.done:
			balancer.completed(worker)
		}
	}
}

func (balancer *Balancer) dispatch(request Request) {
	worker := heap.Pop(&balancer.pool).(*Worker)

	go func (){
		worker.work <- request
	}()

	worker.pending++
	heap.Push(&balancer.pool, worker)
}

func (balancer *Balancer) completed(worker *Worker) {
	worker.pending--

	heap.Remove(&balancer.pool, worker.idx)
	heap.Push(&balancer.pool, worker)
}
