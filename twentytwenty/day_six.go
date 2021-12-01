package twentytwenty

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/types/set"
	"github.com/biesnecker/godvent/utils"
)

func readInputDaySix(fp *bufio.Reader) [][]string {
	var res [][]string
	res = append(res, make([]string, 0))
	utils.ReadStrings(fp, func(s string) {
		if len(s) == 0 {
			res = append(res, make([]string, 0))
			return
		}
		res[len(res)-1] = append(res[len(res)-1], s)
	})
	return res
}

func DaySixA(fp *bufio.Reader) string {
	input := readInputDaySix(fp)
	var counts []int
	for _, g := range input {
		seen := set.New()
		for _, s := range g {
			for _, b := range s {
				seen.Insert(b)
			}

		}
		counts = append(counts, seen.Len())
	}
	return strconv.Itoa(utils.SumOfIntSlice(counts))
}

func DaySixB(fp *bufio.Reader) string {
	input := readInputDaySix(fp)
	var counts []int
	for _, g := range input {
		seen := make(map[rune]int)
		for _, s := range g {
			for _, b := range s {
				seen[b] += 1
			}

		}
		count := 0
		for _, v := range seen {
			if v == len(g) {
				count++
			}
		}
		counts = append(counts, count)
	}
	return strconv.Itoa(utils.SumOfIntSlice(counts))
}
