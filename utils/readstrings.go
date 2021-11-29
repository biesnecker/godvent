package utils

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func ReadStrings(
	fp *bufio.Reader,
	handler func(string),
) {
	ReadStringsUntilBreak(fp, true, func(s string) bool {
		handler(s)
		return true
	})
}

func ReadStringsAsSlice(fp *bufio.Reader) []string {
	var res []string
	ReadStrings(fp, func(s string) {
		res = append(res, s)
	})
	return res
}

func ReadStringsWithIndex(
	fp *bufio.Reader,
	handler func(int, string),
) {
	idx := 0
	ReadStringsUntilBreak(fp, true, func(s string) bool {
		handler(idx, s)
		idx++
		return true
	})
}

func ReadStringsWithIndexNoTrim(
	fp *bufio.Reader,
	handler func(int, string),
) {
	idx := 0
	ReadStringsUntilBreak(fp, false, func(s string) bool {
		handler(idx, s)
		idx++
		return true
	})
}

func ReadStringsUntilBreak(fp *bufio.Reader, trim bool, handler func(string) bool) {
	for {
		if s, err := fp.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if trim {
				s = strings.TrimSpace(s)
			}
			if !handler(s) {
				break
			}
		}
	}
}

func ParseDelimitedIntegerString(s, delimiter string) ([]int, error) {
	var parsingError error
	parts := strings.Split(s, delimiter)
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		if n, err := strconv.Atoi(p); err != nil {
			parsingError = err
			break
		} else {
			nums = append(nums, n)
		}
	}
	if parsingError != nil {
		return nil, parsingError
	}
	return nums, nil
}

func ParseDelimitedByteString(s, delimiter string) ([]byte, error) {
	var parsingError error
	parts := strings.Split(s, delimiter)
	nums := make([]byte, 0, len(parts))
	for _, p := range parts {
		if n, err := strconv.Atoi(p); err != nil {
			parsingError = err
			break
		} else {
			nums = append(nums, byte(n))
		}
	}
	if parsingError != nil {
		return nil, parsingError
	}
	return nums, nil
}

func ReadDelimitedIntegerStrings(
	fp *bufio.Reader,
	delimiter string,
	handler func([]int),
) error {
	var parsingError error = nil
	ReadStringsUntilBreak(fp, true, func(s string) bool {
		if nums, err := ParseDelimitedIntegerString(s, delimiter); err != nil {
			parsingError = err
			return false
		} else {
			handler(nums)
			return true
		}
	})
	if parsingError != nil {
		return parsingError
	} else {
		return nil
	}
}

func ReadOneIntegerPerLine(fp *bufio.Reader) ([]int, error) {
	var parsingError error = nil
	var nums []int
	ReadStringsUntilBreak(fp, true, func(s string) bool {
		if n, err := strconv.Atoi(s); err != nil {
			parsingError = err
			return false
		} else {
			nums = append(nums, n)
			return true
		}
	})
	if parsingError != nil {
		return nil, parsingError
	} else {
		return nums, nil
	}
}
