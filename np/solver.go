package np

import (
	"math"
)

type Solver struct {
	Question [9][9]int
	Solving  [9][9]int
	board    [9][9][]int
	History  [][9][9]int
}

func NewSolver(txt string) Solver {
	board := NewBoard(txt)
	board.solve()
	return Solver{
		Question: board.Question,
		Solving:  board.Solving,
		board:    board.board,
		History:  board.History,
	}
}

func (solver *Solver) Solve(x int, y int, history bool) bool {
	if y > 8 {
		return true
	} else if solver.Solving[y][x] != 0 {
		if x == 8 {
			if solver.Solve(0, y+1, history) {
				return true
			}
		} else {
			if solver.Solve(x+1, y, history) {
				return true
			}
		}
	} else {
		for _, v := range solver.board[y][x] {
			if solver.check(x, y, v) {
				solver.Solving[y][x] = v
				if history {
					solver.History = append(solver.History, solver.Solving)
				}
				if x == 8 {
					if solver.Solve(0, y+1, history) {
						return true
					}
				} else {
					if solver.Solve(x+1, y, history) {
						return true
					}
				}
			}
		}
		solver.Solving[y][x] = 0
		return false
	}
	return false
}

func (solver *Solver) check(x int, y int, i int) bool {
	return solver.row(y, i) && solver.col(x, i) && solver.block(x, y, i)
}

func (solver *Solver) row(y int, i int) bool {
	var ok = true
	for x := 0; x < 9; x++ {
		if solver.Solving[y][x] == i {
			ok = false
		}
	}
	return ok
}

func (solver *Solver) col(x int, i int) bool {
	var ok = true
	for y := 0; y < 9; y++ {
		if solver.Solving[y][x] == i {
			ok = false
		}
	}
	return ok
}

func (solver *Solver) block(x int, y int, i int) bool {
	var ok = true
	baseY := int(math.Floor(float64(y/3))) * 3
	baseX := int(math.Floor(float64(x/3))) * 3
	for _y := baseY; _y < baseY+3; _y++ {
		for _x := baseX; _x < baseX+3; _x++ {
			if solver.Solving[_y][_x] == i {
				ok = false
			}
		}
	}
	return ok
}
