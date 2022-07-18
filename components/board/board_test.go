package board

import (
	"battleShip/components/cell"
	"testing"
)

func TestNew(t *testing.T) { //Test___()
	board := New(5, 6)
	actualRowSize, actualColSize := board.GetRowColSize()
	var (
		expectedRowSize, expectedColSize uint8 = 5, 6
	)

	// size is proper
	if expectedRowSize != actualRowSize {
		t.Error("Actual row size is ", actualRowSize, "but expected row size is ", expectedRowSize)
	}
	if expectedColSize != actualColSize {
		t.Error("Actual col size is ", actualRowSize, "but expected col size is ", expectedRowSize)
	}

	// all cells should have no mark
	var i, j uint8
	for i = 0; i < board.rowSize; i++ {
		for j = 0; j < board.colSize; j++ {
			icell := board.NCells[i][j]
			actualMark := icell.Cell()
			expectedMark := cell.NoMark
			if (actualMark) != expectedMark {
				t.Error("Actual is ", actualMark, "but expected is ", expectedMark)
			}
		}
	}

}
