package grid

/*
	Represents a row in the Sudoku grid. The X field is the identified for the line,
	its value is referenced from the top-left of the grid.
*/
type Row struct {
	X int8
}

func (r Row) Contains(val int8, grid *[9][9]int8) bool {
	var y int8
	for y = 0; y < 9; y++ {
		if grid[r.X][y] == val {
			return true
		}
	}

	return false
}
