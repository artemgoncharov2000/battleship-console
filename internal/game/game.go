package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/artemgoncharov2000/battleship-console/internal/board"
	"github.com/artemgoncharov2000/battleship-console/internal/ships"
)

type Game struct {
	playerBoard board.Board
	enemyBoard  board.Board
	nextMove    string
}

func Create() Game {
	playerBoard := board.Create()

	enemyBoard := board.Create()
	enemyBoard.PlaceShipsRandomly()

	return Game{
		playerBoard: playerBoard,
		enemyBoard:  enemyBoard,
		nextMove:    "player",
	}
}

func (game *Game) Start() {
	game.placePlayerShips()
	clearConsole()
	fmt.Println("Game is started!!!")
	fmt.Println()
	var errReason string = ""
	for !game.enemyBoard.IsAllShipsSunk() && !game.playerBoard.IsAllShipsSunk() {

		if errReason != "" {
			fmt.Println(errReason)
			errReason = ""
		}

		game.enemyBoard.Print(true)
		fmt.Println()
		game.playerBoard.Print(false)
		fmt.Println("Enemy stats")
		fmt.Println("Ships afloat", game.enemyBoard.TotalShips)
		fmt.Println("Sunken ships", game.enemyBoard.SunkenShips)

		if game.nextMove == "player" {
			fmt.Println("Your move")
			fmt.Print("Enter row, column: ")

			input, err := getUserInput()

			if err != nil {
				fmt.Println(err)
				continue
			}

			splittedInput := strings.Split(input, ", ")
			fmt.Println(splittedInput)
			row, err := strconv.Atoi(splittedInput[0])

			if err != nil || row < 0 || row > 9 {
				errReason = "Incorrect row input, row should be integer value in range [0..9]"
				continue
			}

			column, err := strconv.Atoi(splittedInput[1])

			if err != nil {
				errReason = "Incorrect column input, column should be integer value in range [0..9]"
				continue
			}

			damagedOrSunk, isOcean := game.enemyBoard.ShootAt(row, column)

			if damagedOrSunk && !isOcean {
				clearConsole()
				fmt.Println("Enemy's ship was damaged or sunk")
				continue
			}

			if !damagedOrSunk {
				errReason = "You have already shot at this point"
				continue
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
	}

	fmt.Println("Game is over")

	if game.enemyBoard.IsAllShipsSunk() {
		fmt.Println("You won!!!")
	} else {
		fmt.Println("You lose(")
	}
}

func (game *Game) placePlayerShips() {
	totalShipsToPlace := 10
	shipsMapCount := map[string]int{
		ships.Carrier:    1,
		ships.Battleship: 2,
		ships.Submarine:  3,
		ships.Destroyer:  4,
	}

	var errReason string = ""

	for totalShipsToPlace > 0 {

		clearConsole()

		if errReason != "" {
			fmt.Println(errReason)
			errReason = ""
		}

		game.playerBoard.Print(false)
		fmt.Println("Choose option to place")
		fmt.Println("Enter 1 to place ship manually")
		fmt.Println("Enter 2 to place all ships randomly")
		fmt.Print("Input: ")
		input, err := getUserInput()

		if err != nil {
			errReason = "Incorrect input"
			continue
		}

		switch input {
		case "1":
			game.placePlayerShipManually(&shipsMapCount)
			totalShipsToPlace -= 1
			clearConsole()
		case "2":
			game.playerBoard.PlaceShipsRandomly()
			return
		default:
			errReason = "Wrong command, try again"
			continue
		}

	}

}

func (game *Game) placePlayerShipManually(shipsMapCount *map[string]int) {
	var errReason string = ""
	for {
		clearConsole()

		if errReason != "" {
			fmt.Println(errReason)
		}

		game.playerBoard.Print(false)
		fmt.Println("Enter \"type, row, column, horizontal\" (ex. \"Destroyer, 5, 5, true\") to place ship")
		printShipsToPlace(*shipsMapCount)
		fmt.Print("Input: ")
		input, err := getUserInput()

		if err != nil {
			errReason = "Incorrect input"
			continue
		}

		splittedInput := strings.Split(strings.Replace(input, "\n", "", -1), ", ")

		if len(splittedInput) < 4 {
			errReason = "Incorrect input, omitted arguments"
			continue
		}
		fmt.Println(splittedInput)
		shipType := splittedInput[0]
		rowAsString := splittedInput[1]
		columnAsString := splittedInput[2]
		horizontal := splittedInput[3]

		if count, ok := (*shipsMapCount)[shipType]; !ok || count == 0 {
			errReason = "Invalid ship type or all ships of this type have ended"
			continue
		}

		row, err := strconv.Atoi(rowAsString)

		if err != nil || row < 0 || row > 9 {
			errReason = "Incorrect row input, row should be integer value in range [0..9]"
		}

		column, err := strconv.Atoi(columnAsString)

		if err != nil || column < 0 || column > 9 {
			errReason = "Incorrect column input, column should be integer value in range [0..9]"
		}

		if horizontal != "true" && horizontal != "false" {
			errReason = "Incorrect horizontal input, horizonal should be true or false"
			continue
		}

		ship := createShip(shipType, row, column, horizontal == "true")

		if !game.playerBoard.CanPlaceShip(ship) {
			errReason = "Cannot place ship there, try another place"
			continue
		}

		game.playerBoard.PlaceShip(&ship)
		(*shipsMapCount)[shipType] -= 1
		return
	}
}

func createShip(shipType string, row, column int, horizontal bool) ships.Ship {
	var ship ships.Ship
	switch shipType {
	case ships.Carrier:
		ship = ships.CreateCarrier(row, column, horizontal)
	case ships.Battleship:
		ship = ships.CreateBattleship(row, column, horizontal)
	case ships.Submarine:
		ship = ships.CreateSubmarine(row, column, horizontal)
	case ships.Destroyer:
		ship = ships.CreateDestroyer(row, column, horizontal)
	}

	return ship
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

func getUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.Replace(input, "\n", "", -1), err
}
