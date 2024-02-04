package ships

const (
	Carrier    = "Carrier"
	Battleship = "Battleship"
	Submarine  = "Submarine"
	Destroyer  = "Destroyer"
	Ocean      = "Ocean"
)

type Ship struct {
	BowRow       int
	BowColumn    int
	Size         int
	Name         string
	IsHorizontal bool
	IsOcean      bool
	Hit          []bool
}

func (ship Ship) GetString(hideShips bool) string {
	if ship.IsOcean || hideShips {
		return "  "
	}

	return " ■"
}

func (ship Ship) IsOccupied() bool {
	return !ship.IsOcean
}

func CreateCarrier(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         4,
		Name:         Carrier,
		IsHorizontal: isHorizonal,
		IsOcean:      false,
		Hit:          make([]bool, 4),
	}
}

func CreateBattleship(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         3,
		Name:         Battleship,
		IsHorizontal: isHorizonal,
		IsOcean:      false,
		Hit:          make([]bool, 3),
	}
}

func CreateSubmarine(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         2,
		Name:         Submarine,
		IsHorizontal: isHorizonal,
		IsOcean:      false,
		Hit:          make([]bool, 2),
	}
}

func CreateDestroyer(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         1,
		Name:         Destroyer,
		IsHorizontal: isHorizonal,
		IsOcean:      false,
		Hit:          make([]bool, 1),
	}
}

func CreateOcean(bowRow, bowColumn int) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         1,
		Name:         Ocean,
		IsHorizontal: false,
		IsOcean:      true,
		Hit:          make([]bool, 1),
	}
}

func (ship *Ship) ShootAt(row, column int) bool {
	if ship.IsHorizontal {
		if !ship.Hit[column-ship.BowColumn] {
			ship.Hit[column-ship.BowColumn] = true
			return true
		} else {
			return false
		}
	} else {
		if !ship.Hit[row-ship.BowRow] {
			ship.Hit[row-ship.BowRow] = true
			return true
		} else {
			return false
		}
	}
}

func (ship Ship) IsHit(row, column int) bool {
	if ship.IsHorizontal {
		return ship.Hit[column-ship.BowColumn]
	}

	return ship.Hit[row-ship.BowRow]
}

func (ship Ship) IsSunk() bool {
	for _, v := range ship.Hit {
		if !v {
			return false
		}
	}

	return true
}
