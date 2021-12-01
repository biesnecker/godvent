package twentytwentyone

import (
	"bufio"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayOne(fp *bufio.Reader) []int {
	v, err := utils.ReadOneIntegerPerLine(fp)
	if err != nil {
		log.Fatalln(err)
	}
	return v
}

func DayOneA(fp *bufio.Reader) string {
	nums := readInputDayOne(fp)
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		}
	}
	return strconv.Itoa(count)
}

func DayOneB(fp *bufio.Reader) string {
	nums := readInputDayOne(fp)
	count := 0
	for i := 0; i < len(nums)-3; i++ {
		if nums[i]+nums[i+1]+nums[i+2] < nums[i+1]+nums[i+2]+nums[i+3] {
			count++
		}
	}
	return strconv.Itoa(count)
}
