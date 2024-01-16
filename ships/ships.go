package ships

const (
	Carrier    = "Carrier"
	Battleship = "Battleship"
	Cruiser    = "Cruiser"
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

func (ship Ship) GetString() string {
	if ship.IsOcean {
		return "~"
	}

	return "*"
}

func (ship Ship) IsOccupied() bool {
	return !ship.IsOcean
}

func CreateCarrier(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         5,
		Name:         Carrier,
		IsHorizontal: isHorizonal,
		IsOcean:      false,
		Hit:          make([]bool, 5),
	}
}

func CreateBattleship(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         4,
		Name:         Battleship,
		IsHorizontal: isHorizonal,
		IsOcean:      false,
		Hit:          make([]bool, 4),
	}
}

func CreateCruiser(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         3,
		Name:         Cruiser,
		IsHorizontal: isHorizonal,
		IsOcean:      false,
		Hit:          make([]bool, 3),
	}
}

func CreateSubmarine(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         3,
		Name:         Submarine,
		IsHorizontal: isHorizonal,
		IsOcean:      false,
		Hit:          make([]bool, 3),
	}
}

func CreateDestroyer(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		BowRow:       bowRow,
		BowColumn:    bowColumn,
		Size:         2,
		Name:         Destroyer,
		IsHorizontal: isHorizonal,
		IsOcean:      false,
		Hit:          make([]bool, 2),
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
