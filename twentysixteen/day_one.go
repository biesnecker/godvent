package twentysixteen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type directionDayOne struct {
	dir  rune
	step int
}

func readInputDayOne(fp *bufio.Reader) []directionDayOne {
	res := make([]directionDayOne, 0, 64)
	s := utils.ReadSingleString(fp)
	sparts := strings.Split(s, ", ")

	for _, part := range sparts {
		var c rune
		var step int
		fmt.Sscanf(part, "%c%d", &c, &step)
		res = append(res, directionDayOne{dir: c, step: step})
	}
	return res
}

func findSolutionDayOne(fp *bufio.Reader) (int, int) {
	seen := make(map[types.Coord]struct{})

	foundDuplicate := false
	duplicateLocation := types.Coord{}

	steps := readInputDayOne(fp)

	dirs := [4]rune{'N', 'E', 'S', 'W'}

	startingLocation := types.Coord{}
	currentLocation := startingLocation
	currentDir := 0

	for _, step := range steps {
		switch step.dir {
		case 'L':
			currentDir--
			if currentDir < 0 {
				currentDir = 3
			}
		case 'R':
			currentDir++
			if currentDir > 3 {
				currentDir = 0
			}
		}
		var deltaX, deltaY int
		switch dirs[currentDir] {
		case 'N':
			deltaY = 1
		case 'E':
			deltaX = 1
		case 'S':
			deltaY = -1
		case 'W':
			deltaX = -1
		}
		for i := 0; i < step.step; i++ {
			currentLocation.X += deltaX
			currentLocation.Y += deltaY
			if !foundDuplicate {
				if _, ok := seen[currentLocation]; ok {
					foundDuplicate = true
					duplicateLocation = currentLocation
				} else {
					seen[currentLocation] = struct{}{}
				}
			}
		}
	}
	totalDistance := utils.ManhattanDistance(startingLocation, currentLocation)
	duplicateDistance := utils.ManhattanDistance(startingLocation, duplicateLocation)
	return totalDistance, duplicateDistance
}

func DayOneA(fp *bufio.Reader) string {
	dist, _ := findSolutionDayOne(fp)
	return strconv.Itoa(dist)
}

func DayOneB(fp *bufio.Reader) string {
	_, dup := findSolutionDayOne(fp)
	return strconv.Itoa(dup)
}
