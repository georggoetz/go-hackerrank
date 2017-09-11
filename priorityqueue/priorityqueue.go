package priorityqueue

import "container/heap"

// PriorityComparer must be implemented for any type of priority.
type PriorityComparer interface {
	// Less decides if it is a min- or max-priority queue.
	Less(interface{}) bool
}

// Item is an element of the priority queue.
type Item struct {
	Value    interface{}
	priority PriorityComparer
	index    int
}

// PriorityQueue is similar to a stack or queue, but additionally each element
// has a priority associated with it. An element with high priority is served
// before elements with lower priority.
type PriorityQueue struct {
	heap storage
}

// Init initializes or clears the priority queue. The initial length is 0 the
// initial capacity is set to cap.
func (pq *PriorityQueue) Init(cap int) *PriorityQueue {
	pq.heap = make(storage, 0, cap)
	return pq
}

// New returns an initialized priority queue with a capacity for cap elements.
func New(cap int) *PriorityQueue {
	return new(PriorityQueue).Init(cap)
}

// Len returns the length.
func (pq *PriorityQueue) Len() int {
	return pq.heap.Len()
}

// Push pushes a new item with the given value and priority into the priority
// queue and returns the item.
func (pq *PriorityQueue) Push(val interface{}, prio PriorityComparer) *Item {
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
		return pq.heap[0]
	}
	return nil
}

// Fix re-establishes the heap ordering after the priority of the given item
// has changed. Equivalent to, but less expensive, to calling Remove() followed
// by Push().
func (pq *PriorityQueue) Fix(i *Item, prio PriorityComparer) {
	i.priority = prio
	heap.Fix(&pq.heap, i.index)
}

// Remove removes the given item.
func (pq *PriorityQueue) Remove(i *Item) *Item {
	return heap.Remove(&pq.heap, i.index).(*Item)
}

// See examples from container/heap.
type storage []*Item

func (h storage) Len() int {
	return len(h)
}

func (h storage) Less(i, j int) bool {
	return h[i].priority.Less(h[j].priority)
}

func (h storage) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *storage) Push(x interface{}) {
	n := len(*h)
	item := x.(*Item)
	item.index = n
	*h = append(*h, item)
}

func (h *storage) Pop() interface{} {
	prev := *h
	n := len(prev)
	item := prev[n-1]
	item.index = -1
	*h = prev[:n-1]
	return item
}
