package board

import (
	"battleShip/components/cell"
	"fmt"
)

type Board struct {
	NCells  [][]*cell.Cell //a slice of cell structure pointers
	rowSize uint8
	colSize uint8
}

func New(rowSizeByUser, colSizeByUser uint8) *Board {

	var testCells = [][]*cell.Cell{} //a slice of cell structure pointers

	for i := 0; i < int(rowSizeByUser); i++ {
		ithRow := []*cell.Cell{}
		for j := 0; j < int(colSizeByUser); j++ {
			newCell := cell.New() //new cell is created with noMark
			ithRow = append(ithRow, newCell)
		}
		testCells = append(testCells, ithRow)

	}

	var boardTest = &Board{
		NCells:  testCells,
		rowSize: rowSizeByUser,
		colSize: colSizeByUser,
	}
	return boardTest //pointer to board
}

func (currentBoard *Board) GetRowColSize() (uint8, uint8) {
	return currentBoard.rowSize, currentBoard.colSize
}

func (currentBoard *Board) DisplayHitMiss() {
	var i, j uint8
	for i = 0; i < currentBoard.rowSize; i++ {
		for j = 0; j < currentBoard.colSize; j++ {
			icell := currentBoard.NCells[i][j] //get a structure pointer of the cell at a particular index
			if icell.Cell() == cell.Hit || icell.Cell() == cell.Miss {
				fmt.Print(" | ", icell.Cell()) //get the mark of the cell at that index
			} else {
				fmt.Print(" |    ")

			}
		}
		fmt.Printf(" | Row %d \n", i)
	}
}

func (currentBoard *Board) Display() {
	var i, j uint8
	fmt.Print(" ")
	for i = 0; i < currentBoard.rowSize; i++ {
		for j = 0; j < currentBoard.colSize; j++ {
			icell := currentBoard.NCells[i][j]
			shipInCell := icell.Ship()
			name, _ := shipInCell.Ship() //get a structure pointer of the cell at a particular index
			fmt.Print(" | ", name)       //get the mark of the cell at that index
		}
		fmt.Printf(" |  Row %d \n ", i)
	}
}
