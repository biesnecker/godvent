package twentyseventeen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type nodestate int

const (
	CLEAN nodestate = iota
	INFECTED
	WEAKENED
	FLAGGED
)

func turnLeft(direction int) int {
	switch direction {
	case UP:
		return LEFT
	case DOWN:
		return RIGHT
	case LEFT:
		return DOWN
	case RIGHT:
		return UP
	default:
		panic("Unknown direction")
	}
}

func turnRight(direction int) int {
	switch direction {
	case UP:
		return RIGHT
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	case RIGHT:
		return DOWN
	default:
		panic("Unknown direction")
	}
}

func reverse(direction int) int {
	switch direction {
	case UP:
		return DOWN
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	default:
		panic("Unknown direction")
	}
}

func readInputDayTwentyTwo(fp *bufio.Reader) (map[types.Coord]nodestate, types.Coord) {
	res := make(map[types.Coord]nodestate)
	var width, height int
	utils.ReadStringsWithIndex(fp, func(x int, s string) {
		width = len(s)
		height = x
		for y := 0; y < len(s); y++ {
			if s[y] == '#' {
				res[types.Coord{X: x, Y: y}] = INFECTED
			}
		}
	})
	return res, types.Coord{X: height / 2, Y: width / 2}
}

func simulateD22(
	grid map[types.Coord]nodestate,
	location types.Coord,
	steps int,
	partA bool) int {
	infected := 0
	direction := UP
	for i := 0; i < steps; i++ {
		switch grid[location] {
		case CLEAN:
			direction = turnLeft(direction)
			if partA {
				grid[location] = INFECTED
				infected++
			} else {
				grid[location] = WEAKENED
			}
		case INFECTED:
			direction = turnRight(direction)
			if partA {
				delete(grid, location)
			} else {
				grid[location] = FLAGGED
			}
		case WEAKENED:
			if !partA {
				grid[location] = INFECTED
				infected++
			}
		case FLAGGED:
			if !partA {
				direction = reverse(direction)
				grid[location] = CLEAN
			}
		default:
			panic("unknown state")
		}

		newX := location.X
		newY := location.Y
		switch direction {
		case UP:
			newX--
		case DOWN:
			newX++
		case LEFT:
			newY--
		case RIGHT:
			newY++
		}
		location = types.Coord{X: newX, Y: newY}
	}
	return infected
}

func DayTwentyTwoA(fp *bufio.Reader) string {
	grid, center := readInputDayTwentyTwo(fp)
	solution := simulateD22(grid, center, 10000, true)
	return strconv.Itoa(solution)
}

func DayTwentyTwoB(fp *bufio.Reader) string {
	grid, center := readInputDayTwentyTwo(fp)
	solution := simulateD22(grid, center, 10000000, false)
	return strconv.Itoa(solution)
}
