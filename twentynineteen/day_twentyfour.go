package twentynineteen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

func readInputDay24(fp *bufio.Reader) map[types.Coord]bool {
	res := make(map[types.Coord]bool)
	utils.ReadStringsWithIndex(fp, func(y int, s string) {
		for x := range s {
			if s[x] == '#' {
				res[types.Coord{X: x, Y: y}] = true
			}
		}
	})
	return res
}

func boardReprD24(m map[types.Coord]bool) uint32 {
	repr := uint32(0)

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if m[types.Coord{X: x, Y: y}] {
				repr |= (uint32(1) << ((y * 5) + x))
			}
		}
	}

	return repr
}

func scoreBoard(m map[types.Coord]bool) int {
	score := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if m[types.Coord{X: x, Y: y}] {
				score += utils.IntPow(2, (y*5)+x)
			}
		}
	}
	return score
}

func DayTwentyFourA(fp *bufio.Reader) string {
	currentMap := readInputDay24(fp)

	seen := make(map[uint32]bool)
outerloop:
	for {
		newMap := make(map[types.Coord]bool)
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				loc := types.Coord{X: x, Y: y}
				locAlive := currentMap[loc]
				adjcount := 0
				for _, other := range loc.GetAdjacentCoords() {
					if currentMap[other] {
						adjcount++
					}
				}
				if locAlive && adjcount != 1 {
					locAlive = false
				} else if !locAlive && (adjcount == 1 || adjcount == 2) {
					locAlive = true
				}
				if locAlive {
					newMap[loc] = true
				}
			}
		}
		currentMap = newMap
		repr := boardReprD24(currentMap)

		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if currentMap[types.Coord{X: x, Y: y}] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println("")
		}
		fmt.Println("")

		if seen[repr] {
			break outerloop
		} else {
			seen[repr] = true
		}
	}

	return strconv.Itoa(scoreBoard(currentMap))
}

func DayTwentyFourB(fp *bufio.Reader) string {
	return ""
}
