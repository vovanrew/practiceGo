package battleShip

/*
These types are used to match direction to place ship.
If field is empty and matches ship size it can be filled by ship symbols

      ^
      |
<---point--->
	  |
	  V

X - vertical coordinate
Y - horizontal coordinate
*/

type PlaceSideResolver interface {
	checkSpaceAndFill(x, y, shipSize int, field [][]string) bool
}

type LeftSideResolver struct {}

func (l LeftSideResolver) checkSpaceAndFill(x, y, shipSize int, field [][]string) bool {
	if y - shipSize >= 0 {
		for i := y; i >= y - shipSize; i-- {
			if field[x][i] != "O" {
				return false
			}
		}

		for i := y; i > y - shipSize; i-- {
			field[x][i] = "<"
		}

		return true
	}

	return false
}

type RightSideResolver struct {}

func (r RightSideResolver) checkSpaceAndFill(x, y, shipSize int, field [][]string) bool {
	if y + shipSize < len(field) {
		for i := y; i <= y + shipSize; i++ {
			if field[x][i] != "O" {
				return false
			}
		}

		for i := y; i < y + shipSize; i++ {
			field[x][i] = ">"
		}

		return true
	}

	return false
}

type DownSideResolver struct {}

func (b DownSideResolver) checkSpaceAndFill(x, y, shipSize int, field [][]string) bool {
	if x + shipSize < len(field) {
		for i := x; i <= x + shipSize; i++ {
			if field[i][y] != "O" {
				return false
			}
		}

		for i := x; i < x + shipSize; i++ {
			field[i][y] = "V"
		}

		return true
	}

	return false
}

type UpSideResolver struct {}

func (u UpSideResolver) checkSpaceAndFill(x, y, shipSize int, field [][]string) bool {
	if x - shipSize >= 0 {
		for i := x; i >= x - shipSize; i-- {
			if field[i][y] != "O" {
				return false
			}
		}

		for i := x; i > x - shipSize; i-- {
			field[i][y] = "^"
		}

		return true
	}

	return false
}
