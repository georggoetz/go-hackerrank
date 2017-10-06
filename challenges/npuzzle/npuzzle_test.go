package npuzzle

import (
	"fmt"
	"strings"
	"testing"
)

const sampleInput = `3
0
3
8
4
1
7
2
6
5`

func checkPuzzle(t *testing.T, p *Puzzle, tiles [][]int) {
	if len := len(p.tiles); len != p.n {
		t.Errorf("len(tiles) = %d, want %d", len, p.n)
	}
	for i := range p.tiles {
		if len := len(p.tiles[i]); len != p.n {
			t.Errorf("len(tiles[%d]) = %d, want %d", i, len, p.n)
		}
	}
	for i := range tiles {
		for j := range tiles[i] {
			if is, want := p.tiles[i][j], tiles[i][j]; is != want {
				t.Errorf("tiles[%d][%d] = %d, want %d", i, j, is, want)
			}
		}
	}
}

func TestPuzzle(t *testing.T) {
	p := ReadPuzzle(strings.NewReader(sampleInput))
	checkPuzzle(t, p, [][]int{{0, 3, 8}, {4, 1, 7}, {2, 6, 5}})
}

func checkManhattanDistance(t *testing.T, p *Puzzle, want int) {
	if is := p.manhattan(); is != want {
		t.Errorf("manhattanDistance() = %d, want %d", is, want)
	}
}

func TestManhattanDistance(t *testing.T) {
	p := NewPuzzle(2, []int{0, 1, 2, 3})
	checkManhattanDistance(t, p, 0)
	p.init(2, []int{0, 2, 1, 3})
	checkManhattanDistance(t, p, 4)
	p.init(2, []int{1, 0, 3, 2})
	checkManhattanDistance(t, p, 3)
	p.init(2, []int{3, 2, 1, 0})
	checkManhattanDistance(t, p, 6)
}

func checkGoal(t *testing.T, p *Puzzle, want bool) {
	if is := p.goal(); is != want {
		t.Errorf("goal() = %t, want %t", is, want)
	}
}

func TestGoal(t *testing.T) {
	p := NewPuzzle(3, []int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	checkGoal(t, p, true)
	p.init(2, []int{0, 1, 3, 2})
	checkGoal(t, p, false)
}

func checkIndexOf(t *testing.T, p *Puzzle, r, c, v int, ok bool) {
	if y, x, b := p.tiles.indexOf(v); y != r || x != c || b != ok {
		t.Errorf("indexOf(%d) = %d, %d, %t, want %d, %d, %t", v, y, x, b, r, c, ok)
	}
}

func TestIndexOf(t *testing.T) {
	p := NewPuzzle(3, []int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	checkIndexOf(t, p, 0, 0, 0, true)
	checkIndexOf(t, p, 0, 2, 2, true)
	checkIndexOf(t, p, 2, 2, 8, true)
	checkIndexOf(t, p, 3, 3, 9, false)
}

func printMoves(moves []Move) {
	fmt.Println(len(moves))
	for _, move := range moves {
		fmt.Println(move.String())
	}
}

func ExamplePuzzle_Solve() {
	moves, _ := ReadPuzzle(strings.NewReader(sampleInput)).Solve()
	printMoves(moves)
	// Output:
	// 20
	// DOWN
	// DOWN
	// RIGHT
	// UP
	// LEFT
	// UP
	// RIGHT
	// DOWN
	// DOWN
	// RIGHT
	// UP
	// UP
	// LEFT
	// DOWN
	// DOWN
	// RIGHT
	// UP
	// LEFT
	// LEFT
	// UP
}
