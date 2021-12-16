package twentytwentyone

import (
	"bufio"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/types/queue"
	"github.com/biesnecker/godvent/utils"
)

func readInputDayFifteen(fp *bufio.Reader) (int, int, map[types.Coord]int) {
	res := make(map[types.Coord]int)
	var maxx, maxy int
	utils.ReadStringsWithIndex(fp, func(y int, s string) {
		if y > maxy {
			maxy = y
		}
		for x := range s {
			if x > maxx {
				maxx = x
			}
			res[types.Coord{X: x, Y: y}] = int(s[x] - '0')
		}
	})
	return maxx, maxy, res

}

func DayFifteenA(fp *bufio.Reader) string {
	maxx, maxy, input := readInputDayFifteen(fp)

	type step struct {
		loc  types.Coord
		risk int
	}

	target := types.Coord{X: maxx, Y: maxy}

	q := queue.New()
	q.Push(step{loc: types.Coord{X: 0, Y: 0}, risk: 0})

	minrisk := math.MaxInt

	risks := make(map[types.Coord]int)

	for !q.Empty() {
		s := q.Pop().(step)

		if s.loc == target {
			if minrisk > s.risk {
				minrisk = s.risk
				continue
			}
		}

	adjloop:
		for _, nextloc := range s.loc.GetAdjacentCoords() {
			if nextrisk, ok := input[nextloc]; ok {
				nextrisk += s.risk
				if nextrisk >= minrisk {
					continue adjloop
				}
				if c, ok := risks[nextloc]; ok && c <= nextrisk {
					continue adjloop
				}
				risks[nextloc] = nextrisk
				q.Push(step{loc: nextloc, risk: nextrisk})
			}
		}

	}
	return strconv.Itoa(minrisk)
}

func DayFifteenB(fp *bufio.Reader) string {
	maxx, maxy, input := readInputDayFifteen(fp)

	maxxe := (maxx+1)*5 - 1
	maxye := (maxy+1)*5 - 1

	type step struct {
		loc  types.Coord
		risk int
	}

	target := types.Coord{X: maxxe, Y: maxye}

	q := queue.New()
	q.Push(step{loc: types.Coord{X: 0, Y: 0}, risk: 0})

	minrisk := math.MaxInt

	risks := make(map[types.Coord]int)

	for !q.Empty() {
		s := q.Pop().(step)

		if s.loc == target {
			if minrisk > s.risk {
				minrisk = s.risk
				continue
			}
		}

	adjloop:
		for _, actualnextloc := range s.loc.GetAdjacentCoords() {
			if actualnextloc.X < 0 || actualnextloc.X > maxxe ||
				actualnextloc.Y < 0 || actualnextloc.Y > maxye {
				continue
			}
			nx := actualnextloc.X % (maxx + 1)
			ny := actualnextloc.Y % (maxy + 1)
			nextloc := types.Coord{X: nx, Y: ny}
			xadd := actualnextloc.X / (maxx + 1)
			yadd := actualnextloc.Y / (maxy + 1)
			add := xadd + yadd
			if baserisk, ok := input[nextloc]; ok {
				nextrisk := baserisk + add
				if nextrisk > 9 {
					nextrisk = nextrisk%10 + 1
				}
				nextrisk += s.risk
				if c, ok := risks[actualnextloc]; ok && c <= nextrisk {
					continue adjloop
				}
				risks[actualnextloc] = nextrisk
				q.Push(step{loc: actualnextloc, risk: nextrisk})
			}
		}

	}
	return strconv.Itoa(minrisk)
}
