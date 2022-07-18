package service

import (
	"battleShip/components/board"
	"battleShip/components/player"
	"fmt"
)

func GameStart(gameBoard *board.Board, currentPlayer *player.Player) {
	var X, Y uint8
	var (
		shipCells, noOfTries uint = 15, 50
	)
	rowSize, colSize := gameBoard.GetRowColSize()
	fmt.Println("You only have ", noOfTries, " to attack all the ships")
	for shipCells > 0 && noOfTries > 0 {
		X, Y = inputFromPlayerToAttack(currentPlayer.GetName(), rowSize, colSize)
		icell := gameBoard.NCells[X][Y]
		shipInCell := icell.Ship()
		checkIfShipHitOrMiss(icell,gameBoard,currentPlayer,&noOfTries,&shipCells,shipInCell)
		fmt.Println("Number of hits : ", currentPlayer.GetHit(), " \nNumber of miss : ", currentPlayer.GetMiss(), "\nTrials left : ", noOfTries)
	}

	resultAnalysis(shipCells, noOfTries)

}

func resultAnalysis(shipCells, noOfTries uint) {
	fmt.Println("GAME OVER")
	if shipCells == 0 {
		fmt.Println("You won the game!\nYou needed ", 50-noOfTries, " of tries to win the game")
	} else {
		fmt.Println("Oh no!You have exhausted all your tries")
	}
}


