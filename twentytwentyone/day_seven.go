package twentytwentyone

import (
	"bufio"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func getAnswerDaySeven(fp *bufio.Reader, fuelCalc func(pos, v int) int) string {
	input := utils.ReadDeliminatedInts(fp, ",")
	min, max := utils.MinMaxIntSlice(input)
	minfuel := math.MaxInt
	for pos := min; pos < max; pos++ {
		fuel := 0
		for _, v := range input {
			fuel += fuelCalc(pos, v)
		}
		if fuel < minfuel {
			minfuel = fuel
		} else {
			break
		}
	}
	return strconv.Itoa(minfuel)
}

func DaySevenA(fp *bufio.Reader) string {
	return getAnswerDaySeven(fp, func(pos, v int) int {
		return utils.IntAbs(pos - v)
	})
}

func DaySevenB(fp *bufio.Reader) string {
	return getAnswerDaySeven(fp, func(pos, v int) int {
		dist := utils.IntAbs(pos - v)
		return (dist * (dist + 1)) / 2
	})
}
