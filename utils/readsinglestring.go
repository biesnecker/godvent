package utils

import (
	"bufio"
	"log"
	"strings"
)

func ReadSingleString(fp *bufio.Reader) string {
	if s, err := fp.ReadString('\n'); err != nil {
		log.Fatal(err)
	} else {
		return strings.TrimSpace(s)
	}
	return ""
}

func ReadDeliminatedInts(fp *bufio.Reader, delimiter string) []int {
	var res []int
	for _, s := range strings.Split(ReadSingleString(fp), delimiter) {
		res = append(res, ReadInt(s))
	}
	return res
}
