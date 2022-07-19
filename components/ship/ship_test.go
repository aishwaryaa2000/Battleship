package ship

import (
	"testing"
)
func TestNew(t *testing.T) { //Test___()
	newShip := New()
	var actualName = newShip.shipName
	var expectedName = "     "
	var actualSize = newShip.size
	var expectedSize uint= 0
	if actualName != expectedName {
		t.Error("Actual is ", actualName, "but expected is ", expectedName)
	}
	if actualSize != expectedSize {
		t.Error("Actual is ", actualSize, "but expected is ", expectedSize)
	}

}

func TestSetShip(t *testing.T){
	var listTest = []struct { //slice of struct
		input    Ship
		expected Ship
	}{
		{
			Ship{
				"Ship1",1,
			},
			Ship{
				"Ship1",1,
			},
		},
		{
			Ship{
				"Ship2",2,
			},
			Ship{
				"Ship2",2,
			},
		},
		{
			Ship{
				"Ship4",4,
			},
			Ship{
				"Ship4",4,
			},
		},

	}

	for _, structAtiIndex := range listTest {
		newShip := New()
		newShip.Set(structAtiIndex.input.shipName,structAtiIndex.input.size)
		actualName,actualSize := newShip.Ship()
		if actualName != structAtiIndex.expected.shipName{
			t.Error("Actual status is ", actualName, "but expected is ", structAtiIndex.expected.shipName)
		}
		if actualSize != structAtiIndex.expected.size{
			t.Error("Actual status is ", actualSize, "but expected is ", structAtiIndex.expected.size)
		}
	}

}


