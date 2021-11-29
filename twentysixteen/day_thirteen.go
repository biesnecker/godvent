package twentysixteen

import (
	"bufio"
	"log"
	"math/bits"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type positionD13 struct {
	pos   types.Coord
	steps int
}

func isWall(favorite uint, c types.Coord) bool {
	v := uint(c.X*c.X+3*c.X+2*c.X*c.Y+c.Y+c.Y*c.Y) + favorite
	return bits.OnesCount(v)%2 == 1
}

func findSolutionDayThirteen(favorite uint, shouldBreak func(positionD13) bool) (int, int) {
	q := make(chan positionD13, 100)

	seen := make(map[types.Coord]struct{})

	q <- positionD13{pos: types.Coord{X: 1, Y: 1}, steps: 0}
	seen[types.Coord{X: 1, Y: 1}] = struct{}{}

	withinFifty := 0

complete:
	for {
		select {
		case p := <-q:
			if p.steps <= 50 {
				withinFifty++
			}
			if shouldBreak(p) {
				return p.steps, withinFifty
			}
			for dir := 0; dir < 4; dir++ {
				var newPos types.Coord
				switch dir {
				case 0:
					newPos = p.pos.Up()
				case 1:
					newPos = p.pos.Down()
				case 2:
					newPos = p.pos.Left()
				case 3:
					newPos = p.pos.Right()
				}
				if newPos.IsNegative() {
					continue
				}
				if isWall(favorite, newPos) {
					continue
				}
				if _, found := seen[newPos]; !found {
					seen[newPos] = struct{}{}
					q <- positionD13{pos: newPos, steps: p.steps + 1}
				}
			}
		default:
			break complete
		}
	}
	return 0, 0
}

func DayThirteenA(fp *bufio.Reader) string {
	fs := utils.ReadSingleString(fp)
	if favorite, err := strconv.Atoi(fs); err != nil {
		log.Fatalln(err)
	} else {
		steps, _ := findSolutionDayThirteen(
			uint(favorite),
			func(p positionD13) bool {
				return p.pos.X == 31 && p.pos.Y == 39
			})
		return strconv.Itoa(steps)
	}
	return ""
}

func DayThirteenB(fp *bufio.Reader) string {
	fs := utils.ReadSingleString(fp)
	if favorite, err := strconv.Atoi(fs); err != nil {
		log.Fatalln(err)
	} else {
		_, withinFifty := findSolutionDayThirteen(
			uint(favorite),
			func(p positionD13) bool {
				return p.steps > 50
			})
		return strconv.Itoa(withinFifty)
	}
	return ""
}
