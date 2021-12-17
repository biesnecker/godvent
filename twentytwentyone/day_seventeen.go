package twentytwentyone

import (
	"bufio"
	"fmt"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

func inBounds(c types.Coord, b bounds) bool {
	return c.IsInBounds(b.minx, b.miny, b.maxx, b.maxy)
}

func runSimulation(xvel int, yvel int, b bounds) (bool, int) {
	//fmt.Println("vel ", xvel, yvel)
	loc := types.Coord{0, 0}
	maxy := math.MinInt
	for {
		loc = loc.UpBy(yvel)
		loc = loc.RightBy(xvel)
		if xvel > 0 {
			xvel--
		} else if xvel < 0 {
			xvel++
		}
		if loc.Y > maxy {
			maxy = loc.Y
		}
		yvel--
		if inBounds(loc, b) {
			return true, maxy
		} else if (loc.X > b.maxx && xvel > 0) || (loc.X < b.minx && xvel < 0) {
			return false, 0
		} else if loc.Y < b.miny && yvel < 0 {
			return false, 0
		}
	}
}

type bounds struct {
	minx, maxx int
	miny, maxy int
}

func readInputDay17(fp *bufio.Reader) bounds {
	var minx, maxx, miny, maxy int
	s := utils.ReadSingleString(fp)
	fmt.Sscanf(s,
		"target area: x=%d..%d, y=%d..%d",
		&minx, &maxx, &miny, &maxy)
	return bounds{minx, maxx, miny, maxy}
}

func DaySeventeenA(fp *bufio.Reader) string {
	b := readInputDay17(fp)
	maxy := math.MinInt
	for x := 0; x < b.maxx; x++ {
		for y := 0; y < 100; y++ {
			hit, m := runSimulation(x, y, b)
			if hit && m > maxy {
				maxy = m
			}
		}
	}
	return strconv.Itoa(maxy)
}

func DaySeventeenB(fp *bufio.Reader) string {
	type pair struct {
		x, y int
	}

	seen := make(map[pair]bool)
	b := readInputDay17(fp)
	for x := 0; x < b.maxx+1; x++ {
		for y := -100; y < 100; y++ {
			hit, _ := runSimulation(x, y, b)
			if hit {
				seen[pair{x, y}] = true
			}
		}
	}

	return strconv.Itoa(len(seen))
}
