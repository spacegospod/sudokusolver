package grid

/*
	Represents a line in the Sudoku grid. The Y field is the identified for the line,
	its value is referenced from the top-left of the grid.
*/
type Line struct {
	Y int8
}

func (l Line) Contains(val int8, grid *[9][9]int8) bool {
	var x int8
	for x = 0; x < GRID_SIZE; x++ {
		if grid[x][l.Y] == val {
			return true
		}
	}

	return false
}