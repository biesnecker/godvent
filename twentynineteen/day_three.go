package twentynineteen

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type commandDay3 struct {
	direction rune
	count     int
}

type stepD3 struct {
	coord types.Coord
	steps int
}

func parseLineDayThree(s string) []commandDay3 {
	line := make([]commandDay3, 0, 128)
	for _, command := range strings.Split(s, ",") {
		var c commandDay3
		fmt.Sscanf(command, "%c%d", &c.direction, &c.count)
		line = append(line, c)
	}
	return line
}

func readInputDayThree(fp *bufio.Reader) ([]commandDay3, []commandDay3) {
	lines := utils.ReadStringsAsSlice(fp)

	if len(lines) != 2 {
		log.Fatal("too many lines")
	}

	return parseLineDayThree(lines[0]), parseLineDayThree(lines[1])
}

func stepsForLine(line []commandDay3) map[types.Coord]int {
	res := make(map[types.Coord]int)
	x := 0
	y := 0
	steps := 0
	for _, command := range line {
		var deltax, deltay int
		switch command.direction {
		case 'U':
			deltay = 1
		case 'D':
			deltay = -1
		case 'R':
			deltax = 1
		case 'L':
			deltax = -1
		}
		for i := 0; i < command.count; i++ {
			x += deltax
			y += deltay
			steps++
			c := types.Coord{X: x, Y: y}
			if _, ok := res[c]; !ok {
				res[c] = steps
			}
		}
	}
	return res
}

func getCoordMaps(fp *bufio.Reader) (map[types.Coord]int, map[types.Coord]int) {
	lineOne, lineTwo := readInputDayThree(fp)
	c1 := make(chan map[types.Coord]int)
	c2 := make(chan map[types.Coord]int)

	go func() {
		c1 <- stepsForLine(lineOne)
	}()
	go func() {
		c2 <- stepsForLine(lineTwo)
	}()

	return <-c1, <-c2
}

func DayThreeA(fp *bufio.Reader) string {
	coordsOne, coordsTwo := getCoordMaps(fp)

	mindist := math.MaxInt

	start := types.Coord{}

	for c := range coordsOne {
		if _, ok := coordsTwo[c]; ok {
			d := utils.ManhattanDistance(start, c)
			if d < mindist {
				mindist = d
			}
		}
	}

	return strconv.Itoa(mindist)
}

func DayThreeB(fp *bufio.Reader) string {
	coordsOne, coordsTwo := getCoordMaps(fp)

	minsteps := math.MaxInt

	for c, s1 := range coordsOne {
		if s2, ok := coordsTwo[c]; ok {
			d := s1 + s2
			if d < minsteps {
				minsteps = d
			}
		}
	}

	return strconv.Itoa(minsteps)
}
