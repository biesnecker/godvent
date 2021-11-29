package twentyfifteen

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func getIndexForCityName(names *[]string, target string) int {
	for i, name := range *names {
		if name == target {
			return i
		}
	}
	*names = append(*names, target)
	return len(*names) - 1
}

func collectDistances(names *[]string, distances *[8][8]int) func(string) {
	return func(s string) {
		var start, end string
		var dist int
		if _, err := fmt.Sscanf(s, "%s to %s = %d", &start, &end, &dist); err != nil {
			log.Fatal(err)
		}
		cityOneId := getIndexForCityName(names, start)
		cityTwoId := getIndexForCityName(names, end)
		distances[cityOneId][cityTwoId] = dist
		distances[cityTwoId][cityOneId] = dist
	}
}

func permuteRoutes(
	distances *[8][8]int,
	lastCity int,
	distance int,
	mask uint8,
	target *int,
	isBetter func(int, int) bool,
) {
	if mask == 0xff {
		if isBetter(distance, *target) {
			*target = distance
		}
		return
	}

	for i := 0; i < 8; i++ {
		check := uint8(1) << i
		if mask&check > 0 {
			// Already done this one.
			continue
		}
		var d int
		if lastCity != -1 {
			d = distances[lastCity][i]
		}
		permuteRoutes(distances, i, distance+d, mask|check, target, isBetter)
	}

}

func DayNineA(fp *bufio.Reader) string {
	names := make([]string, 0, 10)
	var distances [8][8]int
	utils.ReadStrings(fp, collectDistances(&names, &distances))

	bestDistance := math.MaxInt32
	permuteRoutes(&distances, -1, 0, 0, &bestDistance, func(candidate, target int) bool {
		return candidate < target
	})
	return strconv.Itoa(bestDistance)
}

func DayNineB(fp *bufio.Reader) string {
	names := make([]string, 0, 10)
	var distances [8][8]int
	utils.ReadStrings(fp, collectDistances(&names, &distances))

	bestDistance := math.MinInt32
	permuteRoutes(&distances, -1, 0, 0, &bestDistance, func(candidate, target int) bool {
		return candidate > target
	})
	return strconv.Itoa(bestDistance)
}
