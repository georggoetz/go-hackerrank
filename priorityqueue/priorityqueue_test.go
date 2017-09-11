package priorityqueue

import "testing"

type minprio int

func (prio minprio) Less(x interface{}) bool {
	return prio < x.(minprio)
}

type maxprio int

func (prio maxprio) Less(x interface{}) bool {
	return prio > x.(maxprio)
}

func checkLength(t *testing.T, pq *PriorityQueue, len int) {
	if n := pq.Len(); n != len {
		t.Errorf("Len() = %d, want %d", n, len)
	}
}

// checkPointers repeatedly pops items from the priority queue and checks if
// the exprected items are returned. is contains the items in the order they
//  are expected to pop out of the priority queue.
func checkPQPointers(t *testing.T, pq *PriorityQueue, is []*Item) {
	// Check length
	checkLength(t, pq, len(is))

	// Compare pointers
	for count, item := range is {
		top := pq.Top()
		pop := pq.Pop()
		if item != pop {
			t.Errorf("Pop() times(%d) = %p, want %p", count+1, pop, item)
		}
		if top != pop {
			t.Errorf("Top() = %p, want %p", top, pop)
		}
	}

	// Priority queue must now be empty
	checkLength(t, pq, 0)

	// Check that Pop() returns nil when empty
	if pop := pq.Pop(); pop != nil {
		t.Errorf("Pop() = %p, want %p", pop, nil)
	}

	// Check that Top() returns nil when empty
	if top := pq.Top(); top != nil {
		t.Errorf("Top() = %p, want %p", top, nil)
	}
}

func TestPriorityQueue(t *testing.T) {
	// Empty priority queue
	pq := New(0)
	checkPQPointers(t, pq, []*Item{})

	// Check Pop() returning the item with min priority
	pq.Init(3)
	i0 := pq.Push(nil, minprio(2))
	i1 := pq.Push(nil, minprio(1))
	i2 := pq.Push(nil, minprio(3))
	checkPQPointers(t, pq, []*Item{i1, i0, i2})

	// Check Pop() returning the item with max priority
	pq.Init(3)
	i0 = pq.Push(nil, maxprio(2))
	i1 = pq.Push(nil, maxprio(1))
	i2 = pq.Push(nil, maxprio(3))
	checkPQPointers(t, pq, []*Item{i2, i0, i1})

	// Check Fix()
	pq.Init(3)
	i0 = pq.Push(nil, minprio(2))
	i1 = pq.Push(nil, minprio(1))
	i2 = pq.Push(nil, minprio(3))
	pq.Fix(i1, minprio(4))
	checkPQPointers(t, pq, []*Item{i0, i2, i1})

	// Check Remove()
	pq.Init(3)
	i0 = pq.Push(nil, minprio(2))
	i1 = pq.Push(nil, minprio(1))
	i2 = pq.Push(nil, minprio(3))
	pq.Remove(i1)
	checkPQPointers(t, pq, []*Item{i0, i2})
}
