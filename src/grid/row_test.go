package grid

import "testing"

func TestRowContainsWhenRowContainsValue(t *testing.T) {
	row := Row{0}

	var val int8

	var res bool

	var grid [9][9]int8
	grid[0] = [9]int8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	val = 3

	res = row.Contains(val, &grid)

	if !res {
		t.Errorf("row->Contains(%d) == %b, want %b", val, false, true)
	}
}

func TestRowContainsWhenRowDoesNotContainValue(t *testing.T) {
	row := Row{0}

	var val int8

	var res bool

	var grid [9][9]int8
	grid[0] = [9]int8{1, 2, EMPTY_VALUE, 4, 5, 6, 7, 8, 9}

	val = 3

	res = row.Contains(val, &grid)

	if res {
		t.Errorf("row->Contains(%d) == %b, want %b", val, true, false)
	}
}
