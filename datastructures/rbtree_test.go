package rbtree

import "fmt"

func ExampleTree_Insert() {
	t := NewTree()
	t.Insert(1407482458, nil)
	t.Insert(923960288, nil)
	t.Insert(225310218, nil)
	t.Insert(1166620750, nil)
	t.Insert(458607644, nil)
	t.Insert(1403830842, nil)
	t.Insert(1164433078, nil)
	t.Insert(1273045782, nil)
	t.Insert(1856817205, nil)
	t.Insert(198887065, nil)
	t.Insert(262719564, nil)
	t.Insert(270288546, nil)
	t.Insert(1713131318, nil)
	t.Insert(440430166, nil)
	t.Insert(272333804, nil)
	t.Insert(824337406, nil)

	fmt.Println(BlackHeight(t.root))
	// Output:
	// -1
}

func ExampleTree_Search() {
	t := NewTree()
	t.Insert(1, nil)
	t.Insert(2, nil)
	t.Insert(3, nil)
	t.Insert(4, nil)
	t.Insert(5, nil)
	t.Insert(6, nil)
	n := t.Search(4)
	if n != nil {
		fmt.Println(n.key)
	}
	n = t.Search(2)
	if n != nil {
		fmt.Println(n.key)
	}
	n = t.Search(5)
	if n != nil {
		fmt.Println(n.key)
	}
	// Output:
	// 4
	// 2
	// 5
}

func ExampleNode_Maximum() {
	t := NewTree()
	t.Insert(2, nil)
	t.Insert(3, nil)
	t.Insert(4, nil)
	t.Insert(6, nil)
	t.Insert(7, nil)
	t.Insert(9, nil)
	t.Insert(13, nil)
	t.Insert(15, nil)
	t.Insert(17, nil)
	t.Insert(18, nil)
	t.Insert(20, nil)

	fmt.Println(t.root.Maximum().key)
	// Output:
	// 20
}

func ExampleNode_Minimum() {
	var t = NewTree()
	t.Insert(2, nil)
	t.Insert(3, nil)
	t.Insert(4, nil)
	t.Insert(6, nil)
	t.Insert(7, nil)
	t.Insert(9, nil)
	t.Insert(13, nil)
	t.Insert(15, nil)
	t.Insert(17, nil)
	t.Insert(18, nil)
	t.Insert(20, nil)

	fmt.Println(t.root.Minimum().key)
	// Output:
	// 2
}

func ExampleNode_Successor() {
	var t = NewTree()
	t.Insert(2, nil)
	t.Insert(3, nil)
	t.Insert(4, nil)
	t.Insert(6, nil)
	t.Insert(7, nil)
	t.Insert(9, nil)
	t.Insert(13, nil)
	t.Insert(15, nil)
	t.Insert(17, nil)
	t.Insert(18, nil)
	t.Insert(20, nil)

	n := t.Search(13)
	fmt.Println(n.Successor().key)
	// Output:
	// 15
}

func ExampleNode_Predecessor() {
	var t = NewTree()
	t.Insert(2, nil)
	t.Insert(3, nil)
	t.Insert(4, nil)
	t.Insert(6, nil)
	t.Insert(7, nil)
	t.Insert(9, nil)
	t.Insert(13, nil)
	t.Insert(15, nil)
	t.Insert(17, nil)
	t.Insert(18, nil)
	t.Insert(20, nil)

	n := t.Search(13)
	fmt.Println(n.Predecessor().key)
	// Output:
	// 9
}
