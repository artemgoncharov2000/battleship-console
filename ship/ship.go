package ship

import "fmt"

type Ship struct{
	bowRow int
	bowColumn int
	size int
	name string
	isHorizontal bool
	isOcean bool
	hit []int
}

func New(name string, size int, isHorizontal bool, isOcean bool, bowRow, bowColumn int) Ship {
	return Ship{
		name: name,
		size: size,
		isHorizontal: isHorizontal,
		isOcean: isOcean,
		hit: make([]int, size),
		bowRow: bowRow,
		bowColumn: bowColumn,
	}
}

func (ship Ship) PrintName() {
	fmt.Printf("Ship name is %v", ship.name)
}

