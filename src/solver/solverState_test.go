package solver

import (
    "testing"
    "grid"
)

func TestNewSolverStateCopiesTheValueOfItsGridArgument(t *testing.T) {
    var g [9][9]int8

    var sample int8 = 1
    for x := 0; int8(x) < grid.GRID_SIZE; x++ {
        for y := 0; int8(y) < grid.GRID_SIZE; y++ {
            g[x][y] = sample
            sample++
        }
    }

    state := NewSolverState(1, 3, &g)

    state.Grid[3][3] = grid.EMPTY_VALUE

    if state.Grid[3][3] == g[3][3] {
        t.Errorf("Expected SolverState to have a local copy of the grid")
    }
}