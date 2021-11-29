package twentytwenty

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

type inputD2 struct {
	pw       string
	c        rune
	min, max int
}

func readInput(fp *bufio.Reader) []inputD2 {
	var res []inputD2
	utils.ReadStrings(fp, func(s string) {
		s = strings.ReplaceAll(s, ":", "")
		s = strings.ReplaceAll(s, "-", " ")
		var v inputD2
		var cstr string
		_, err := fmt.Sscanf(s, "%d %d %s %s", &v.min, &v.max, &cstr, &v.pw)
		if err != nil {
			log.Fatalln(err)
		}
		if len(cstr) != 1 {
			log.Fatalln("Bad length: ", cstr, len(cstr))
		}
		v.c = rune(cstr[0])
		res = append(res, v)
	})
	return res
}

func DayTwoA(fp *bufio.Reader) string {
	vals := readInput(fp)
	c := 0
	for _, v := range vals {
		count := 0
		for _, b := range v.pw {
			if b == v.c {
				count++
			}
		}
		if count >= v.min && count <= v.max {
			c++
		}
	}
	return strconv.Itoa(c)
}

func DayTwoB(fp *bufio.Reader) string {
	vals := readInput(fp)
	c := 0
	for _, v := range vals {
		letterA := v.pw[v.min-1] == byte(v.c)
		letterB := v.pw[v.max-1] == byte(v.c)
		if (letterA || letterB) && !(letterA && letterB) {
			c++
		}
	}
	return strconv.Itoa(c)
}
