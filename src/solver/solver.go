package solver

import (
    "grid"
    "errors"
)

func Solve(g *[9][9]int8) *[9][9]int8 {
    initialState := NewSolverState(0, 0, grid.EMPTY_VALUE, g)

    return step(initialState)
}

func step(state *SolverState) *[9][9]int8 {
    var solution *[9][9]int8

    nextX, nextY, err := getNextCell(state.X, state.Y)

    // current cell is filled
    if state.Grid[state.X][state.Y] != grid.EMPTY_VALUE {
        // last cell, return solution
        if err != nil {
            return &state.Grid
        }

        // step forward
        newState := NewSolverState(nextX, nextY, grid.EMPTY_VALUE, &state.Grid)

        solution = step(newState)
        if solution != nil {
            return solution
        }

        return solution
    }

    // current cell is available
    var valueToTest int8 = grid.MIN_VALUE
    if state.Value != grid.EMPTY_VALUE {
        state.Value += 1
        valueToTest = state.Value
    } else {
        state.Value = valueToTest
    }

    for val := valueToTest; val <= grid.MAX_VALUE; val++ {
        if isValueValidForCell(val, state.X, state.Y, &state.Grid) {
            // step forward
            newState := NewSolverState(nextX, nextY, grid.EMPTY_VALUE, &state.Grid)
            newState.Grid[state.X][state.Y] = val

            // last cell, return solution
            if err != nil {
                return &newState.Grid
            }

            solution = step(newState)
            if solution != nil {
                return solution
            }
        }
    }

    return nil
}

func isValueValidForCell(value, x, y int8, g *[9][9]int8) bool {
    inRow := grid.Row{x}.Contains(value, g)
    inLine := grid.Line{y}.Contains(value, g)
    inSector := grid.Sector{(x - (x % 3)) / 3, (y - (y % 3)) / 3}.Contains(value, g)
    return !inRow && !inLine && !inSector
}

func getNextCell(currX, currY int8) (int8, int8, error) {
    if currX == grid.GRID_SIZE - 1 && currY == grid.GRID_SIZE -1 {
        return grid.EMPTY_VALUE, grid.EMPTY_VALUE, errors.New("Last cell reached")
    }

    if currX < 8 {
        return currX + 1, currY, nil
    } else {
        return 0, currY + 1, nil
    }
}