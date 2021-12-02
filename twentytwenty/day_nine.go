package twentytwenty

import (
	"bufio"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

const windowSize int = 25

func findInvalidNumber(nums []int) int {
	wvalues := make(map[int]struct{})
	for _, value := range nums[:windowSize] {
		wvalues[value] = struct{}{}
	}

	for i, value := range nums[windowSize:] {
		found := false
		for k := range wvalues {
			target := value - k
			if _, ok := wvalues[target]; ok {
				found = true
				break
			}
		}
		if !found {
			return value
		}

		// Remove the first value from the map and
		// add the current one.
		delete(wvalues, nums[i])
		wvalues[value] = struct{}{}
	}
	log.Fatalln("Didn't find the solution")
	return 0
}

func DayNineA(fp *bufio.Reader) string {
	nums := utils.ReadOneIntegerPerLine(fp)
	return strconv.Itoa(findInvalidNumber(nums))
}

func DayNineB(fp *bufio.Reader) string {
	nums := utils.ReadOneIntegerPerLine(fp)
	targetNumber := findInvalidNumber(nums)
	tailIdx := 0
	headIdx := 0
	total := 0
	for {
		if total < targetNumber {
			// Add the head to the total
			if headIdx == len(nums) {
				break
			}
			total += nums[headIdx]
			headIdx++
		} else if total > targetNumber {
			// Remove the tail from the total
			if tailIdx == headIdx {
				break
			}
			total -= nums[tailIdx]
			tailIdx++
		} else {
			// We've found the answer
			min, max := utils.MinMaxIntSlice(nums[tailIdx:headIdx])
			return strconv.Itoa(min + max)
		}
	}
	log.Fatalln("Didn't find solution")
	return ""
}
