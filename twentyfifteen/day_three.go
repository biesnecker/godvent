package twentyfifteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

func updateSleighLocation(c rune, loc *types.Coord) {
	switch c {
	case '^':
		loc.Y++
	case 'v':
		loc.Y--
	case '<':
		loc.X--
	case '>':
		loc.X++
	}
}

func DayThreeA(fp *bufio.Reader) string {
	santaLoc := types.Coord{X: 0, Y: 0}
	visited := make(map[types.Coord]struct{})

	utils.ReadChars(fp, func(c rune, _ int) bool {
		updateSleighLocation(c, &santaLoc)
		visited[santaLoc] = struct{}{}
		return true
	})

	return strconv.Itoa(len(visited))
}

func DayThreeB(fp *bufio.Reader) string {

	santaLoc := types.Coord{X: 0, Y: 0}
	robotLoc := types.Coord{X: 0, Y: 0}
	visited := make(map[types.Coord]struct{})
	isSanta := true

	visited[santaLoc] = struct{}{}
	utils.ReadChars(fp, func(c rune, _ int) bool {
		var loc *types.Coord
		if isSanta {
			loc = &santaLoc
		} else {
			loc = &robotLoc
		}
		updateSleighLocation(c, loc)
		visited[*loc] = struct{}{}

		isSanta = !isSanta
		return true
	})

	return strconv.Itoa(len(visited))
}
