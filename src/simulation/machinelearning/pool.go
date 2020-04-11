package machinelearning

// Pool is a pool of workers
type Pool [] *Worker

// Len returns the pool length
func (pool Pool) Len() int {
	return len(pool)
}

// Less comparated if the element in position i is less than the element in the position j
func (pool Pool) Less(i, j int) bool {
	return pool[i].pending < pool[j].pending
}

// Swap changes the element in position i to the position j, and vice versa
func (pool *Pool) Swap(i, j int) {
	tmpPool := *pool
	tmpPool[i], tmpPool[j] = tmpPool[j], tmpPool[i]
	tmpPool[i].idx = i
	tmpPool[j].idx = j
}

// Push adds a new element to the Pool
func (pool *Pool) Push(element interface{}) {
	lenght := len(*pool)
	item := element.(*Worker)
	item.idx = lenght

	*pool = append(*pool, item)
}

// Pop returns the worker more available
func (pool *Pool) Pop() interface{} {
	oldPool := *pool
	length := len(oldPool)
	item := oldPool[length - 1]
	item.idx = -1
	*pool = oldPool[0 : length - 1]

	return item
}