package ship

import "fmt"

type Ship struct {
	shipName string
	size     uint
}

func New() *Ship {
	var shipTest = &Ship{
		shipName: "     ",
		size:     0,
	}
	return shipTest //pointer to cell
}

func (s *Ship) Set(name string, sizeOfShip uint) {
	s.shipName = name
	s.size = sizeOfShip
}

func (s *Ship) Ship() (string, uint) {
	return s.shipName, s.size
}

func (s *Ship) DecrementSize() {
	s.size = s.size - 1
	if s.size == 0 {
		fmt.Println("You have attacked the entire ", s.shipName)
	}
}
