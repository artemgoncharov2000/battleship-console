package main

import "github.com/artemgoncharov2000/battleship-console/internal/game"

func main() {
	game := game.Create()
	game.Start()
}
