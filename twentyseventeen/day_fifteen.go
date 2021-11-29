package twentyseventeen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type genD15 struct {
	val, factor int
}

func (g *genD15) Next() int {
	g.val = (g.val * g.factor) % 2147483647
	return g.val & 0xffff
}

func (g *genD15) NextOfMultiple(m int) int {
	nextVal := g.val
	for {
		nextVal = (nextVal * g.factor) % 2147483647
		if nextVal%m == 0 {
			g.val = nextVal
			break
		}
	}
	return g.val & 0xffff
}

func readInputDayFifteen(fp *bufio.Reader) (int, int) {
	var vals []int
	var c rune
	var i int
	utils.ReadStrings(fp, func(s string) {
		fmt.Sscanf(s, "Generator %c starts with %d", &c, &i)
		vals = append(vals, i)
	})
	return vals[0], vals[1]
}

const (
	aFactor int = 16807
	bFactor int = 48271
)

func DayFifteenA(fp *bufio.Reader) string {
	a, b := readInputDayFifteen(fp)

	genA := genD15{val: a, factor: aFactor}
	genB := genD15{val: b, factor: bFactor}

	total := 0
	for i := 0; i < 40000000; i++ {
		if genA.Next() == genB.Next() {
			total++
		}
	}
	return strconv.Itoa(total)
}

func DayFifteenB(fp *bufio.Reader) string {
	a, b := readInputDayFifteen(fp)

	genA := genD15{val: a, factor: aFactor}
	genB := genD15{val: b, factor: bFactor}

	total := 0
	for i := 0; i < 5000000; i++ {
		if genA.NextOfMultiple(4) == genB.NextOfMultiple(8) {
			total++
		}
	}
	return strconv.Itoa(total)
}
