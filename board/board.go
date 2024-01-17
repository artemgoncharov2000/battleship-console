package board

import (
	"fmt"
	"math/rand"

	"github.com/artemgoncharov2000/battleship-console/ships"
)

type Board struct{
	ocean [10][10]ships.Ship
}

func (board Board) Print() {
	for _, row := range board.ocean {
		str := ""
		for _, item := range row {
			str += item.GetString()
		}
		fmt.Println(str)
	}
}

func Create() Board {
	var ocean = [10][10]ships.Ship{}

	for i, row:= range ocean {
		for j := range row {
			ocean[i][j] = ships.CreateOcean(i, j)
		}
	}

	return Board{
		ocean: ocean,
	}
}

func (board *Board) PlaceShipsRandomly() {
	// Carrier
	board.createShips(ships.CreateCarrier, 1)

	// Battleships
	board.createShips(ships.CreateBattleship, 2)

	// Submarine
	board.createShips(ships.CreateSubmarine, 3)

	// Destroyer
	board.createShips(ships.CreateDestroyer, 4)
}

func (board *Board) createShips(createFn func(bowRow, bowColumn int, isHorizonal bool) ships.Ship, count int) {
	for i := 0; i < count; i++ {
		flag := false

		for !flag {
			row := rand.Intn(10)
			column := rand.Intn(10)
			horizontal := rand.Intn(2) == 1
	
			ship := createFn(row, column, horizontal)
			
			if board.canPlaceShip(&ship) {
				board.placeShip(&ship)
				flag = true
			}
		}
	}
}

func (board *Board) placeShip(ship *ships.Ship) {
	if ship.IsHorizontal {
		for i := ship.BowColumn; i < ship.BowColumn + ship.Size; i++ {
			board.ocean[ship.BowRow][i] = *ship
		}
	} else {
		for i := ship.BowRow; i < ship.BowRow + ship.Size; i++ {
			board.ocean[i][ship.BowColumn] = *ship
		}
	}
}

func (board Board) canPlaceShip(ship *ships.Ship) bool {
	if ship.IsHorizontal {
		if ship.BowColumn + ship.Size - 1 > 9 {
			return false
		}
		
		for i := max(ship.BowColumn - 1, 0); i < min(ship.BowColumn + ship.Size + 1, 10); i++ {

			if board.ocean[min(ship.BowRow + 1, 9)][i].IsOccupied() || board.ocean[max(ship.BowRow - 1, 0)][i].IsOccupied() || board.ocean[ship.BowRow][i].IsOccupied() {
				return false
			}
		}
	} else {
		if ship.BowRow + ship.Size - 1 > 9 {
			return false
		}

		for i := max(ship.BowRow - 1, 0); i < min(ship.BowRow + ship.Size + 1, 10); i++ {
			if board.ocean[i][min(ship.BowColumn + 1, 9)].IsOccupied() || board.ocean[i][max(ship.BowColumn - 1, 0)].IsOccupied() || board.ocean[i][ship.BowColumn].IsOccupied() {
				return false
			}
		}
	}

	return true
}