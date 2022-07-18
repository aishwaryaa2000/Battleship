package main

import (
	"battleShip/service"
)

func main() {

	board := service.MakeBoard()
	player := service.MakePlayer()
	service.GameStart(board, player)

}
