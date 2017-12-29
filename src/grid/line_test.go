package grid

import "testing"

func TestLineContainsWhenLineContainsValue(t *testing.T) {
	line := Line{0}

	var val int8

	var res bool

	var grid [9][9]int8

	var sample int8 = 1
	var x int8
	for x = 0; x < GRID_SIZE; x++ {
		grid[x][line.Y] = sample
		sample++
	}

	val = 3

	res = line.Contains(val, &grid)

	if !res {
		t.Errorf("line->Contains(%d) == %b, want %b", val, false, true)
	}
}

func TestLineContainsWhenLineDoesNotContainValue(t *testing.T) {
	line := Line{0}

	var val int8

	var res bool

	var grid [9][9]int8

	var sample int8 = 1
	var x int8
	for x = 0; x < GRID_SIZE-1; x++ {
		grid[x][line.Y] = sample
		sample++
	}

	val = 9

	res = line.Contains(val, &grid)

	if res {
		t.Errorf("line->Contains(%d) == %b, want %b", val, true, false)
	}
}
