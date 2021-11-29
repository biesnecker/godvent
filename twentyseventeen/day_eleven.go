package twentyseventeen

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

func findSolutionDayEleven(fp *bufio.Reader) (int, int) {
	directions := strings.Split(utils.ReadSingleString(fp), ",")
	startPos := types.HexCoord{}
	endPos := startPos
	maxDistance := math.MinInt64
	for _, dir := range directions {
		switch dir {
		case "n":
			endPos = endPos.North()
		case "ne":
			endPos = endPos.Northeast()
		case "se":
			endPos = endPos.Southeast()
		case "s":
			endPos = endPos.South()
		case "sw":
			endPos = endPos.Southwest()
		case "nw":
			endPos = endPos.Northwest()
		}

		distance := utils.HexDistance(startPos, endPos)
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return utils.HexDistance(startPos, endPos), maxDistance
}

func DayElevenA(fp *bufio.Reader) string {
	dist, _ := findSolutionDayEleven(fp)
	return strconv.Itoa(dist)
}

func DayElevenB(fp *bufio.Reader) string {
	_, maxDistance := findSolutionDayEleven(fp)
	return strconv.Itoa(maxDistance)
}
