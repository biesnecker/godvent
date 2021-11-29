package twentysixteen

import (
	"bufio"
	"math"

	"github.com/biesnecker/godvent/utils"
)

func getMostAndLeastCommon(letters []rune) (rune, rune) {
	counts := make(map[rune]int)

	for _, letter := range letters {
		counts[letter] += 1
	}

	max := math.MinInt64
	min := math.MaxInt64
	var maxLetter, minLetter rune

	for k, v := range counts {
		if v > max {
			max = v
			maxLetter = k
		}
		if v < min {
			min = v
			minLetter = k
		}
	}

	return maxLetter, minLetter
}

func DaySixA(fp *bufio.Reader) string {
	var columns [][]rune
	utils.ReadStrings(fp, func(s string) {
		for i, c := range s {
			if i >= len(columns) {
				columns = append(columns, make([]rune, 0))
			}
			columns[i] = append(columns[i], c)
		}
	})
	res := make([]rune, 0, len(columns))
	for _, column := range columns {
		max, _ := getMostAndLeastCommon(column)
		res = append(res, max)
	}
	return string(res)
}

func DaySixB(fp *bufio.Reader) string {
	var columns [][]rune
	utils.ReadStrings(fp, func(s string) {
		for i, c := range s {
			if i >= len(columns) {
				columns = append(columns, make([]rune, 0))
			}
			columns[i] = append(columns[i], c)
		}
	})
	res := make([]rune, 0, len(columns))
	for _, column := range columns {
		_, min := getMostAndLeastCommon(column)
		res = append(res, min)
	}
	return string(res)
}
