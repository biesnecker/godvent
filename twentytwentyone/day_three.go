package twentytwentyone

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func highLowBinaryStrings(nums []string) (string, string) {
	digits := make([]int, len(nums[0]))

	for _, num := range nums {
		for idx, bit := range num {
			if bit == '1' {
				digits[idx]++
			} else {
				digits[idx]--
			}
		}
	}
	var high []byte
	var low []byte
	for _, i := range digits {
		if i >= 0 {
			// Tie behavior is the same as 1 being the most common.
			high = append(high, '1')
			low = append(low, '0')
		} else {
			high = append(high, '0')
			low = append(low, '1')
		}
	}

	return string(high), string(low)
}

func DayThreeA(fp *bufio.Reader) string {
	input := utils.ReadStringsAsSlice(fp)

	gammaString, epsilonString := highLowBinaryStrings(input)
	gamma := utils.ParseBinaryString(gammaString)
	epsilon := utils.ParseBinaryString(epsilonString)
	return strconv.Itoa(gamma * epsilon)
}

func findCandidate(candidates []string, takeHigh bool) string {
	i := 0
	for len(candidates) > 1 {
		high, low := highLowBinaryStrings(candidates)
		var match string
		if takeHigh {
			match = high
		} else {
			match = low
		}

		candidates = utils.FilterStrings(
			candidates,
			func(s string) bool {
				return s[i] == match[i]
			})

		i++
	}
	return candidates[0]
}

func DayThreeB(fp *bufio.Reader) string {
	input := utils.ReadStringsAsSlice(fp)

	oxygen := findCandidate(input, true)
	co2 := findCandidate(input, false)

	oi := utils.ParseBinaryString(oxygen)
	ci := utils.ParseBinaryString(co2)

	return strconv.Itoa(oi * ci)
}
