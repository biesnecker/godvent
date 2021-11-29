package twentyfifteen

import (
	"bufio"
	"log"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type daySeventeenState struct {
	numCombos    int
	minBucket    int
	numMinCombos int
}

func stepDaySeventeen(
	buckets []int,
	goal, total, numBuckets, lastBucket int,
	result *daySeventeenState,
) {
	if total == goal {
		result.numCombos++
		if numBuckets == result.minBucket {
			result.numMinCombos++
		} else if numBuckets < result.minBucket {
			result.minBucket = numBuckets
			result.numMinCombos = 1
		}
		return
	}
	for i := lastBucket + 1; i < len(buckets); i++ {
		stepDaySeventeen(buckets, goal, total+buckets[i], numBuckets+1, i, result)
	}

}

func findSolutionDaySeventeen(fp *bufio.Reader, goal int) (int, int) {
	buckets := make([]int, 0, 50)
	utils.ReadStrings(fp, func(s string) {
		if val, err := strconv.Atoi(s); err != nil {
			log.Fatal(err)
		} else {
			buckets = append(buckets, val)
		}
	})
	state := daySeventeenState{
		numCombos:    0,
		minBucket:    math.MaxInt32,
		numMinCombos: 0}
	stepDaySeventeen(buckets, goal, 0, 0, -1, &state)
	return state.numCombos, state.numMinCombos
}

func DaySeventeenA(fp *bufio.Reader) string {
	n, _ := findSolutionDaySeventeen(fp, 150)
	return strconv.Itoa(n)
}

func DaySeventeenB(fp *bufio.Reader) string {
	_, n := findSolutionDaySeventeen(fp, 150)
	return strconv.Itoa(n)
}
