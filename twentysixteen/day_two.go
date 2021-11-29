package twentysixteen

import (
	"bufio"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

func findSolutionDayTwo(
	fp *bufio.Reader,
	startingLocation types.Coord,
	moveHandler func(loc types.Coord, move rune) types.Coord,
	outputHandler func(loc types.Coord) byte,
) []byte {
	var res []byte
	sloc := startingLocation
	utils.ReadStrings(fp, func(s string) {
		loc := sloc
		for _, c := range s {
			loc = moveHandler(loc, c)
		}
		res = append(res, outputHandler(loc))
		sloc = loc
	})
	return res
}

func moveHandlerDayTwoA(loc types.Coord, move rune) types.Coord {
	switch move {
	case 'L':
		loc.X--
	case 'R':
		loc.X++
	case 'U':
		loc.Y--
	case 'D':
		loc.Y++
	}
	loc.X = utils.BoundInt(loc.X, 0, 2)
	loc.Y = utils.BoundInt(loc.Y, 0, 2)
	return loc
}

func outputHandlerDayTwoA(loc types.Coord) byte {
	return byte((loc.Y*3)+loc.X) + 1 + '0'
}

var fancyKeypad = [5][5]byte{
	{0, 0, '1', 0, 0},
	{0, '2', '3', '4', 0},
	{'5', '6', '7', '8', '9'},
	{0, 'A', 'B', 'C', 0},
	{0, 0, 'D', 0, 0},
}

func moveHandlerDayTwoB(loc types.Coord, move rune) types.Coord {
	newLoc := loc
	switch move {
	case 'L':
		newLoc.X--
	case 'R':
		newLoc.X++
	case 'U':
		newLoc.Y--
	case 'D':
		newLoc.Y++
	}
	if newLoc.X < 0 || newLoc.X > 4 ||
		newLoc.Y < 0 || newLoc.Y > 4 ||
		fancyKeypad[newLoc.Y][newLoc.X] == 0 {
		return loc
	}
	return newLoc
}

func outputHandlerDayTwoB(loc types.Coord) byte {
	return fancyKeypad[loc.Y][loc.X]
}

func DayTwoA(fp *bufio.Reader) string {
	return string(findSolutionDayTwo(
		fp,
		types.Coord{X: 1, Y: 1},
		moveHandlerDayTwoA,
		outputHandlerDayTwoA))
}

func DayTwoB(fp *bufio.Reader) string {
	return string(findSolutionDayTwo(
		fp,
		types.Coord{X: 0, Y: 2},
		moveHandlerDayTwoB,
		outputHandlerDayTwoB))
}
