package rbtree

type color bool

const (
	red, black color = true, false
)

type Interface interface {
	Less(interface{}) bool
}

type Tree struct {
	root *Node
	size int
}

func (t *Tree) Size() int {
	return t.size
}

type Node struct {
	value  Interface
	color  color
	parent *Node
	left   *Node
	right  *Node
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Search(v Interface) *Node {
	x := n
	for x != nil && (v.Less(x.value) || x.value.Less(v)) {
		if v.Less(x.value) {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (n *Node) Minimum() *Node {
	x := n
	for x.left != nil {
		x = x.left
	}
	return x
}

func (n *Node) Maximum() *Node {
	x := n
	for x.right != nil {
		x = x.right
	}
	return x
}

func (n *Node) Successor() *Node {
	if n.right != nil {
		return n.right.Minimum()
	}
	x := n
	y := n.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

func (n *Node) Predecessor() *Node {
	if n.left != nil {
		return n.left.Maximum()
	}
	x := n
	y := n.parent
	for y != nil && x == y.left {
		x = y
		y = y.parent
	}
	return y
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Search(v Interface) *Node {
	if t.root == nil {
		return nil
	}
	return t.root.Search(v)
}

func (t *Tree) Insert(v Interface) *Node {
	x, y, z := t.root, (*Node)(nil), &Node{value: v, color: red}
	for x != nil {
		y = x
		if z.value.Less(x.value) {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		t.root = z
	} else if z.value.Less(y.value) {
		y.left = z
	} else {
		y.right = z
	}
	t.insertFixup(z)
	t.size += 1
	return z
}

func (t *Tree) insertFixup(z *Node) {
	var y *Node
	for z.parent != nil && z.parent.parent != nil && z.parent.color == red {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else if z == z.parent.right {
				z = z.parent
				t.leftRotate(z)
			} else {
				z.parent.color = black
				if z.parent.parent != nil {
					z.parent.parent.color = red
					t.rightRotate(z.parent.parent)
				}
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else if z == z.parent.left {
				z = z.parent
				t.rightRotate(z)
			} else {
				z.parent.color = black
				if z.parent.parent != nil {
					z.parent.parent.color = red
					t.leftRotate(z.parent.parent)
				}
			}
		}
	}
	t.root.color = black
}

func (t *Tree) leftRotate(x *Node) {
	if x.right == nil {
		return
	}
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}
	y.left = x
	x.parent = y
}

func (t *Tree) rightRotate(x *Node) {
	if x.left == nil {
		return
	}
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else {
		if x == x.parent.right {
			x.parent.right = y
		} else {
			x.parent.left = y
		}
	}
	y.right = x
	x.parent = y
}

func BlackHeight(n *Node) int {
	if n == nil {
		return 0
	}
	l := BlackHeight(n.left)
	r := BlackHeight(n.right)
	a := 0
	if n.color == black {
		a = 1
	}

	if l == -1 || r == -1 || l != r {
		return -1
	} else {
		return l + a
	}
}
