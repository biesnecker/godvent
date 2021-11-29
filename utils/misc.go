package utils

import (
	"strconv"
)

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
