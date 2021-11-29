package twentyseventeen

import (
	"bufio"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

func readInputDayThree(fp *bufio.Reader) int {
	s := utils.ReadSingleString(fp)
	if n, err := strconv.Atoi(s); err != nil {
		log.Fatalln(err)
	} else {
		return n
	}
	return 0
}

func DayThreeA(fp *bufio.Reader) string {
	target := readInputDayThree(fp)
	var targetCoord types.Coord
	utils.GenerateUlamSpiral(func(idx int, c types.Coord) bool {
		// Problem is 1-indexed but generated spiral is 0-indexed
		if idx+1 == target {
			targetCoord = c
			return false
		}
		return true
	})

	return strconv.Itoa(utils.ManhattanDistance(types.Coord{}, targetCoord))
}

func DayThreeB(fp *bufio.Reader) string {
	target := readInputDayThree(fp)
	seen := make(map[types.Coord]int)
	var solution int

	seen[types.Coord{}] = 1

	utils.GenerateUlamSpiral(func(idx int, c types.Coord) bool {
		if idx == 0 {
			// Skip
			return true
		}

		total := 0
		for _, adjCoord := range c.GetSurroundingCoords() {
			if val, ok := seen[adjCoord]; ok {
				total += val
			}
		}
		if total > target {
			solution = total
			return false
		}
		seen[c] = total
		return true
	})
	return strconv.Itoa(solution)
}
