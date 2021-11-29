package twentyfifteen

import (
	"bufio"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func sumFactors(factors map[int]int) int {
	total := 1
	for base, exp := range factors {
		subtotal := base
		s := base
		for i := 0; i < exp-1; i++ {
			s *= base
			subtotal += s
		}
		total *= (subtotal + 1)
	}
	return total
}

func DayTwentyA(fp *bufio.Reader) string {
	target, err := strconv.Atoi(utils.ReadSingleString(fp))
	if err != nil {
		log.Fatalln(err)
	}

	fac := utils.NewFactorizer()

	step := 5040
	idx := 0
	result := 0
	for result < target {
		idx += step
		result = sumFactors(fac.PrimeFactorMap(idx)) * 10
	}
	return strconv.Itoa(idx)
}

func DayTwentyB(fp *bufio.Reader) string {
	target, err := strconv.Atoi(utils.ReadSingleString(fp))
	if err != nil {
		log.Fatalln(err)
	}

	step := 5040
	i := 0
	var result int
	for result < target {
		result = 0
		i += step
		for j := i / 50; j <= i; j++ {
			if i%j == 0 {
				result += j * 11
			}
		}
	}
	return strconv.Itoa(i)
}
