package main

import "github.com/artemgoncharov2000/battleship-console/board"


func main() {
	playerBoard := board.Create()
	playerBoard.PlaceShipsRandomly()
	playerBoard.Print()
}