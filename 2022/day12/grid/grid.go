package grid

import (
	"container/list"
	"fmt"
	"math"
	"strings"
)

type Grid struct {
	heights [][]rune
	si, sj  int
}

type listnode struct {
	char rune
	i, j int
}

func NewGrid(in string) *Grid {
	lines := strings.Split(in, "\n")
	grid := new(Grid)
	grid.heights = make([][]rune, len(lines))
	for i, line := range lines {
		grid.heights[i] = make([]rune, len(line))
		for j, char := range line {
			if char == 'S' {
				grid.si = i
				grid.sj = j
				grid.heights[i][j] = 'a' - 1
			} else if char == 'E' {
				grid.heights[i][j] = 'z' + 1
			} else {
				grid.heights[i][j] = char

			}
		}
	}
	return grid
}

func (g Grid) setupSPL() ([][]int, *list.List) {
	dists := make([][]int, len(g.heights))
	for i := 0; i < len(dists); i++ {
		dists[i] = make([]int, len(g.heights[i]))
		for j := 0; j < len(dists[i]); j++ {
			dists[i][j] = math.MaxInt
		}
	}

	dists[g.si][g.sj] = 0

	queue := list.New()

	for i := 0; i < len(dists); i++ {
		for j := 0; j < len(dists[i]); j++ {
			queue.PushFront(listnode{g.heights[i][j], i, j})
		}
	}

	return dists, queue
}

func (g Grid) setupSPLNoStart() ([][]int, *list.List) {
	dists := make([][]int, len(g.heights))
	for i := 0; i < len(dists); i++ {
		dists[i] = make([]int, len(g.heights[i]))
		for j := 0; j < len(dists[i]); j++ {
			if g.heights[i][j] > 'a' {
				dists[i][j] = math.MaxInt
			}
		}
	}

	// time to find minimum is O(n), a priority queue would
	// cost O(logn)
	queue := list.New()

	for i := 0; i < len(dists); i++ {
		for j := 0; j < len(dists[i]); j++ {
			queue.PushFront(listnode{g.heights[i][j], i, j})
		}
	}
	return dists, queue
}

func (g Grid) dijkstra(dists [][]int, queue *list.List) int {
	var cw listnode
	var ei, ej int
	for queue.Len() != 0 {
		cw = queue.Remove(minEl(queue, dists)).(listnode)
		if cw.char == 'z'+1 {
			ei = cw.i
			ej = cw.j
		}
		if dists[cw.i][cw.j] == math.MaxInt {
			break
		}
		if cw.i > 0 && g.heights[cw.i-1][cw.j] <= cw.char+1 {
			updateDists(dists, cw.i-1, cw.j, dists[cw.i][cw.j]+1)
		}
		if cw.j > 0 && g.heights[cw.i][cw.j-1] <= cw.char+1 {
			updateDists(dists, cw.i, cw.j-1, dists[cw.i][cw.j]+1)
		}
		if cw.i < len(g.heights)-1 && g.heights[cw.i+1][cw.j] <= cw.char+1 {
			updateDists(dists, cw.i+1, cw.j, dists[cw.i][cw.j]+1)
		}
		if cw.j < len(g.heights[cw.i])-1 && g.heights[cw.i][cw.j+1] <= cw.char+1 {
			updateDists(dists, cw.i, cw.j+1, dists[cw.i][cw.j]+1)
		}
	}
	return dists[ei][ej]
}

func (g Grid) ShortestPathLen() int {
	dists, queue := g.setupSPL()
	return g.dijkstra(dists, queue)
}

func (g Grid) ShortestPathLenNoStart() int {
	dists, queue := g.setupSPLNoStart()
	return g.dijkstra(dists, queue)
}

func (l listnode) String() string {
	var s string
	s += fmt.Sprintf("coord: %d %d\n", l.i, l.j)
	s += fmt.Sprintf("char: %c", l.char)
	return s
}

func minEl(l *list.List, dists [][]int) *list.Element {
	min := 10000
	var ret *list.Element = l.Front()
	for e := l.Front(); e != nil; e = e.Next() {
		node := e.Value.(listnode)
		if dists[node.i][node.j] <= min {
			min = dists[node.i][node.j]
			ret = e
		}
	}
	return ret
}

func updateDists(dists [][]int, x, y, value int) {
	if dists[x][y] > value {
		dists[x][y] = value
	}
}
