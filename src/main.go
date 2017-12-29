package main

import (
    "fmt"
    "solver"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
    solution := solver.Solve(readSudoku())

    if solution != nil {
        for y := 0; y < 9; y++ {
            for x := 0; x < 9; x++ {
                fmt.Printf("%d ", solution[x][y])
            }
            fmt.Println()
        }
    }
}

func readSudoku() *[9][9]int8 {
	reader := bufio.NewReader(os.Stdin)
	var sudoku [9][9]int8
	
	lineNumber := 0
	
	for lineNumber < 9 {
		line, _ := reader.ReadString('\n')
		for i, c := range strings.Split(line, "") {
			num, _ := strconv.ParseInt(c, 10, 8)
			if i < 9 {
				if int8(num) < 1 {
					sudoku[i][int8(lineNumber)] = -1
				} else {
					sudoku[i][int8(lineNumber)] = int8(num)
				}
			}
		}
		lineNumber++
	}
	
    return &sudoku
}