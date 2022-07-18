package service

import (
	"battleShip/components/board"
	"battleShip/components/cell"
	"battleShip/components/ship"
	"math/rand"
	"strconv"
	"time"
)

func MakeBoard() *board.Board {

	rowSize, colSize := inputSizeFromPlayer()
	board := board.New(rowSize, colSize)
	BoardInit(board) //Initializing board with 5 random ships of size 5 4 3 2 1
	board.Display()
	return board
}

func randomNumberGenerator(rowSize, colSize uint8) (int, int, int) {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	XcordinateRandom := (random.Intn(int(rowSize)))
	YcordinateRandom := (random.Intn(int(colSize)))
	orientation := random.Intn(2)

	return XcordinateRandom, YcordinateRandom, orientation
}

func BoardInit(currentBoard *board.Board) {
	//Initializing board with 5 random ships of size 5 4 3 2 1 horizontally or vertically
	noOfShip := 4
	var gameShips []*ship.Ship
	for i := 0; i < 5; i++ {
		t := strconv.Itoa(i + 1)
		name := "Ship" + t
		newShip := ship.New()
		newShip.Set(name, uint(i+1))
		gameShips = append(gameShips, newShip)
	}

	rowSize, colSize := currentBoard.GetRowColSize()

	for noOfShip >= 0 {
		XcordinateRandom, YcordinateRandom, orientation := randomNumberGenerator(rowSize, colSize)
		/*Orientation is to check if vertical placement should be done first then horizontal or
		vice versa thereby increasing randomness*/
		if orientation == 1 {
			rowStart, rowEnd, okBool := checkIfHorizontalOrVertical(currentBoard.NCells, XcordinateRandom, YcordinateRandom, int(rowSize), gameShips[noOfShip], 'V')
			if okBool {
				placeShip(currentBoard.NCells, rowStart, rowEnd, YcordinateRandom, 'V', gameShips[noOfShip])
				noOfShip--
				continue
			}
			colStart, colEnd, okBool := checkIfHorizontalOrVertical(currentBoard.NCells, XcordinateRandom, YcordinateRandom, int(colSize), gameShips[noOfShip], 'H')
			if okBool {
				placeShip(currentBoard.NCells, colStart, colEnd, XcordinateRandom, 'H', gameShips[noOfShip])
				noOfShip--
				continue
			}
		} else {
			colStart, colEnd, okBool := checkIfHorizontalOrVertical(currentBoard.NCells, XcordinateRandom, YcordinateRandom, int(colSize), gameShips[noOfShip], 'H')
			if okBool {
				placeShip(currentBoard.NCells, colStart, colEnd, XcordinateRandom, 'H', gameShips[noOfShip])
				noOfShip--
				continue
			}
			rowStart, rowEnd, okBool := checkIfHorizontalOrVertical(currentBoard.NCells, XcordinateRandom, YcordinateRandom, int(rowSize), gameShips[noOfShip], 'V')
			if okBool {
				placeShip(currentBoard.NCells, rowStart, rowEnd, YcordinateRandom, 'V', gameShips[noOfShip])
				noOfShip--
				continue
			}

		}
	}

}

func checkIfHorizontalOrVertical(b [][]*cell.Cell, XcordinateRandom, YcordinateRandom, rowOrColSize int, s *ship.Ship, orientation rune) (int, int, bool) {
	var icell *cell.Cell
	_, sizeOfShip := s.Ship()
	icell = b[XcordinateRandom][YcordinateRandom]
	if icell.IfShipInCell() {
		return -1, -1, false
	}

	upOrLeftMin, downOrRightMax := YcordinateRandom, YcordinateRandom

	if orientation == 'V' {
		upOrLeftMin, downOrRightMax = XcordinateRandom, XcordinateRandom
	}

	for upOrLeftMin >= 0 {
		if orientation == 'V' {
			icell = b[upOrLeftMin][YcordinateRandom]
		} else {
			icell = b[XcordinateRandom][upOrLeftMin]
		}
		if icell.IfShipInCell() {
			break
		}
		upOrLeftMin--
	}
	upOrLeftMin++

	for downOrRightMax < rowOrColSize {
		if orientation == 'V' {
			icell = b[downOrRightMax][YcordinateRandom]
		} else {
			icell = b[XcordinateRandom][downOrRightMax]
		}
		if icell.IfShipInCell() {
			break
		}
		downOrRightMax++
	}
	downOrRightMax--

	if downOrRightMax-upOrLeftMin+1 >= int(sizeOfShip) {
		return upOrLeftMin, upOrLeftMin + int(sizeOfShip), true
	}

	return -1, -1, false

}
