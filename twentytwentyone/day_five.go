package twentytwentyone

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type lineD5 struct {
	start, end types.Coord
}

func readInputDayFive(fp *bufio.Reader) []lineD5 {
	var res []lineD5
	utils.ReadStrings(fp, func(s string) {
		var x1, y1, x2, y2 int
		fmt.Sscanf(s, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		res = append(
			res,
			lineD5{
				types.Coord{X: x1, Y: y1},
				types.Coord{X: x2, Y: y2}})
	})
	return res
}

func isStraight(line *lineD5) bool {
	return line.start.X == line.end.X || line.start.Y == line.end.Y
}

func drawLinesOnMap(m map[types.Coord]int, line *lineD5) {
	xdelta := 0
	ydelta := 0
	if line.start.X > line.end.X {
		xdelta = -1
	} else if line.start.X < line.end.X {
		xdelta = 1
	}
	if line.start.Y > line.end.Y {
		ydelta = -1
	} else if line.start.Y < line.end.Y {
		ydelta = 1
	}
	current := line.start
	for current != line.end {
		m[current]++
		current.X += xdelta
		current.Y += ydelta
	}
	m[current]++
}

func countOverlaps(m map[types.Coord]int) int {
	count := 0
	for _, v := range m {
		if v > 1 {
			count++
		}
	}
	return count
}

func DayFiveA(fp *bufio.Reader) string {
	input := readInputDayFive(fp)

	m := make(map[types.Coord]int)

	for _, line := range input {
		if !isStraight(&line) {
			continue
		}
		drawLinesOnMap(m, &line)
	}
	return strconv.Itoa(countOverlaps(m))
}

func DayFiveB(fp *bufio.Reader) string {
	input := readInputDayFive(fp)

	m := make(map[types.Coord]int)

	for _, line := range input {
		drawLinesOnMap(m, &line)
	}
	return strconv.Itoa(countOverlaps(m))
}
