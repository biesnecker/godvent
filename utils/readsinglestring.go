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
