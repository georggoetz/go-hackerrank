// http://www.hackerrank.com/challenges/n-puzzle

package npuzzle

import (
	"bufio"
	"errors"
	"io"
	"math"
	"strconv"

	moremath "github.com/georggoetz/hackerrank/math"
	"github.com/georggoetz/hackerrank/priorityqueue"
	"github.com/georggoetz/hackerrank/rbtree"
)

type tiles [][]int
type priority struct{}

// Move denotes the movements of a tile of the puzzle.
type Move int

// Puzzle represents a sliding puzzle.
type Puzzle struct {
	n, g   int
	open   *priorityqueue.Item
	closed bool
	prev   *Puzzle
	move   Move
	tiles  tiles
}

var (
	moves = map[Move]struct {
		row, col int
	}{
		up:    {row: -1, col: 0},
		down:  {row: 1, col: 0},
		left:  {row: 0, col: -1},
		right: {row: 0, col: 1}}
)

const (
	up Move = iota
	down
	left
	right
)

// NewPuzzle returns a new puzzle of the given size and tiles.
func NewPuzzle(n int, tiles []int) *Puzzle {
	return new(Puzzle).init(n, tiles)
}

// Less compares two puzzles. On the occurence of the first difference it
// returns true if the number on the tile of this puzzle is less than on the
// other one. Otherwise, it returns false.
func (p *Puzzle) Less(x interface{}) bool {
	q := x.(*Puzzle)
	for r := range p.tiles {
		for c := range p.tiles[r] {
			if v, w := p.tiles[r][c], q.tiles[r][c]; v != w {
				return v < w
			}
		}
	}
	return false
}

// ReadPuzzle returns a new puzzle created from the spcified reader. See
// hackerrank for input details.
func ReadPuzzle(r io.Reader) *Puzzle {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	len := n * n
	tiles := make([]int, len)
	for i := 0; i < len && scanner.Scan(); i++ {
		tiles[i], _ = strconv.Atoi(scanner.Text())
	}
	return NewPuzzle(n, tiles)
}

// Solve calculates the moves required to solve the puzzle. It applies the
// A* algorithm to determine an optimal solution.
func (p *Puzzle) Solve() ([]Move, error) {
	db := rbtree.New()
	q := priorityqueue.New(priority{})
	p.g = 0
	p.open = q.Push(p, p.manhattan())
	db.Insert(p)
	for q.Len() > 0 {
		u := q.Pop().Value.(*Puzzle)
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
	return []Move{}, errors.New("path not found")
}

// String returns a move as string.
func (m Move) String() (s string) {
	switch m {
	case up:
		s = "UP"
	case down:
		s = "DOWN"
	case left:
		s = "LEFT"
	case right:
		s = "RIGHT"
	}
	return
}

func (p priority) Less(x, y interface{}) bool {
	return x.(int) < y.(int)
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

func (p *Puzzle) init(n int, tiles []int) *Puzzle {
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

func (p *Puzzle) copy() *Puzzle {
	n := p.n
	q := NewPuzzle(n, make([]int, n*n))
	for r := range q.tiles {
		for c := range q.tiles[r] {
			q.tiles[r][c] = p.tiles[r][c]
		}
	}
	return q
}

func (p *Puzzle) manhattan() int {
	n, d := p.n, 0
	for r := range p.tiles {
		for c := range p.tiles[r] {
			if v := p.tiles[r][c]; v != 0 {
				d += moremath.AbsInt(r-(v/n)) + moremath.AbsInt(c-(v%n))
			}
		}
	}
	return d
}

func (p *Puzzle) goal() bool {
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

func (p *Puzzle) next(db *rbtree.Tree) map[Move]*Puzzle {
	a := make(map[Move]*Puzzle)
	if r, c, ok := p.tiles.indexOf(0); ok {
		for m, v := range moves {
			q := p.copy()
			if q.tiles.swap(r, c, v.row, v.col) {
				if rec := db.Search(q); rec != nil {
					q = rec.Value.(*Puzzle)
				}
				a[m] = q
			}
		}
	}
	return a
}

func (p *Puzzle) path() []Move {
	a := make([]Move, 0)
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
