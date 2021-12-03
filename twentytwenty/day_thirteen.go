package twentytwenty

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayThirteen(fp *bufio.Reader) (int, []int, []int) {
	var departureTime int
	var offsets []int
	var buses []int
	lineOne := true
	utils.ReadStrings(fp, func(s string) {
		if lineOne {
			lineOne = false
			departureTime = utils.ReadInt(s)
		} else {
			for i, t := range strings.Split(s, ",") {
				if t != "x" {
					offsets = append(offsets, i)
					buses = append(buses, utils.ReadInt(t))
				}
			}

		}
	})
	return departureTime, offsets, buses
}

func DayThirteenA(fp *bufio.Reader) string {
	departureTime, _, buses := readInputDayThirteen(fp)

	minWait := math.MaxInt
	minBus := 0

	for _, bus := range buses {
		wait := bus - departureTime%bus
		if wait < minWait {
			minWait = wait
			minBus = bus
		}
	}
	return strconv.Itoa(minBus * minWait)
}

func DayThirteenB(fp *bufio.Reader) string {
	_, offsets, buses := readInputDayThirteen(fp)

	target := 0
	step := 1

	for i := 0; i < len(buses); i++ {
		for (target+offsets[i])%buses[i] != 0 {
			target += step
		}
		step *= buses[i]
	}
	return strconv.Itoa(target)
}
