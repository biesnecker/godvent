package twentytwenty

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayTen(fp *bufio.Reader) []int {
	adapters := utils.ReadOneIntegerPerLine(fp)
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	return adapters
}

func DayTenA(fp *bufio.Reader) string {
	adapters := readInputDayTen(fp)
	fmt.Println(adapters)

	var ones, threes int

	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		switch diff {
		case 3:
			threes++
		case 1:
			ones++
		}
	}
	return strconv.Itoa(threes * ones)
}

func pathCount(cache map[int]int, adapters []int, idx int) int {
	if idx == len(adapters)-1 {
		return 1
	}
	total := 0
	// The most we can possibly connect to is three ahead, but don't check past
	// the end of the slice, either.
	maxIdx := utils.MinInt(idx+4, len(adapters))
	for i := idx + 1; i < maxIdx; i++ {
		if adapters[i]-adapters[idx] < 4 {
			if cache[i] == 0 {
				subtotal := pathCount(cache, adapters, i)
				cache[i] = subtotal
				total += subtotal
			} else {
				total += cache[i]
			}
		}
	}
	return total
}

func DayTenB(fp *bufio.Reader) string {
	cache := make(map[int]int)
	adapters := readInputDayTen(fp)
	return strconv.Itoa(pathCount(cache, adapters, 0))
}
