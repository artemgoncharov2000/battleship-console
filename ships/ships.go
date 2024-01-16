package ships

const (
	Carrier    = "Carrier"
	Battleship = "Battleship"
	Cruiser    = "Cruiser"
	Submarine  = "Submarine"
	Destroyer  = "Destroyer"
)

type Ship struct {
	bowRow       int
	bowColumn    int
	size         int
	name         string
	isHorizontal bool
	isOcean      bool
	hit          []bool
}

func (ship Ship) getString() string {
	if ship.isOcean {
		return "*"
	}

	return "x"
}

func CreateCarrier(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		bowRow:       bowRow,
		bowColumn:    bowColumn,
		size:         5,
		name:         Carrier,
		isHorizontal: isHorizonal,
		isOcean:      false,
		hit:          make([]bool, 5),
	}
}

func CreateBattleship(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		bowRow:       bowRow,
		bowColumn:    bowColumn,
		size:         4,
		name:         Battleship,
		isHorizontal: isHorizonal,
		isOcean:      false,
		hit:          make([]bool, 4),
	}
}

func CreateCruiser(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		bowRow:       bowRow,
		bowColumn:    bowColumn,
		size:         3,
		name:         Cruiser,
		isHorizontal: isHorizonal,
		isOcean:      false,
		hit:          make([]bool, 3),
	}
}

func CreateSubmarine(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		bowRow:       bowRow,
		bowColumn:    bowColumn,
		size:         3,
		name:         Submarine,
		isHorizontal: isHorizonal,
		isOcean:      false,
		hit:          make([]bool, 3),
	}
}

func CreateDestroyer(bowRow, bowColumn int, isHorizonal bool) Ship {
	return Ship{
		bowRow:       bowRow,
		bowColumn:    bowColumn,
		size:         2,
		name:         Destroyer,
		isHorizontal: isHorizonal,
		isOcean:      false,
		hit:          make([]bool, 2),
	}
}