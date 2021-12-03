package utils

import (
	"log"
	"strconv"
	"unicode"
)

func IsHexDigit(r rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, r)
}

func ReadInt(s string) int {
	if n, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return n
	}
}

func ReadByte(s string) byte {
	return byte(ReadInt(s))
}

func IntInBounds(x, min, max int) bool {
	return x >= min && x <= max
}

func ParseBinaryString(s string) int {
	var i int64
	var e error
	if i, e = strconv.ParseInt(s, 2, 64); e != nil {
		log.Fatalln(e)
	}
	return int(i)
}
