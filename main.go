package main

import "github.com/artemgoncharov2000/battleship-console/game"


func main() {
	game := game.Create()
	game.Start()
}