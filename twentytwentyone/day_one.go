package twentytwentyone

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func DayOneA(fp *bufio.Reader) string {
	nums := utils.ReadOneIntegerPerLine(fp)
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		}
	}
	return strconv.Itoa(count)
}

func DayOneB(fp *bufio.Reader) string {
	nums := utils.ReadOneIntegerPerLine(fp)
	count := 0
	for i := 0; i < len(nums)-3; i++ {
		if nums[i]+nums[i+1]+nums[i+2] < nums[i+1]+nums[i+2]+nums[i+3] {
			count++
		}
	}
	return strconv.Itoa(count)
}
