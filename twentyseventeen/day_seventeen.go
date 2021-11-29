package twentyseventeen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDaySeventeen(fp *bufio.Reader) int {
	return utils.ReadInt(utils.ReadSingleString(fp))
}

func DaySeventeenA(fp *bufio.Reader) string {
	step := readInputDaySeventeen(fp)

	buffer := make([]int, 2048)
	bufferLen := 1
	currentIndex := 0

	for i := 1; i < 2018; i++ {
		nextIndex := ((currentIndex + step) % bufferLen) + 1
		bufferLen++

		// Move everything from the nextIndex forward one position.
		for j := bufferLen; j >= nextIndex; j-- {
			buffer[j] = buffer[j-1]
		}
		buffer[nextIndex] = i
		currentIndex = nextIndex
	}
	return strconv.Itoa(buffer[currentIndex+1])
}

func DaySeventeenB(fp *bufio.Reader) string {
	step := readInputDaySeventeen(fp)

	idx := 0
	answer := 0
	for i := 1; i < 50000000; i++ {
		idx = (idx+step)%i + 1
		if idx == 1 {
			answer = i
		}
	}
	return strconv.Itoa(answer)
}
