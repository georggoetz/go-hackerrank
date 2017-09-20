// https://www.hackerrank.com/challenges/queries-with-fixed-length

package fixedlengthqueries

import (
	"container/list"
	"fmt"
	"io"

	"github.com/georggoetz/hackerrank/math"
)

// Solve calculates a solution in O(n*k). This is not good enough to complete
// the challenge. Several tests time out.
func Solve(a []int, k int) int {
	len := len(a)
	m := make([]int, len-k+1)
	for i, j := 0, k; j <= len; i, j = i+1, j+1 {
		m[i] = math.MaxIntSlice(a[i:j])
	}
	return math.MinIntSlice(m)
}

// Solve2 is able to find a solution in O(n).
// Explanation: A sliding window of size k is moved over the slice a. A double
// ended queue stores all indices of elements that are larger then the elements
// to the left of it, such that the queue contains elements in the order they
// were added and in decreasing order of values.
// Because every element is only added and then removed once from the queue
// the number of operations is 2n.
func Solve2(a []int, k int) int {
	len := len(a)
	q := list.New()
	m := make([]int, 0, len-k+1)
	i := 0
	// Process the first window
	for ; i < k; i++ {
		// Remove previously added elements left of the current element When
		// less or equal
		for e := q.Back(); e != nil && a[e.Value.(int)] <= a[i]; e = q.Back() {
			q.Remove(e)
		}
		q.PushBack(i)
	}
	// Process the remaining elements
	for ; i < len; i++ {
		// Report the largest element of the previous window. It is at the front.
		m = append(m, a[q.Front().Value.(int)])
		// Again remove previously added elements left of the...
		for e := q.Back(); e != nil && a[e.Value.(int)] <= a[i]; e = q.Back() {
			q.Remove(e)
		}
		// Remove elements that do not belong to the current window.
		for e := q.Front(); e != nil && e.Value.(int) <= i-k; e = q.Front() {
			q.Remove(e)
		}
		q.PushBack(i)
	}
	// Add the largest value of the last window.
	m = append(m, a[q.Front().Value.(int)])
	return math.MinIntSlice(m)
}

func read(r io.Reader) ([]int, []int) {
	var n, q int
	fmt.Fscanf(r, "%d %d\n", &n, &q)
	a := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(r, "%d", &a[i])
	}
	fmt.Fscanf(r, "%d\n", &a[n-1])
	k := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscanf(r, "%d\n", &k[i])
	}
	return a, k
}
