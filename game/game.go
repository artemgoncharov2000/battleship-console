package game

import (

	"fmt"
	"math/rand"


	"github.com/artemgoncharov2000/battleship-console/board"
)

type Game struct{
	playerBoard board.Board
	enemyBoard board.Board
	nextMove string
}

func Create() Game {
	playerBoard := board.Create()
	playerBoard.PlaceShipsRandomly()

	enemyBoard := board.Create()
	enemyBoard.PlaceShipsRandomly()


	return Game{
		playerBoard: playerBoard,
		enemyBoard: enemyBoard,
		nextMove: "player",
	}
}

func (game *Game) Start() {

	fmt.Println("Game is started!!!")
	fmt.Print()

	for !game.enemyBoard.IsAllShipsSunk() || !game.playerBoard.IsAllShipsSunk() {
		game.enemyBoard.Print(true)
		fmt.Println()
		game.playerBoard.Print(false)

		if game.nextMove == "player" {
			fmt.Println("Your move")
			fmt.Println("Enter row: ")

			var row int
			_, err := fmt.Scanf("%d", &row)

			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Enter column: ")
			
			var column int
			_, err = fmt.Scanf("%d", &column)

			if err != nil {
				fmt.Println("Wrong value of column")
				continue
			}

			if game.enemyBoard.ShootAt(row, column) {
				fmt.Println("Enemy's ship was damaged or sunk")
			}

			game.nextMove = "enemy"
		} else {
			fmt.Println("Enemy move")
			var row int
			var column int

			for game.playerBoard.IsDamaged(row, column) {
				row = rand.Intn(10)
				column = rand.Intn(10)
			}
			
			game.playerBoard.ShootAt(row, column)

			game.nextMove = "player"
		}

	}

	fmt.Println("Game is over")

	if game.enemyBoard.IsAllShipsSunk() {
		fmt.Println("You won!!!")
	} else {
		fmt.Println("You lose(")
	}
}