package twentytwenty

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayFour(fp *bufio.Reader) []map[string]string {
	var res []map[string]string
	var m *map[string]string
	utils.ReadStrings(fp, func(s string) {
		if len(s) == 0 && m != nil {
			res = append(res, *m)
			m = nil
		} else {
			if m == nil {
				newMap := make(map[string]string)
				m = &newMap
			}
			for _, kv := range strings.Split(s, " ") {
				kvs := strings.Split(kv, ":")
				if len(kvs) != 2 {
					log.Fatalln("Bad length: ", len(kvs), kv)
				}
				key := kvs[0]
				val := kvs[1]
				(*m)[key] = val
			}
		}
	})
	if m != nil {
		res = append(res, *m)
	}
	return res
}

func getRequired() map[string]bool {
	return map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
		"cid": false,
	}
}

func checkFourDigit(s string, min, max int) bool {
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return v >= min && v <= max
}

func isValidDayFour(id map[string]string, validate bool) bool {
	req := getRequired()
	for k := range id {
		req[k] = true
	}
	for k, v := range req {
		if k == "cid" {
			continue
		}
		if !v {
			return false
		}
		if !validate {
			continue
		}
		switch k {
		case "byr":
			if !checkFourDigit(id[k], 1920, 2002) {
				return false
			}
		case "iyr":
			if !checkFourDigit(id[k], 2010, 2020) {
				return false
			}
		case "eyr":
			if !checkFourDigit(id[k], 2020, 2030) {
				return false
			}
		case "hgt":
			v := id[k]
			var i int
			var s string
			if _, err := fmt.Sscanf(v, "%d%s", &i, &s); err != nil {
				return false
			}
			switch s {
			case "in":
				if i < 59 || i > 76 {
					return false
				}
			case "cm":
				if i < 150 || i > 193 {
					return false
				}
			default:
				return false
			}
		case "hcl":
			v := id[k]
			if len(v) != 7 {
				return false
			}
			if v[0] != '#' {
				return false
			}
			for i := 1; i < 7; i++ {
				if !utils.IsHexDigit(rune(v[i])) {
					return false
				}
			}
		case "ecl":
			v := id[k]
			if !(v == "amb" || v == "blu" || v == "brn" || v == "gry" ||
				v == "grn" || v == "hzl" || v == "oth") {
				return false
			}
		case "pid":
			v := id[k]
			if len(v) != 9 {
				return false
			}
			for i := 0; i < 9; i++ {
				if !unicode.IsDigit(rune(v[i])) {
					return false
				}
			}
		default:
			return false
		}
	}
	return true
}

func DayFourA(fp *bufio.Reader) string {
	count := 0
	for _, id := range readInputDayFour(fp) {
		if isValidDayFour(id, false) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func DayFourB(fp *bufio.Reader) string {
	count := 0
	for _, id := range readInputDayFour(fp) {
		if isValidDayFour(id, true) {
			count++
		}
	}
	return strconv.Itoa(count)
}
