package twentytwenty

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

func readInputDayThree(fp *bufio.Reader) (map[types.Coord]bool, int, int) {
	res := make(map[types.Coord]bool)
	rowLen := 0
	nRows := 0
	utils.ReadStringsWithIndex(fp, func(i int, s string) {
		rowLen = len(s)
		nRows++
		for j, v := range s {
			if v == '#' {
				res[types.Coord{X: i, Y: j}] = true
			}
		}
	})
	return res, rowLen, nRows
}

func traverse(m map[types.Coord]bool, rowLen, nRows, xstep, ystep int) int {
	res := 0
	for xpos, ypos := 0, 0; xpos < nRows; {
		if m[types.Coord{X: xpos, Y: ypos}] {
			res++
		}
		ypos = (ypos + ystep) % rowLen
		xpos += xstep
	}
	return res
}

func DayThreeA(fp *bufio.Reader) string {
	m, rowLen, nRows := readInputDayThree(fp)
	res := traverse(m, rowLen, nRows, 1, 3)
	return strconv.Itoa(res)
}

func DayThreeB(fp *bufio.Reader) string {
	m, rowLen, nRows := readInputDayThree(fp)
	res := traverse(m, rowLen, nRows, 1, 1) *
		traverse(m, rowLen, nRows, 1, 3) *
		traverse(m, rowLen, nRows, 1, 5) *
		traverse(m, rowLen, nRows, 1, 7) *
		traverse(m, rowLen, nRows, 2, 1)
	return strconv.Itoa(res)
}
