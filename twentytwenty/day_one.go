package twentytwenty

import (
	"bufio"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayOne(fp *bufio.Reader) []int {
	var res []int
	var err error
	if res, err = utils.ReadOneIntegerPerLine(fp); err != nil {
		log.Fatalln(err)
	}
	return res
}

const target int = 2020

func DayOneA(fp *bufio.Reader) string {
	nums := readInputDayOne(fp)

	nmap := make(map[int]struct{})

	for _, n := range nums {
		t := target - n
		if _, ok := nmap[t]; ok {
			return strconv.Itoa(n * t)
		} else {
			nmap[n] = struct{}{}
		}
	}
	return ""
}

func DayOneB(fp *bufio.Reader) string {
	nums := readInputDayOne(fp)

	type stepTwoValue struct {
		a, b int
	}

	stepOne := make(map[int]struct{})
	stepTwo := make(map[int]stepTwoValue)

	for _, n := range nums {
		t := target - n
		// First check stepTwo to make sure we don't already have a match.
		if stv, ok := stepTwo[t]; ok {
			return strconv.Itoa(n * stv.a * stv.b)
		}

		// Iterate through stepOne dictionary and add pairs to stepTwo
		for k := range stepOne {
			stepTwo[n+k] = stepTwoValue{n, k}
		}
		stepOne[n] = struct{}{}
	}
	return ""
}
