package battleShip

import (
	"fmt"
)

func PlayShipButtleground(matrixOrdex int) {

	ships := []int {1, 1, 1, 1, 2, 2, 2, 3, 3, 4}
	counter := 20

	battleField := generateBattleField(matrixOrdex)
	userField := generateBattleField(matrixOrdex)

	for i := 0; i < len(ships); i++ {
		placeShip(ships[i], &battleField)
	}

	fmt.Println("Just to see that ships are placed randomly every time.")
	displayField(battleField)

	var x, y int

	for counter != 0 {
		fmt.Print("Vertical coordinate: ")
		fmt.Scanf("%d", &x)
		fmt.Print("Horizontal coordinate: ")
		fmt.Scanf("%d", &y)

		if battleField[x][y] != "O" {
			userField[x][y] = battleField[x][y]
			displayField(userField)
			fmt.Println("GOT IT!")
			counter--
		} else {
			displayField(userField)
			fmt.Println("MISSED")
		}
	}

	fmt.Println("CONGRATULATIONS")
}

func displayField(field [][]string) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			fmt.Print(field[i][j], " ")
		}

		fmt.Println()
	}
}
