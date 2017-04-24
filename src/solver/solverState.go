package solver

/*
    Represents a preserved state of the sudoku solver. X and Y are the indices
    where the solver was at this state and Grid contains the values of the 
    grid at the time of this state. Value is the last tested number.
*/
type SolverState struct {
    X int8
    Y int8
    Value int8
    Grid [9][9]int8
}

func NewSolverState(x, y, value int8, grid *[9][9]int8) *SolverState {
    state := new(SolverState)
    state.X = x
    state.Y = y
    state.Value = value

    var gridCopy [9][9]int8
    for i, v := range grid {
        copy(gridCopy[i][:], v[:])
    }

    state.Grid = gridCopy

    return state
}