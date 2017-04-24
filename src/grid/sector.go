package grid

/*
	Represents a 3x3 sector in the Sudoku grid. The N and M indices are identifiers
	for the sector. N is the horizontal position of the sector and M is the vertical.
	Their values range from 0 to 2.
*/
type Sector struct {
	N int8
	M int8
}

func (s Sector) Contains(val int8, grid *[9][9]int8) bool {
	x_start := s.N * SECTOR_SIZE
	x_end := (s.N * SECTOR_SIZE) + SECTOR_SIZE
	y_start := s.M * SECTOR_SIZE
	y_end := (s.M * SECTOR_SIZE) + SECTOR_SIZE

	for x := x_start; x < x_end; x++ {
		for y := y_start; y < y_end; y++ {
			if grid[x][y] == val {
				return true
			}
		}
	}

	return false
}
