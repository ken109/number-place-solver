package main

import (
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/ken109/sudoku-go/np"
)

func main() {
	if len(os.Args) < 3 {
		color.Red("Usage:   nump <問題ファイルのパス> <ディレイ(ms)>")
		color.Green("Example: nump question.txt 50")
		os.Exit(1)
	}

	delay, err := strconv.Atoi(os.Args[2])
	if err != nil {
		color.Red("ディレイの値が正しくありません。")
		os.Exit(1)
	}

	np.Draw(os.Args[1], delay)
}
