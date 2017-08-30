package pacman

var directions = [4]node{
	node{row: -1, col: 0},
	node{row: 0, col: -1},
	node{row: 0, col: 1},
	node{row: 1, col: 0}}

type grid [][]rune

func (g grid) rows() int {
	return len(g)
}

func (g grid) cols() int {
	if len(g) > 0 {
		return len(g[0])
	}
	return 0
}

func (g grid) neighbors(x node) []node {
	var nextRow, nextCol int
	n := make([]node, 0, 4)
	for _, direction := range directions {
		nextRow, nextCol = x.row+direction.row, x.col+direction.col
		if nextRow < 0 || nextRow >= g.rows() || nextCol < 0 || nextCol >= g.cols() ||
			(g[nextRow][nextCol] != '.' && g[nextRow][nextCol] != '-') {
			continue
		}
		n = append(n, newNode(nextRow, nextCol))
	}
	return n
}
