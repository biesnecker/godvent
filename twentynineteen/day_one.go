package twentynineteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func DayOneA(fp *bufio.Reader) string {
	in := utils.ReadOneIntegerPerLine(fp)
	total := 0
	for _, n := range in {
		total += (n / 3) - 2
	}
	return strconv.Itoa(total)
}

func DayOneB(fp *bufio.Reader) string {
	in := utils.ReadOneIntegerPerLine(fp)
	var total, fuel int
	for _, n := range in {
		fuel = (n / 3) - 2
		total += fuel
		for {
			fuel = (fuel / 3) - 2
			if fuel <= 0 {
				break
			}
			total += fuel
		}
	}
	return strconv.Itoa(total)
}
