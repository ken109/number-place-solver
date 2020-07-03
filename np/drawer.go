package np

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"time"
)

func Draw(txt string, delay int) {
	solver := NewSolver(txt)
	if delay <= 0 {
		go solver.Solve(0, 0, false)
	} else {
		go solver.Solve(0, 0, true)
	}
	var before [9][9]int
	for true {
		if delay <= 0 {
			npPrint(solver.Solving, solver.Question)
			time.Sleep(time.Microsecond * 5)
			if before != solver.Solving {
				reset()
				before = solver.Solving
			} else {
				break
			}
		} else {
			for true {
				if len(solver.History) > 0 {
					break
				}
				time.Sleep(time.Millisecond)
			}
			npPrint(solver.History[0], solver.Question)
			if len(solver.History) > 1 {
				solver.History = solver.History[1:]
			}
			time.Sleep(time.Millisecond * time.Duration(delay))
			reset()
		}
	}
}

func npPrint(np [9][9]int, question [9][9]int) {
	for y, row := range np {
		for x, col := range row {
			if question[y][x] != 0 {
				green := color.New(color.FgGreen).PrintfFunc()
				green("%s ", strconv.Itoa(col))
			} else {
				fmt.Printf("%s ", strconv.Itoa(col))
			}
			if x == 2 || x == 5 {
				fmt.Printf("| ")
			}
		}
		if y == 2 || y == 5 {
			fmt.Printf("\n")
			for i := 0; i < 21; i++ {
				if i == 6 || i == 14 {
					fmt.Printf("+")
				} else {
					fmt.Printf("-")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func reset() {
	for i := 0; i < 11; i++ {
		fmt.Printf("\033[1A\033[2K\033[G")
	}
}
