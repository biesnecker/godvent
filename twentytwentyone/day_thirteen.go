package twentytwentyone

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type foldD13 struct {
	dir byte
	idx int
}

func readInputDayThirteen(fp *bufio.Reader) (map[types.Coord]bool, []foldD13) {
	m := make(map[types.Coord]bool)
	f := make([]foldD13, 0, 16)

	isCoordinates := true
	utils.ReadStrings(fp, func(s string) {
		if len(s) == 0 {
			isCoordinates = false
		} else if isCoordinates {
			var x, y int
			fmt.Sscanf(s, "%d,%d", &x, &y)
			m[types.Coord{X: x, Y: y}] = true
		} else {
			var dir byte
			var idx int
			fmt.Sscanf(s, "fold along %c=%d", &dir, &idx)
			f = append(f, foldD13{dir: dir, idx: idx})
		}
	})

	return m, f
}

func doFoldsD13(m map[types.Coord]bool, folds []foldD13) {
	for _, fold := range folds {
		if fold.dir == 'x' {
			for loc := range m {
				if loc.X > fold.idx {
					nx := fold.idx - (loc.X - fold.idx)
					nloc := types.Coord{X: nx, Y: loc.Y}
					m[nloc] = true
					delete(m, loc)
				}
			}
		} else {
			for loc := range m {
				if loc.Y > fold.idx {
					ny := fold.idx - (loc.Y - fold.idx)
					nloc := types.Coord{X: loc.X, Y: ny}
					m[nloc] = true
					delete(m, loc)
				}
			}
		}
	}
}

func DayThirteenA(fp *bufio.Reader) string {
	m, f := readInputDayThirteen(fp)
	doFoldsD13(m, f[:1])
	return strconv.Itoa(len(m))
}

func DayThirteenB(fp *bufio.Reader) string {
	m, f := readInputDayThirteen(fp)
	doFoldsD13(m, f)

	var maxx, maxy int
	for k := range m {
		if k.X > maxx {
			maxx = k.X
		}
		if k.Y > maxy {
			maxy = k.Y
		}
	}

	maxx++
	maxy++

	display := make([]string, maxy)
	for y := 0; y < maxy; y++ {
		row := make([]byte, maxx)
		for x := 0; x < maxx; x++ {
			if m[types.Coord{X: x, Y: y}] {
				row[x] = '#'
			} else {
				row[x] = '.'
			}
		}
		display[y] = string(row)
	}
	return strings.Join(display, "\n")
}
