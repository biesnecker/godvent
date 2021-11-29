package twentysixteen

import (
	"bufio"
	"log"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/biesnecker/godvent/utils"
)

// dkqjcbctfqwu-uecxgpigt-jwpv-fgrctvogpv-128[cgptv]

func parseInputDayFour(handler func(hash, checksum string, sector int)) func(string) {
	return func(s string) {
		sectorEnd := strings.IndexByte(s, '[')
		checksumStart := sectorEnd + 1
		hashEnd := strings.LastIndexByte(s, '-')
		sectorStart := hashEnd + 1

		hash := s[:hashEnd]
		sectorString := s[sectorStart:sectorEnd]
		checksum := s[checksumStart : checksumStart+5]

		var sector int
		var err error

		if sector, err = strconv.Atoi(sectorString); err != nil {
			log.Fatalln(err)
		}

		handler(hash, checksum, sector)
	}
}

func DayFourA(fp *bufio.Reader) string {
	sum := 0
	utils.ReadStrings(fp, parseInputDayFour(func(hash, checksum string, sector int) {
		letterCounts := make(map[rune]int)

		for _, c := range hash {
			if unicode.IsLetter(c) {
				letterCounts[c] += 1
			}
		}

		type letCnt struct {
			c     rune
			count int
		}

		counts := make([]letCnt, 0, len(letterCounts))
		for k, v := range letterCounts {
			counts = append(counts, letCnt{c: k, count: v})
		}
		sort.SliceStable(counts, func(i, j int) bool {
			if counts[i].count == counts[j].count {
				return counts[i].c < counts[j].c
			} else {
				return counts[i].count > counts[j].count
			}
		})

		match := true
		for i, sc := range checksum {
			if counts[i].c != sc {
				match = false
				break
			}
		}

		if match {
			sum += sector
		}
	}))
	return strconv.Itoa(sum)
}

func DayFourB(fp *bufio.Reader) string {
	matchingSector := 0
	utils.ReadStrings(fp, parseInputDayFour(func(hash, checksum string, sector int) {
		bytes := []byte(hash)
		for i, b := range bytes {
			if unicode.IsLetter(rune(b)) {
				bytes[i] = byte(((int(b-'a') + sector) % 26) + 'a')
			}
		}
		if string(bytes) == "northpole-object-storage" {
			matchingSector = sector
		}
	}))

	return strconv.Itoa(matchingSector)
}
