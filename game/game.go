package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/artemgoncharov2000/battleship-console/board"
	"github.com/artemgoncharov2000/battleship-console/ships"
)

type Game struct {
	playerBoard board.Board
	enemyBoard  board.Board
	nextMove    string
}

func Create() Game {
	playerBoard := board.Create()
	playerBoard.PlaceShipsRandomly()

	enemyBoard := board.Create()
	enemyBoard.PlaceShipsRandomly()

	return Game{
		playerBoard: playerBoard,
		enemyBoard:  enemyBoard,
		nextMove:    "player",
	}
}

func (game *Game) Start() {

	fmt.Println("Game is started!!!")
	game.placePlayerShips()

	for !game.enemyBoard.IsAllShipsSunk() && !game.playerBoard.IsAllShipsSunk() {
		game.enemyBoard.Print(false)
		fmt.Println()
		game.playerBoard.Print(false)

		if game.nextMove == "player" {
			fmt.Println("Your move")
			fmt.Print("Enter row, column: ")

			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println(err)
				continue
			}

			splittedInput := strings.Split(strings.Replace(input, "\n", "", -1), ", ")
			fmt.Println(splittedInput)
			row, err := strconv.Atoi(splittedInput[0])

			if err != nil {
				fmt.Println("Wrong value of row", err)
				continue
			}

			column, err := strconv.Atoi(splittedInput[1])

			if err != nil {
				fmt.Println("Wrong value of column", err)
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

		clearConsole()

		fmt.Println("Enemy stats")
		fmt.Println("Total ships", game.enemyBoard.TotalShips)
		fmt.Println("Sunken ships", game.enemyBoard.SunkenShips)
	}

	fmt.Println("Game is over")

	if game.enemyBoard.IsAllShipsSunk() {
		fmt.Println("You won!!!")
	} else {
		fmt.Println("You lose(")
	}
}

func (game *Game) placePlayerShips() {

	shipsMapCount := map[string]int{
		ships.Carrier:    1,
		ships.Battleship: 2,
		ships.Submarine:  3,
		ships.Destroyer:  4,
	}

	// for totalShipToPlace > 0 {

	// }
	printShipsToPlace(shipsMapCount)

	fmt.Println("Choose option to place")
}

func printShipsToPlace(shipsMapCount map[string]int) {
	fmt.Printf("------------------------------\n")
	fmt.Printf("Piece          | Size | Amount\n")
	fmt.Printf("------------------------------\n")
	fmt.Printf("Destroyer      | 1    | %v\n", shipsMapCount[ships.Destroyer])
	fmt.Printf("Submarine      | 2    | %v\n", shipsMapCount[ships.Submarine])
	fmt.Printf("Battleship     | 3    | %v\n", shipsMapCount[ships.Battleship])
	fmt.Printf("Carrier        | 4    | %v\n", shipsMapCount[ships.Carrier])

}

func clearConsole() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
