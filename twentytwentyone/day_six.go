package twentytwentyone

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func getAnswerDaySix(fp *bufio.Reader, ndays int) int {
	input := utils.ReadDeliminatedInts(fp, ",")
	var days [9]int
	for _, i := range input {
		days[i]++
	}
	for day := 0; day < ndays; day++ {
		var newDays [9]int
		for d := range days {
			if d == 0 {
				newDays[6] += days[d]
				newDays[8] += days[d]
			} else {
				newDays[d-1] += days[d]
			}
		}
		for i := 0; i < 9; i++ {
			days[i] = newDays[i]
		}
	}
	count := 0
	for i := range days {
		count += days[i]
	}
	return count
}

func DaySixA(fp *bufio.Reader) string {
	return strconv.Itoa(getAnswerDaySix(fp, 80))
}

func DaySixB(fp *bufio.Reader) string {
	return strconv.Itoa(getAnswerDaySix(fp, 256))
}
