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
		d := day % 9
		days[(day+7)%9] += days[d]
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
