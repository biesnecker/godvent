package twentyseventeen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayOne(fp *bufio.Reader) []int {
	var res []int
	s := utils.ReadSingleString(fp)
	for _, c := range s {
		res = append(res, int(c-'0'))
	}
	return res
}

func DayOneA(fp *bufio.Reader) string {
	nums := readInputDayOne(fp)
	sum := 0
	for i, n := range nums {
		var nextIdx int
		if i == len(nums)-1 {
			nextIdx = 0
		} else {
			nextIdx = i + 1
		}
		if n == nums[nextIdx] {
			sum += n
		}
	}
	return strconv.Itoa(sum)
}

func DayOneB(fp *bufio.Reader) string {
	nums := readInputDayOne(fp)
	ln := len(nums)
	half := ln / 2
	sum := 0
	for i, n := range nums {
		if n == nums[(i+half)%ln] {
			sum += n
		}
	}
	return strconv.Itoa(sum)
}
