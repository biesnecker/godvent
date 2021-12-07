package twentytwentyone

import (
	"bufio"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func DaySevenA(fp *bufio.Reader) string {
	input := utils.ReadDeliminatedInts(fp, ",")
	max := utils.MaxIntSlice(input)
	minfuel := math.MaxInt
	for pos := 0; pos < max; pos++ {
		fuel := 0
		for _, v := range input {
			fuel += utils.IntAbs(pos - v)
		}
		if fuel < minfuel {
			minfuel = fuel
		}
	}
	return strconv.Itoa(minfuel)
}

func DaySevenB(fp *bufio.Reader) string {
	input := utils.ReadDeliminatedInts(fp, ",")
	max := utils.MaxIntSlice(input)
	minfuel := math.MaxInt
	for pos := 0; pos < max; pos++ {
		fuel := 0
		for _, v := range input {
			diff := utils.IntAbs(pos - v)
			fuel += (diff * (diff + 1)) / 2
		}
		if fuel < minfuel {
			minfuel = fuel
		}
	}
	return strconv.Itoa(minfuel)
}
