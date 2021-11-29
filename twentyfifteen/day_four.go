package twentyfifteen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func hashHasLeadingZeros(s *string, n int) bool {
	for i := 0; i < n; i++ {
		if (*s)[i] != '0' {
			return false
		}
	}
	return true
}

func DayFourA(fp *bufio.Reader) string {
	input := utils.ReadSingleString(fp)

	idx := 0
	for {
		s := fmt.Sprintf("%s%d", input, idx)

		hsh := utils.GetMD5StringOfString(s)
		if hashHasLeadingZeros(&hsh, 5) {
			return strconv.Itoa(idx)
		}
		idx++
	}
}

func DayFourB(fp *bufio.Reader) string {
	input := utils.ReadSingleString(fp)

	idx := 0
	for {
		s := fmt.Sprintf("%s%d", input, idx)

		hsh := utils.GetMD5StringOfString(s)
		if hashHasLeadingZeros(&hsh, 6) {
			return strconv.Itoa(idx)
		}
		idx++
	}
}
