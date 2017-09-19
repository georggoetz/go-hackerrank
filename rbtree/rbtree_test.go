package rbtree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type key int

func (n key) Less(v interface{}) bool {
	value, _ := v.(key)
	return n < value
}

func blackHeight(n *Node) int {
	if n == nil {
		return 0
	}
	l := blackHeight(n.left)
	r := blackHeight(n.right)
	a := 0
	if n.color == black {
		a = 1
	}

	if l == -1 || r == -1 || l != r {
		return -1
	}
	return l + a
}

func TestTree_Insert(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	r := New()
	var v int
	for i := 0; i < 10000; i++ {
		v = rand.Int()
		r.Insert(key(v))
		if blackHeight(r.root) == -1 {
			t.Error("invalid black height")
		}
	}
}

func testTree() *Tree {
	t := New()
	t.Insert(key(2))
	t.Insert(key(3))
	t.Insert(key(4))
	t.Insert(key(6))
	t.Insert(key(7))
	t.Insert(key(9))
	t.Insert(key(13))
	t.Insert(key(15))
	t.Insert(key(17))
	t.Insert(key(18))
	t.Insert(key(20))
	return t
}

func ExampleTree_Search() {
	t := testTree()
	fmt.Println(t.Search(key(4)).Value)
	fmt.Println(t.Search(key(2)).Value)
	fmt.Println(t.Search(key(5)))
	// Output:
	// 4
	// 2
	// <nil>
}

func ExampleNode_Maximum() {
	t := testTree()
	fmt.Println(t.root.Maximum().Value)
	// Output:
	// 20
}

func ExampleNode_Minimum() {
	t := testTree()
	fmt.Println(t.root.Minimum().Value)
	// Output:
	// 2
}

func ExampleNode_Successor() {
	t := testTree()
	n := t.Search(key(13))
	fmt.Println(n.Successor().Value)
	// Output:
	// 15
}

func ExampleNode_Predecessor() {
	t := testTree()
	n := t.Search(key(13))
	fmt.Println(n.Predecessor().Value)
	// Output:
	// 9
}
