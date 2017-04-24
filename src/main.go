package main

import (
    "fmt"
    "solver"
)

var solvedGrid = [9][9]int8 {
        { 5, 6, 1, 8, 4, 7, 9, 2, 3 },
        { 3, 7, 9, 5, 2, 1, 6, 8, 4 },
        { 4, 2, 8, 9, 6, 3, 1, 7, 5 },
        { 6, 1, 3, 7, 8, 9, 5, 4, 2 },
        { 7, 9, 4, 6, 5, 2, 3, 1, 8 },
        { 8, 5, 2, 1, 3, 4, 7, 9, 6 },
        { 9, 3, 5, 4, 7, 8, 2, 6, 1 },
        { 1, 4, 6, 2, 9, 5, 8, 3, 7 },
        { 2, 8, 7, 3, 1, 6, 4, 5, 9 },
    }

var unolvedGrid = [9][9]int8 {
        { 5, 6, -1, 8, 4, 7, -1, -1, -1 },
        { 3, -1, 9, -1, -1, -1, 6, -1, -1 },
        { -1, -1, 8, -1, -1, -1, -1, -1, 5 },
        { -1, 1, -1, -1, 8, -1, -1, 4, -1 },
        { 7, 9, -1, 6, -1, 2, -1, 1, 8 },
        { -1, 5, -1, -1, 3, -1, -1, 9, -1 },
        { -1, -1, -1, -1, -1, -1, 2, -1, -1 },
        { -1, -1, 6, -1, -1, -1, 8, -1, 7 },
        { -1, -1, -1, 3, 1, 6, -1, 5, 9 },
    }

func main() {
    solution := solver.Solve(&unolvedGrid)

    if solution != nil {
        if validateSolution(solution) {
            fmt.Println("SUCCESS")
        } else {
            fmt.Println("FAIL")
        }

        for y := 0; y < 9; y++ {
            for x := 0; x < 9; x++ {
                fmt.Printf("%d ", solution[x][y])
            }
            fmt.Println()
        }

    } else {
        fmt.Println("FAIL")
    }
}

func validateSolution(solution *[9][9]int8) bool {
    for y := 0; y < 9; y++ {
        for x := 0; x < 9; x++ {
            if solvedGrid[x][y] != solution[x][y] {
                return false
            }
        }
    }

    return true
}