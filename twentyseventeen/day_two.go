package twentyseventeen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func DayTwoA(fp *bufio.Reader) string {
	sum := 0
	utils.ReadDelimitedIntegerStrings(fp, "\t", func(nums []int) {
		min, max := utils.MinMaxIntSlice(nums)
		sum += max - min
	})
	return strconv.Itoa(sum)
}

func DayTwoB(fp *bufio.Reader) string {
	sum := 0
	utils.ReadDelimitedIntegerStrings(fp, "\t", func(nums []int) {
		utils.Combinations(len(nums), 2, func(combo []int) bool {
			v1 := nums[combo[0]]
			v2 := nums[combo[1]]
			if v1 > v2 && v1%v2 == 0 {
				sum += v1 / v2
				return false
			} else if v2 > v1 && v2%v1 == 0 {
				sum += v2 / v1
				return false
			}
			return true
		})
	})
	return strconv.Itoa(sum)
}
