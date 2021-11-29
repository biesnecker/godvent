package twentyfifteen

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputTwo(
	fp *bufio.Reader,
	handler func(int, int, int, interface{}),
	userData interface{},
) {
	var x, y, z int
	for {
		if _, err := fmt.Fscanf(fp, "%dx%dx%d\n", &x, &y, &z); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			handler(x, y, z, userData)
		}
	}
}

func handlerTwoA(x int, y int, z int, userData interface{}) {
	area := (2 * x * y) + (2 * x * z) + (2 * y * z)
	largestSide := utils.MinInt(x*y, utils.MinInt(y*z, x*z))

	var total = userData.(*int)
	*total += area + largestSide
}

func perimeter(x, y int) int {
	return (2 * x) + (2 * y)
}

func handlerTwoB(x, y, z int, userData interface{}) {
	var total = userData.(*int)
	*total += x * y * z
	*total += utils.MinInt(
		perimeter(x, y),
		utils.MinInt(perimeter(x, y), perimeter(y, z)))
}

func DayTwoA(fp *bufio.Reader) string {
	total := 0
	readInputTwo(fp, handlerTwoA, &total)
	return strconv.Itoa(total)
}

func DayTwoB(fp *bufio.Reader) string {
	total := 0
	readInputTwo(fp, handlerTwoB, &total)
	return strconv.Itoa(total)
}
