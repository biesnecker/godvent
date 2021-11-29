package utils

import "github.com/biesnecker/godvent/types"

func GenerateUlamSpiral(handler func(int, types.Coord) bool) {
	idx := 0
	lap := 0
	//dirs := [5]byte{'L', 'D', 'R', 'U', 'L'}
	currentDir := 4
	nextTurn := 0

	currentLocation := types.Coord{X: 0, Y: 0}

	for {
		switch idx {
		case 0:
		case 1:
			currentLocation.X++
		case 2:
			currentLocation.Y++
		case 3:
			currentLocation.X--
			currentDir = 4 // Set this so that the turn logic works correctly.
			nextTurn = idx + 1
		default:
			if idx == nextTurn {
				currentDir = (currentDir + 1) % 5
				if currentDir == 0 {
					lap++
				}

				// Set up the next turn.
				switch currentDir {
				case 0:
					nextTurn = idx + 1
				case 1:
					nextTurn = idx + (lap * 2)
				case 2:
					nextTurn = idx + (lap * 2) + 1
				case 3:
					nextTurn = idx + (lap * 2) + 1
				case 4:
					nextTurn = idx + (lap * 2) + 1
				}
			}

			// Move the location and call the handler.
			switch currentDir {
			case 0:
				currentLocation.X--
			case 1:
				currentLocation.Y--
			case 2:
				currentLocation.X++
			case 3:
				currentLocation.Y++
			case 4:
				currentLocation.X--

			}
		}
		if !handler(idx, currentLocation) {
			break
		}
		idx++
	}
}
