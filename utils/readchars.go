package utils

import (
	"bufio"
	"io"
	"log"
)

func ReadChars(
	fp *bufio.Reader,
	handler func(c rune, sz int) bool,
) {
	for {
		if c, sz, err := fp.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if !handler(c, sz) {
				break
			}
		}
	}
}
