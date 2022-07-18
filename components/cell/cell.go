package cell

import "battleShip/components/ship"

type Mark string

const (
	NoMark Mark = "   "
	Hit    Mark = " X "
	Miss   Mark = " M "
)

type Cell struct {
	cellMark Mark
	shipInCell     *ship.Ship
}

func New() *Cell {
	var cellTest = &Cell{
		cellMark: NoMark,
		shipInCell: ship.New(),
	}
	return cellTest //pointer to cell
}

func (c *Cell) Cell() Mark {
	return c.cellMark
}


func (c *Cell) Ship() *ship.Ship {
	return c.shipInCell
}

func (c *Cell) IfShipInCell()  bool{
	shipCurrent := c.shipInCell
	_,ssize := shipCurrent.Ship()
	return ssize!=0
}

func (c *Cell) SetShip(shipCurrent *ship.Ship) {

	c.shipInCell = shipCurrent

}

func (c *Cell) SetMark(markByUser Mark) {
	
	c.cellMark = markByUser

}