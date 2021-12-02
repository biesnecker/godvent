package twentytwenty

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

func readInputDayEleven(fp *bufio.Reader) (map[types.Coord]bool, int, int) {
	var xSize, ySize int
	res := make(map[types.Coord]bool)
	utils.ReadStringsWithIndex(fp, func(x int, s string) {
		xSize = x
		for y, b := range s {
			if b == 'L' {
				res[types.Coord{X: x, Y: y}] = false
			}
			if y > ySize {
				ySize = y
			}
		}
	})
	return res, xSize, ySize
}

type occupiedCalcFn func(map[types.Coord]bool, int, int, types.Coord) int

func getOccupiedAdjacent(m map[types.Coord]bool, xSize int, ySize int, loc types.Coord) int {
	occupiedCount := 0
	for _, adj := range loc.GetSurroundingCoords() {
		if adj.IsInBounds(0, 0, xSize, ySize) {
			if o, ok := m[adj]; ok {
				if o {
					occupiedCount++
				}
			}
		}
	}
	return occupiedCount
}

// Order: N NE E SE S SW W NW
var xOffsets = [8]int{-1, -1, 0, 1, 1, 1, 0, -1}
var yOffsets = [8]int{0, 1, 1, 1, 0, -1, -1, -1}

func getOccupiedLOS(m map[types.Coord]bool, xSize int, ySize int, loc types.Coord) int {
	occupiedCount := 0

	for d := 0; d < 8; d++ {
		xOffset := xOffsets[d]
		yOffset := yOffsets[d]

		step := 1
		for {
			next := types.Coord{X: loc.X + (xOffset * step), Y: loc.Y + (yOffset * step)}
			if !next.IsInBounds(0, 0, xSize, ySize) {
				break
			}
			if v, ok := m[next]; ok {
				if v {
					occupiedCount++
				}
				break
			} else {
				step++
			}
		}
	}

	return occupiedCount
}

func stepDayEleven(
	original map[types.Coord]bool,
	xSize int,
	ySize int,
	flipSize int,
	ofn occupiedCalcFn) (map[types.Coord]bool, bool) {

	changed := false

	newMap := make(map[types.Coord]bool)
	for loc, occupied := range original {
		occupiedCount := ofn(original, xSize, ySize, loc)
		if occupied && occupiedCount >= flipSize {
			newMap[loc] = false
			changed = true
		} else if !occupied && occupiedCount == 0 {
			newMap[loc] = true
			changed = true
		} else {
			// No change.
			newMap[loc] = occupied
		}
	}

	return newMap, changed
}

func countOccupied(m map[types.Coord]bool) int {
	count := 0
	for _, v := range m {
		if v {
			count++
		}
	}
	return count
}

func DayElevenA(fp *bufio.Reader) string {
	m, xSize, ySize := readInputDayEleven(fp)
	for {
		newMap, changed := stepDayEleven(m, xSize, ySize, 4, getOccupiedAdjacent)
		if !changed {
			return strconv.Itoa(countOccupied(newMap))
		} else {
			m = newMap
		}
	}
}

func DayElevenB(fp *bufio.Reader) string {
	m, xSize, ySize := readInputDayEleven(fp)
	for {
		newMap, changed := stepDayEleven(m, xSize, ySize, 5, getOccupiedLOS)
		if !changed {
			return strconv.Itoa(countOccupied(newMap))
		} else {
			m = newMap
		}
	}
}
