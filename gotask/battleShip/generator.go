package battleShip

import (
	"softgo/gotask"
)

func generateBattleField(matrixOrder int) [][]string {

	//8 - minimum matrix size allowed

	if matrixOrder < 8 {
		matrixOrder = 8
	}

	battleField := make([][]string, matrixOrder)

	row := make([]string, matrixOrder)

	for i := 0; i < len(battleField); i++ {
		for j := 0; j < len(row); j++ {
			row[j] = "O"		// '~' means water/waves
		}
		battleField[i] = row
		row = make([]string, matrixOrder)
	}

	return battleField
}

func placeShip(shipSize int, field *[][]string) {
	randX := gotask.GenerateRandWithRange(0, len(*field))
	randY := gotask.GenerateRandWithRange(0, len(*field))

	sideResolvers := [4]PlaceSideResolver{LeftSideResolver{}, RightSideResolver{}, UpSideResolver{}, DownSideResolver{}}
	randSide := gotask.GenerateRandWithRange(0, 4)

	isPlased := sideResolvers[randSide].checkSpaceAndFill(randX, randY, shipSize, *field)

	for isPlased == false {
		for i := 0; i < len(sideResolvers) && isPlased == false; i++ {
			isPlased = sideResolvers[i].checkSpaceAndFill(randX, randY, shipSize, *field)
		}

		if isPlased == false {
			randX, randY = nextCell(randX, randY, field)
		}
	}
}

//this function is used when ship can't be placed in current cell
func nextCell(x, y int, field *[][]string) (int, int) {
	if y + 1 < len(*field) {
		return x, y + 1
	} else if x + 1 < len(*field) {
		return x + 1, 0
	} else {
		return 0, 0
	}
}