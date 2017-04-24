package grid

import "testing"

func TestRowContainsWhenRowContainsValue(t *testing.T) {
	row := Row{0}

	var val int8

	var res bool
	var err error

	var grid [9][9]int8
	grid[0] = [9]int8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	val = 3

	res, err = row.Contains(val, grid)

	if !res {
		t.Errorf("row->Contains(%d) == %b, want %b", val, false, true)
	}

	if err != nil {
		t.Errorf("row->Contains(%d) throws error but was not supposed to", val)
	}
}

func TestRowContainsWhenRowDoesNotContainValue(t *testing.T) {
	row := Row{0}

	var val int8

	var res bool
	var err error

	var grid [9][9]int8
	grid[0] = [9]int8{1, 2, EMPTY_VALUE, 4, 5, 6, 7, 8, 9}

	val = 3

	res, err = row.Contains(val, grid)

	if res {
		t.Errorf("row->Contains(%d) == %b, want %b", val, true, false)
	}

	if err != nil {
		t.Errorf("row->Contains(%d) throws error but was not supposed to", val)
	}
}
