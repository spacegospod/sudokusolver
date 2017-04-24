package solver

import (
    "testing"
)

func TestgetNextCellWhenCurrentCellIsLastCell(t *testing.T) {
    _, _, err := getNextCell(8, 8)

    if err == nil {
        t.Errorf("Expected getNextCell to return error")
    }
}

func TestgetNextCellWhenCurrentRowIsLastRow(t *testing.T) {
    x, y, err := getNextCell(8, 7)

    if err != nil {
        t.Errorf("Expected getNextCell not to return error")
    }

    if x != 0 || y != 8 {
        t.Errorf("Expected getNextCell to return (%d, %d, nil), but got (%d, %d, %s",
            0, 8, x, y, err)
    }
}

func TestgetNextCellWhenCurrentRowIsNotLastRow(t *testing.T) {
    x, y, err := getNextCell(3, 4)

    if err != nil {
        t.Errorf("Expected getNextCell not to return error")
    }

    if x != 4 || y != 4 {
        t.Errorf("Expected getNextCell to return (%d, %d, nil), but got (%d, %d, %s",
            4, 4, x, y, err)
    }
}