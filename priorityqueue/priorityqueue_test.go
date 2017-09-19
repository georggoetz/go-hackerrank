package priorityqueue

import "testing"

type minComparer struct{}

func (c minComparer) Less(x, y interface{}) bool {
	return x.(int) < y.(int)
}

type maxComparer struct{}

func (c maxComparer) Less(x, y interface{}) bool {
	return x.(int) > y.(int)
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
	pq := New(minComparer{})
	checkPQPointers(t, pq, []*Item{})

	// Check Pop() returning the item with min priority
	pq.Init(minComparer{})
	i0 := pq.Push(nil, 2)
	i1 := pq.Push(nil, 1)
	i2 := pq.Push(nil, 3)
	checkPQPointers(t, pq, []*Item{i1, i0, i2})

	// Check Pop() returning the item with max priority
	pq.Init(maxComparer{})
	i0 = pq.Push(nil, 2)
	i1 = pq.Push(nil, 1)
	i2 = pq.Push(nil, 3)
	checkPQPointers(t, pq, []*Item{i2, i0, i1})

	// Check Fix()
	pq.Init(minComparer{})
	i0 = pq.Push(nil, 2)
	i1 = pq.Push(nil, 1)
	i2 = pq.Push(nil, 3)
	pq.Fix(i1, 4)
	checkPQPointers(t, pq, []*Item{i0, i2, i1})

	// Check Remove()
	pq.Init(minComparer{})
	i0 = pq.Push(nil, 2)
	i1 = pq.Push(nil, 1)
	i2 = pq.Push(nil, 3)
	pq.Remove(i1)
	checkPQPointers(t, pq, []*Item{i0, i2})
}
