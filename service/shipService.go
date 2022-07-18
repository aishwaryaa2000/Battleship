package service

import (
	"battleShip/components/board"
	"battleShip/components/cell"
	"battleShip/components/player"
	"battleShip/components/ship"
	"fmt"

)

func checkIfShipHitOrMiss(icell *cell.Cell,b *board.Board, currentPlayer *player.Player, noOfTries ,shipCells *uint,s* ship.Ship){

	mark := icell.Cell()
	switch mark{
	case cell.Hit,cell.Miss:
		fmt.Println("Already attacked this cell")
		return
	}
	Okbool := icell.IfShipInCell()
	switch Okbool{
	case true:
		hitTheShip(icell, b, currentPlayer, shipCells,noOfTries,s)
	case false:
		missedTheShip(icell, b, currentPlayer, noOfTries)
	}

}

func hitTheShip(icell *cell.Cell, b *board.Board, currentPlayer *player.Player, shipCells *uint, noOfTries *uint,s *ship.Ship) {
	fmt.Println("Hurray! You hit the ship")
	icell.SetMark(cell.Hit)
	b.DisplayHitMiss()
	currentPlayer.IncrementHit()
	s.DecrementSize()
	*shipCells--
	*noOfTries--

}

func missedTheShip(icell *cell.Cell, b *board.Board, currentPlayer *player.Player, noOfTries *uint) {
	fmt.Println("Oh no,you missed the ship")
	icell.SetMark(cell.Miss)
	b.DisplayHitMiss()
	currentPlayer.IncrementMiss()
	*noOfTries--
}

func placeShip(b [][]*cell.Cell, min, max, cordinate int, direction rune,s *ship.Ship) {
	//to place ships on a board
	var icell *cell.Cell
	for i := min; i < max; i++ {
		if direction == 'H' {
			icell = b[cordinate][i]
		} else {
			icell = b[i][cordinate]
		}
		icell.SetShip(s)
	}

}