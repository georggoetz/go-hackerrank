package npuzzle

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/georggoetz/hackerrank/priorityqueue"
	"github.com/georggoetz/hackerrank/rbtree"
)

type tiles [][]int
type move int
type priority struct{}

type puzzle struct {
	n, g   int
	open   *priorityqueue.Item
	closed bool
	prev   *puzzle
	move   move
	tiles  tiles
}

var (
	moves = map[move]struct {
		row, col int
	}{
		up:    {row: -1, col: 0},
		down:  {row: 1, col: 0},
		left:  {row: 0, col: -1},
		right: {row: 0, col: 1}}
)

const (
	up move = iota
	down
	left
	right
)

func (p priority) Less(x, y interface{}) bool {
	return x.(int) < y.(int)
}

func absi(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (t tiles) indexOf(v int) (int, int, bool) {
	var r, c int
	n := len(t)
loop:
	for r = 0; r < n; r++ {
		for c = 0; c < n; c++ {
			if t[r][c] == v {
				break loop
			}
		}
	}
	return r, c, r < n && c < n
}

func (t tiles) swap(r, c, dr, dc int) bool {
	n := len(t)
	if row, col := r+dr, c+dc; row >= 0 && col >= 0 && row < n && col < n {
		t[r][c], t[row][col] = t[row][col], t[r][c]
		return true
	}
	return false
}

func readPuzzle(r io.Reader) *puzzle {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	len := n * n
	tiles := make([]int, len)
	for i := 0; i < len && scanner.Scan(); i++ {
		tiles[i], _ = strconv.Atoi(scanner.Text())
	}
	return newPuzzle(n, tiles)
}

func printMoves(moves []move) {
	fmt.Println(len(moves))
	for _, move := range moves {
		switch move {
		case up:
			fmt.Println("UP")
		case down:
			fmt.Println("DOWN")
		case left:
			fmt.Println("LEFT")
		case right:
			fmt.Println("RIGHT")
		}
	}
}

func newPuzzle(n int, tiles []int) *puzzle {
	return new(puzzle).init(n, tiles)
}

func (p *puzzle) Less(x interface{}) bool {
	q := x.(*puzzle)
	for r := range p.tiles {
		for c := range p.tiles[r] {
			if v, w := p.tiles[r][c], q.tiles[r][c]; v != w {
				return v < w
			}
		}
	}
	return false
}

func (p *puzzle) init(n int, tiles []int) *puzzle {
	p.n = n
	p.open = nil
	p.prev = nil
	p.g = math.MaxInt32
	p.tiles = make([][]int, n)
	for r := range p.tiles {
		p.tiles[r] = make([]int, n)
		for c := range p.tiles[r] {
			p.tiles[r][c] = tiles[(r*n)+c]
		}
	}
	return p
}

func (p *puzzle) copy() *puzzle {
	n := p.n
	q := newPuzzle(n, make([]int, n*n))
	for r := range q.tiles {
		for c := range q.tiles[r] {
			q.tiles[r][c] = p.tiles[r][c]
		}
	}
	return q
}

func (p *puzzle) manhattan() int {
	n, d := p.n, 0
	for r := range p.tiles {
		for c := range p.tiles[r] {
			if v := p.tiles[r][c]; v != 0 {
				d += absi(r-(v/n)) + absi(c-(v%n))
			}
		}
	}
	return d
}

func (p *puzzle) goal() bool {
	n := p.n
	for r := range p.tiles {
		for c := range p.tiles[r] {
			if v := p.tiles[r][c]; v != (r*n)+c {
				return false
			}
		}
	}
	return true
}

func (p *puzzle) next(db *rbtree.Tree) map[move]*puzzle {
	a := make(map[move]*puzzle)
	if r, c, ok := p.tiles.indexOf(0); ok {
		for m, v := range moves {
			q := p.copy()
			if q.tiles.swap(r, c, v.row, v.col) {
				if rec := db.Search(q); rec != nil {
					q = rec.Value.(*puzzle)
				}
				a[m] = q
			}
		}
	}
	return a
}

func (p *puzzle) path() []move {
	a := make([]move, 0)
	x := p
	for x != nil {
		a = append(a, x.move)
		x = x.prev
	}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a[1:len(a)]
}

func (p *puzzle) solve() ([]move, error) {
	db := rbtree.NewTree()
	q := priorityqueue.New(priority{})
	p.g = 0
	p.open = q.Push(p, p.manhattan())
	db.Insert(p)
	for q.Len() > 0 {
		u := q.Pop().Value.(*puzzle)
		if u.goal() {
			return u.path(), nil
		}
		u.closed = true
		for k, v := range u.next(db) {
			if v.closed {
				continue
			}
			tentativeG := u.g + 1
			if v.open != nil && tentativeG >= v.g {
				continue
			}
			v.move = k
			v.prev = u
			v.g = tentativeG
			f := tentativeG + v.manhattan()
			if v.open != nil {
				q.Fix(v.open, f)
			} else {
				v.open = q.Push(v, f)
				db.Insert(v)
			}
		}
	}
	return []move{}, errors.New("path not found")
}
