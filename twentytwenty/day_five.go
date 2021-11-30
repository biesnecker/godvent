package twentytwenty

import (
	"bufio"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func seatFromString(s string) (int, int) {
	rowTotal := utils.IntPow(2, len(s)-3)
	row := 0
	for _, v := range s[:len(s)-3] {
		switch v {
		case 'F':
			rowTotal /= 2
		case 'B':
			nt := rowTotal / 2
			row += nt
			rowTotal = nt
		default:
			log.Fatalln("Unknown value: ", v)
		}
	}
	seatTotal := 8
	seat := 0
	for _, v := range s[len(s)-3:] {
		switch v {
		case 'L':
			seatTotal /= 2
		case 'R':
			nt := seatTotal / 2
			seat += nt
			seatTotal = nt
		default:
			log.Fatal("Unknown value: ", v)
		}
	}
	return row, seat
}

func DayFiveA(fp *bufio.Reader) string {
	m := 0
	utils.ReadStrings(fp, func(s string) {
		row, seat := seatFromString(s)
		id := row*8 + seat
		if id > m {
			m = id
		}
	})
	return strconv.Itoa(m)
}

func DayFiveB(fp *bufio.Reader) string {
	seats := make([]bool, 128*8)
	utils.ReadStrings(fp, func(s string) {
		row, seat := seatFromString(s)
		id := row*8 + seat
		seats[id] = true
	})

	for i := 1; i < len(seats)-1; i++ {
		if !seats[i] && seats[i-1] && seats[i+1] {
			return strconv.Itoa(i)
		}
	}
	return ""
}
