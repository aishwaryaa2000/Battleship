package cell

import (
	"battleShip/components/ship"
	"testing"
	"strconv"
)

func TestNew(t *testing.T) { //Test___()
	newCell := New()
	var actual Mark = newCell.cellMark
	var expected Mark = NoMark

	if actual != expected {
		t.Error("Actual is ", actual, "but expected is ", expected)
	}
}

func TestCell(t *testing.T) {
	newCell := New()
	newCell.SetMark(Hit)
	var actual Mark = newCell.Cell()
	var expected Mark = Hit

	if actual != expected {
		t.Error("Actual is ", actual, "but expected is ", expected)
	}

}

func TestCellBulk(t *testing.T) {

	var listTest = []struct { //slice of struct
		input    Mark
		expected Mark
	}{
		{
			Hit, Hit,
		},
		{
			Miss, Miss,
		},
		{
			NoMark, NoMark,
		},
	}

	for _, structAtiIndex := range listTest {
		newCell := New()

		newCell.SetMark(structAtiIndex.input)
		actual := newCell.Cell()
		if actual != structAtiIndex.expected {
			t.Error("Actual status is ", actual, "but expected is ", structAtiIndex.expected)
		}
	}
}

func TestShipSetShipBulk(t *testing.T) {
	var gameShips []*ship.Ship
	for i := 0; i < 2; i++ {
		t := strconv.Itoa(i + 1)
		name := "Ship" + t
		newShip := ship.New()
		newShip.Set(name, uint(i+1))
		gameShips = append(gameShips, newShip)
	}

	ship1 := ship.New()
	ship1.Set("Ship1",1)
	ship2 := ship.New()
	ship2.Set("Ship2",2)

	var listTest = []struct { //slice of struct
		input    ship.Ship
		expected ship.Ship
	}{
		{
			*gameShips[0], 
			*ship1,
		},
		{
			*gameShips[1], 
			*ship2,
		},

	}

	for _, structAtiIndex := range listTest {
		newCell := New()

		newCell.SetShip(&structAtiIndex.input)
		actual := newCell.Ship()
		if *actual != structAtiIndex.expected {
			t.Error("Actual status is ", actual, "but expected is ", structAtiIndex.expected)
		}
	}
}

func TestIfShipInCellBulk(t *testing.T) {
	var gameShips []*ship.Ship
	for i := 0; i < 4; i++ {
		t := strconv.Itoa(i + 1)
		name := "Ship" + t
		newShip := ship.New()
		newShip.Set(name, uint(i))
		gameShips = append(gameShips, newShip)
		//Ship1 has size 0,ship2 has size 1 and so on for testing purposes
	}


	var listTest = []struct { //slice of struct
		input    ship.Ship
		expected bool
	}{
		{
			*gameShips[0], 
			false,
		},
		{
			*gameShips[1], 
			true,
		},
		{
			*gameShips[2], 
			true,
		},
		{
			*gameShips[3], 
			true,
		},

	}

	for _, structAtiIndex := range listTest {
		newCell := New()

		newCell.SetShip(&structAtiIndex.input)
		actual := newCell.IfShipInCell()
		if actual != structAtiIndex.expected {
			t.Error("Actual status is ", actual, "but expected is ", structAtiIndex.expected)
		}
	}
}