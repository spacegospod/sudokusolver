package grid

import "testing"

func TestSectorContainsWhenSectorContainsValue(t *testing.T) {
	sector := Sector{0, 0}

	var val int8

	var res bool

	var grid [9][9]int8

	var sample int8 = 1
	for x := 0; int8(x) < SECTOR_SIZE; x++ {
		for y := 0; int8(y) < SECTOR_SIZE; y++ {
			grid[x][y] = sample
			sample++
		}
	}

	val = 1

	res = sector.Contains(val, &grid)

	if !res {
		t.Errorf("sector->Contains(%d) == %b, want %b", val, false, true)
	}
}

func TestSectorContainsWhenSectorDoesNotContainValue(t *testing.T) {
	sector := Sector{0, 0}

	var val int8

	var res bool

	var grid [9][9]int8

	var sample int8 = 1
	for x := 0; int8(x) < SECTOR_SIZE; x++ {
		for y := 0; int8(y) < SECTOR_SIZE; y++ {
			grid[x][y] = sample
			sample++
		}
	}

	grid[0][0] = EMPTY_VALUE

	val = 1

	res = sector.Contains(val, &grid)

	if res {
		t.Errorf("sector->Contains(%d) == %b, want %b", val, true, false)
	}
}
