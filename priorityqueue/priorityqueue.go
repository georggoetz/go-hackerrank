package priorityqueue

import "container/heap"

// PriorityComparer must be implemented for any type of priority.
type PriorityComparer interface {
	// Less decides if it is a min- or max-priority queue.
	Less(interface{}, interface{}) bool
}

// Item is an element of the priority queue.
type Item struct {
	Value    interface{}
	priority interface{}
	index    int
}

// PriorityQueue is similar to a stack or queue, but additionally each element
// has a priority associated with it. An element with high priority is served
// before elements with lower priority.
type PriorityQueue struct {
	heap heapStorage
}

// Init initializes or clears the priority queue.
func (pq *PriorityQueue) Init(c PriorityComparer) *PriorityQueue {
	pq.heap = heapStorage{data: make([]*Item, 0), comp: c}
	return pq
}

// New returns an initialized priority queue.
func New(c PriorityComparer) *PriorityQueue {
	return new(PriorityQueue).Init(c)
}

// Len returns the length.
func (pq *PriorityQueue) Len() int {
	return pq.heap.Len()
}

// Push pushes a new item with the given value and priority into the priority
// queue and returns the item.
func (pq *PriorityQueue) Push(val interface{}, prio interface{}) *Item {
	item := &Item{Value: val, priority: prio}
	heap.Push(&pq.heap, item)
	return item
}

// Pop removes the item with the highest priority, according to Comparer.Less().
// If the queue is empty it returns nil.
func (pq *PriorityQueue) Pop() *Item {
	if pq.Len() > 0 {
		return heap.Pop(&pq.heap).(*Item)
	}
	return nil
}

// Top returns but does not remove the item with the highest priority.
func (pq *PriorityQueue) Top() *Item {
	if n := pq.Len(); n > 0 {
		return pq.heap.data[0]
	}
	return nil
}

// Fix re-establishes the heap ordering after the priority of the given item
// has changed. Equivalent to, but less expensive, to calling Remove() followed
// by Push().
func (pq *PriorityQueue) Fix(i *Item, prio interface{}) {
	i.priority = prio
	heap.Fix(&pq.heap, i.index)
}

// Remove removes the given item.
func (pq *PriorityQueue) Remove(i *Item) *Item {
	return heap.Remove(&pq.heap, i.index).(*Item)
}

// See examples from container/heap.
type heapStorage struct {
	data []*Item
	comp PriorityComparer
}

func (s heapStorage) Len() int {
	return len(s.data)
}

func (s heapStorage) Less(i, j int) bool {
	return s.comp.Less(s.data[i].priority, s.data[j].priority)
}

func (s heapStorage) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
	s.data[i].index = i
	s.data[j].index = j
}

func (s *heapStorage) Push(x interface{}) {
	n := len(s.data)
	item := x.(*Item)
	item.index = n
	s.data = append(s.data, item)
}

func (s *heapStorage) Pop() interface{} {
	prev := s.data
	n := len(prev)
	item := prev[n-1]
	item.index = -1
	s.data = prev[:n-1]
	return item
}
