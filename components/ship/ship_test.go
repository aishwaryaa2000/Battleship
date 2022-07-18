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


