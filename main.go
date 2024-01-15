package main

import (


	"github.com/artemgoncharov2000/battleship-console/ship"
)

func main() {
	ship := ship.New("Cruiser", 3, false, false, 0, 0)
	ship.PrintName()
}