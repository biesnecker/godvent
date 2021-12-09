package twentyeighteen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func DayOneA(fp *bufio.Reader) string {
	total := 0
	utils.ReadStrings(fp, func(s string) {
		var n int
		fmt.Sscanf(s, "%d", &n)
		total += n
	})
	return strconv.Itoa(total)
}

func DayOneB(fp *bufio.Reader) string {
	total := 0
	seen := make(map[int]struct{})
	nums := utils.ReadOneIntegerPerLine(fp)
	idx := 0
	for {
		i := idx % len(nums)
		total += nums[i]
		if _, ok := seen[total]; ok {
			return strconv.Itoa(total)
		} else {
			seen[total] = struct{}{}
		}
		idx++
	}

}
