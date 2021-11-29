package twentyfifteen

import (
	"bufio"
	"log"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayTwentyFour(fp *bufio.Reader) ([]int, int) {
	res := make([]int, 0, 64)
	sum := 0
	utils.ReadStrings(fp, func(s string) {
		if n, err := strconv.Atoi(s); err != nil {
			log.Fatalln(err)
		} else {
			res = append(res, n)
			sum += n
		}
	})
	return res, sum
}

func findSolutionDayTwentyFour(packages []int, target int) int {
	groupSize := 1
	minQe := math.MaxInt64
	found := false

	for !found {
		groupSize++
		utils.Combinations(len(packages), groupSize, func(combo []int) bool {
			sum := 0
			product := 1
			for _, i := range combo {
				sum += packages[i]
				product *= packages[i]
			}
			if sum == target {
				found = true
				if product < minQe {
					minQe = product
				}
			}
			return true
		})
	}
	return minQe
}

func DayTwentyFourA(fp *bufio.Reader) string {
	packages, sum := readInputDayTwentyFour(fp)
	minQe := findSolutionDayTwentyFour(packages, sum/3)
	return strconv.Itoa(minQe)
}

func DayTwentyFourB(fp *bufio.Reader) string {
	packages, sum := readInputDayTwentyFour(fp)
	minQe := findSolutionDayTwentyFour(packages, sum/4)
	return strconv.Itoa(minQe)
}
