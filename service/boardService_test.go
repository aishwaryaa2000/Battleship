package service

import (
	"testing"
	"battleShip/components/board"
	"battleShip/components/ship"
	"strconv"

)

func TestCheckIfHorizontalOrVertical(t *testing.T) {
	currentBoard := board.New(5, 5)
	// noOfShip := 4
	var gameShips []*ship.Ship
	for i := 0; i < 5; i++ {
		t := strconv.Itoa(i + 1)
		name := "Ship" + t
		newShip := ship.New()
		newShip.Set(name, uint(i+1))
		gameShips = append(gameShips, newShip)
	}


	type input struct{
		XcordinateRandom,YcordinateRandom,rowOrColSize int
		currentShip *ship.Ship
		orientation rune
	}
	type output struct{
		start,end int
		okBool bool
	}
	var listTest = []struct { //slice of struct
		inputParameters  input
		expected output
	}{
		{
			input{
				2,
				3,
				5,
				gameShips[4], //ship5
				'H',
			},
			output{
				0,
				5,
				true,
			},
		},
		{
			input{
				0,
				2,
				5,
				gameShips[3], //ship4
				'V',
			},
			output{
				-1,
				-1,
				false,
			},
		},
		{
			input{
				0,
				2,
				5,
				gameShips[3], //ship4
				'H',
			},
			output{
				0,
				4,
				true,
			},
		},
		{
			input{
				3,
				0,
				5,
				gameShips[1], //ship2
				'V',
			},
			output{
				3,
				5,
				true,
			},
		},
		{
			input{
				2,
				2,
				5,
				gameShips[0], //ship1
				'V',
			},
			output{
				-1,
				-1,
				false,
			},
		},
	}

	for _, structAtiIndex := range listTest {
		startActual, endActual, okBoolActual := checkIfHorizontalOrVertical(currentBoard.NCells, structAtiIndex.inputParameters.XcordinateRandom, structAtiIndex.inputParameters.YcordinateRandom,structAtiIndex.inputParameters.rowOrColSize,structAtiIndex.inputParameters.currentShip,structAtiIndex.inputParameters.orientation)
		if okBoolActual {
			placeShip(currentBoard.NCells, startActual, endActual, structAtiIndex.inputParameters.XcordinateRandom, structAtiIndex.inputParameters.orientation, structAtiIndex.inputParameters.currentShip)
		}
		if startActual != structAtiIndex.expected.start{
			t.Error("Actual status is ", startActual, "but expected is ", structAtiIndex.expected.start)
		}
		if endActual != structAtiIndex.expected.end{
			t.Error("Actual status is ", endActual, "but expected is ", structAtiIndex.expected.end)
		}
		if okBoolActual != structAtiIndex.expected.okBool{
			t.Error("Actual status is ", okBoolActual, "but expected is ", structAtiIndex.expected.okBool)
		}
	}

}