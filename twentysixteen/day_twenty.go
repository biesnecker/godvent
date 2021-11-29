package twentysixteen

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type rangeD20 struct {
	begin, end int
}

func mergeSortedRanges(rs []rangeD20) []rangeD20 {
	merged := []rangeD20{rs[0]}
	for _, r := range rs[1:] {
		mh := len(merged) - 1
		if r.begin <= merged[mh].end+1 {
			// Merge these together.
			merged[mh].end = utils.MaxInt(merged[mh].end, r.end)
		} else {
			merged = append(merged, r)
		}
	}
	return merged
}

func readInputDayTwenty(fp *bufio.Reader) []rangeD20 {
	var res []rangeD20

	utils.ReadStrings(fp, func(s string) {
		var from, to int
		fmt.Sscanf(s, "%d-%d", &from, &to)
		res = append(res, rangeD20{begin: from, end: to})
	})

	sort.Slice(res, func(i, j int) bool {
		if res[i].begin == res[j].begin {
			return res[i].end < res[j].end
		} else {
			return res[i].begin < res[j].begin
		}
	})

	return mergeSortedRanges(res)
}

func findSolutionDayTwenty(fp *bufio.Reader) (int, int) {
	ranges := readInputDayTwenty(fp)
	head := ranges[0]
	first := 0
	if head.begin == 0 {
		first = head.end + 1
	}
	total := head.begin - 0
	for i := 0; i < len(ranges)-1; i++ {
		total += ranges[i+1].begin - ranges[i].end - 1
	}
	total += math.MaxUint32 - ranges[len(ranges)-1].end

	return first, total
}

func DayTwentyA(fp *bufio.Reader) string {
	first, _ := findSolutionDayTwenty(fp)
	return strconv.Itoa(first)
}

func DayTwentyB(fp *bufio.Reader) string {
	_, total := findSolutionDayTwenty(fp)
	return strconv.Itoa(total)
}
