package grid

import "testing"

func TestLineContainsWhenLineContainsValue(t *testing.T) {
	line := Line{0}

	var val int8

	var res bool
	var err error

	var grid [9][9]int8

	var sample int8 = 1
	var x int8
	for x = 0; x < GRID_SIZE; x++ {
		grid[x][line.Y] = sample
		sample++
	}

	val = 3

	res, err = line.Contains(val, grid)

	if !res {
		t.Errorf("line->Contains(%d) == %b, want %b", val, false, true)
	}

	if err != nil {
		t.Errorf("line->Contains(%d) throws error but was not supposed to", val)
	}
}

func TestLineContainsWhenLineDoesNotContainValue(t *testing.T) {
	line := Line{0}

	var val int8

	var res bool
	var err error

	var grid [9][9]int8

	var sample int8 = 1
	var x int8
	for x = 0; x < GRID_SIZE-1; x++ {
		grid[x][line.Y] = sample
		sample++
	}

	val = 9

	res, err = line.Contains(val, grid)

	if res {
		t.Errorf("line->Contains(%d) == %b, want %b", val, true, false)
	}

	if err != nil {
		t.Errorf("line->Contains(%d) throws error but was not supposed to", val)
	}
}
