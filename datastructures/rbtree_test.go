package rbtree

import "fmt"

type key int

func (n key) Less(v interface{}) bool {
	value, _ := v.(key)
	return n < value
}

func ExampleTree_Insert() {
	t := NewTree()
	t.Insert(key(1407482458))
	t.Insert(key(923960288))
	t.Insert(key(225310218))
	t.Insert(key(1166620750))
	t.Insert(key(458607644))
	t.Insert(key(1403830842))
	t.Insert(key(1164433078))
	t.Insert(key(1273045782))
	t.Insert(key(1856817205))
	t.Insert(key(198887065))
	t.Insert(key(262719564))
	t.Insert(key(270288546))
	t.Insert(key(1713131318))
	t.Insert(key(440430166))
	t.Insert(key(272333804))
	t.Insert(key(824337406))
}

func ExampleTree_Search() {
	t := NewTree()
	t.Insert(key(1))
	t.Insert(key(2))
	t.Insert(key(3))
	t.Insert(key(4))
	t.Insert(key(5))
	t.Insert(key(6))
	n := t.Search(key(4))
	if n != nil {
		fmt.Println(n.Value())
	}
	n = t.Search(key(2))
	if n != nil {
		fmt.Println(n.Value())
	}
	n = t.Search(key(5))
	if n != nil {
		fmt.Println(n.Value())
	}
	// Output:
	// 4
	// 2
	// 5
}

func ExampleNode_Maximum() {
	t := NewTree()
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

	fmt.Println(t.root.Maximum().Value())
	// Output:
	// 20
}

func ExampleNode_Minimum() {
	var t = NewTree()
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

	fmt.Println(t.root.Minimum().Value())
	// Output:
	// 2
}

func ExampleNode_Successor() {
	var t = NewTree()
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

	n := t.Search(key(13))
	fmt.Println(n.Successor().Value())
	// Output:
	// 15
}

func ExampleNode_Predecessor() {
	var t = NewTree()
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

	n := t.Search(key(13))
	fmt.Println(n.Predecessor().Value())
	// Output:
	// 9
}
