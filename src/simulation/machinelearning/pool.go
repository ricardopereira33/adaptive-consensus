package machinelearning

// Pool is a pool of workers
type Pool [] *Worker

// Len returns the pool length
func (p Pool) Len() int {
	return len(p)
}

// Less comparated if the element in position i is less than the element in the position j
func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

// Swap changes the element in position i to the position j, and vice versa
func (p *Pool) Swap(i, j int) {
	a := *p
	a[i], a[j] = a[j], a[i]
	a[i].idx = i
	a[j].idx = j
}

// Push adds a new element to the Pool
func (p *Pool) Push(x interface{}) {
	n := len(*p)
	item := x.(*Worker)
	item.idx = n

	*p = append(*p, item)
}

// Pop returns the worker more available
func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n - 1]
	item.idx = -1
	*p = old[0 : n - 1]

	return item
}