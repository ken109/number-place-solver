package np

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Question      [9][9]int
	Solving       [9][9]int
	beforeSolving [9][9]int
	board         [9][9][]int
	History       [][9][9]int
}

func NewBoard(txt string) Board {
	board := Board{
		Question:      [9][9]int{},
		Solving:       [9][9]int{},
		beforeSolving: [9][9]int{},
		board:         [9][9][]int{},
		History:       nil,
	}
	board.readQuestion(txt)
	return board
}

func (board *Board) readQuestion(txt string) {
	f, err := os.Open(txt)
	if err != nil {
		color.Red("ファイルを開けませんでした。")
		os.Exit(1)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if len(strings.Split(strings.TrimRight(string(b), "\n"), "\n")) != 9 {
		formatErr()
	}
	for i, row := range strings.Split(strings.TrimRight(string(b), "\n"), "\n") {
		if len(strings.Split(row, "")) != 9 {
			formatErr()
		}
		for j, col := range strings.Split(row, "") {
			if board.Question[i][j], err = strconv.Atoi(col); err != nil {
				formatErr()
			}
		}
	}
	board.Solving = board.Question
}

func formatErr() {
	color.Red("問題のフォーマットが正しくありません。")
	fmt.Println("Example: \n023000000\n070000000\n005000010\n201079340\n084100002\n030004000\n300000001\n010800065\n700000090")
	os.Exit(1)
}

func (board *Board) solve() {
	double(func(y int, x int) {
		if board.Solving[y][x] == 0 {
			board.board[y][x] = board.check(x, y)
		} else {
			board.board[y][x] = []int{board.Solving[y][x]}
		}
	})

	board.beforeSolving = board.Solving
	double(func(y int, x int) {
		if len(board.board[y][x]) == 1 {
			board.Solving[y][x] = board.board[y][x][0]
		} else {
			board.Solving[y][x] = 0
		}
	})

	if board.beforeSolving != board.Solving {
		board.History = append(board.History, board.beforeSolving)
		board.solve()
	}
}

func (board *Board) check(x int, y int) []int {
	var ok []int
	var bad = set(sliceMerge(board.row(y), board.col(x), board.block(x, y)))
	for i := 1; i <= 9; i++ {
		if c, _ := search(bad, i); !c {
			ok = append(ok, i)
		}
	}
	return ok
}

func (board *Board) row(y int) []int {
	var rows []int
	for x := 0; x < 9; x++ {
		if board.Solving[y][x] != 0 {
			rows = append(rows, board.Solving[y][x])
		}
	}
	return rows
}

func (board *Board) col(x int) []int {
	var cols []int
	for y := 0; y < 9; y++ {
		if board.Solving[y][x] != 0 {
			cols = append(cols, board.Solving[y][x])
		}
	}
	return cols
}

func (board *Board) block(x int, y int) []int {
	var blocks []int
	baseY := int(math.Floor(float64(y/3))) * 3
	baseX := int(math.Floor(float64(x/3))) * 3
	for _y := baseY; _y < baseY+3; _y++ {
		for _x := baseX; _x < baseX+3; _x++ {
			if board.Solving[_y][_x] != 0 {
				blocks = append(blocks, board.Solving[_y][_x])
			}
		}
	}
	return blocks
}
