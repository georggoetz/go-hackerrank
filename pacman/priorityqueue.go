package pacman

type priorityqueue struct {
	data   []node
	lookup map[string]bool
}

func newPriorityQueue(cap int) priorityqueue {
	return priorityqueue{data: make([]node, 0, cap), lookup: make(map[string]bool)}
}

func (pq priorityqueue) Len() int {
	return len(pq.data)
}

func (pq priorityqueue) Less(i, j int) bool {
	return pq.data[i].f > pq.data[j].f
}

func (pq priorityqueue) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *priorityqueue) Push(x interface{}) {
	n := x.(node)
	pq.data = append(pq.data, n)
	pq.lookup[n.key()] = true
}

func (pq *priorityqueue) Pop() interface{} {
	n := len(pq.data)
	x := (pq.data)[n-1]
	pq.data = (pq.data)[0 : n-1]
	delete(pq.lookup, x.key())
	return x
}

func (pq priorityqueue) contains(n node) bool {
	_, ok := pq.lookup[n.toString()]
	return ok
}
