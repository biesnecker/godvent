package twentysixteen

import (
	"bufio"
	"fmt"

	"github.com/biesnecker/godvent/utils"
)

func DayFiveA(fp *bufio.Reader) string {
	doorId := utils.ReadSingleString(fp)
	password := make([]byte, 8)
	found := 0
	idx := 0
	for found < 8 {
		s := fmt.Sprintf("%s%d", doorId, idx)
		hash := utils.GetMD5StringOfString(s)
		for i, c := range hash {
			if i < 5 && c != '0' {
				break
			} else if i == 5 {
				password[found] = byte(c)
				found++
				break
			}
		}
		idx++
	}
	return string(password)
}

func DayFiveB(fp *bufio.Reader) string {
	doorId := utils.ReadSingleString(fp)
	password := make([]byte, 8)
	found := 0
	idx := 0
	for found < 8 {
		var pos int
		s := fmt.Sprintf("%s%d", doorId, idx)
		hash := utils.GetMD5StringOfString(s)
		for i, c := range hash {
			if i < 5 && c != '0' {
				break
			} else if i == 5 {
				pos = int(c - '0')
				if pos < 0 || pos > 7 {
					break
				}
				if password[pos] != 0 {
					break
				}
			} else if i == 6 {
				password[pos] = byte(c)
				found++
				break
			}
		}
		idx++
	}
	return string(password)
}
