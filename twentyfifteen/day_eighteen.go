package twentyfifteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type grid [100][100]bool

func isAlive(g *grid, x, y int, force bool) bool {
	if force {
		if (x == 0 && (y == 0 || y == 99)) || (x == 99 && (y == 0 || y == 99)) {
			return true
		}
	}
	isOn := g[x][y]
	count := 0
	for xs := x - 1; xs <= x+1; xs++ {
		for ys := y - 1; ys <= y+1; ys++ {
			if xs == x && ys == y {
				continue
			}
			if xs < 0 || ys < 0 || xs > 99 || ys > 99 {
				continue
			}
			if g[xs][ys] {
				count++
			}
		}
	}
	if isOn {
		return count == 2 || count == 3
	} else {
		return count == 3
	}
}

func stepDayEighteen(original *grid, force bool) *grid {
	var updated grid
	for i := range original {
		for j := range original[i] {
			updated[i][j] = isAlive(original, i, j, force)
		}
	}
	return &updated
}

func countGridDayEighteen(g *grid) int {
	c := 0
	for i := range g {
		for j := range g[i] {
			if g[i][j] {
				c++
			}
		}
	}
	return c
}

func readInputDayEighteen(fp *bufio.Reader) *grid {
	var g grid
	utils.ReadStringsWithIndex(fp, func(i int, s string) {
		for j, c := range s {
			if c == '#' {
				g[i][j] = true
			}
		}
	})
	return &g
}

func DayEighteenA(fp *bufio.Reader) string {
	g := readInputDayEighteen(fp)
	for s := 0; s < 100; s++ {
		g = stepDayEighteen(g, false)
	}

	return strconv.Itoa(countGridDayEighteen(g))
}

func DayEighteenB(fp *bufio.Reader) string {
	g := readInputDayEighteen(fp)
	for s := 0; s < 100; s++ {
		g = stepDayEighteen(g, true)
	}

	return strconv.Itoa(countGridDayEighteen(g))
}
