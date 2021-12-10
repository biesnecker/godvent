package twentynineteen

import (
	"bufio"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

const layerSize int = 25 * 6

func readInputDayEight(fp *bufio.Reader) []int {
	res := make([]int, 0, 1024)
	for _, b := range utils.ReadSingleString(fp) {
		res = append(res, int(b-'0'))
	}
	return res
}

func DayEightA(fp *bufio.Reader) string {
	input := readInputDayEight(fp)

	minZeros := math.MaxInt
	var res int
	var counts [3]int
	for i, v := range input {
		switch v {
		case 0:
			counts[0]++
		case 1:
			counts[1]++
		case 2:
			counts[2]++
		}
		if i%layerSize == layerSize-1 {
			if counts[0] < minZeros {
				minZeros = counts[0]
				res = counts[1] * counts[2]
			}
			counts[0] = 0
			counts[1] = 0
			counts[2] = 0
		}
	}
	return strconv.Itoa(res)
}

func DayEightB(fp *bufio.Reader) string {
	input := readInputDayEight(fp)

	var image [layerSize]int
	for i := range image {
		image[i] = 2
	}

	for i, v := range input {
		idx := i % layerSize
		if image[idx] != 2 {
			continue
		}
		image[idx] = v
	}

	display := make([]byte, 0, len(image)+6)
	for i, p := range image {
		if i > 0 && i%25 == 0 {
			display = append(display, '\n')
		}
		if p == 1 {
			display = append(display, '#')
		} else {
			display = append(display, ' ')
		}
	}

	return string(display)
}
