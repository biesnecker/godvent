package twentyfifteen

import (
	"bufio"
	"hash/maphash"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayNineteen(fp *bufio.Reader) (map[string]*[]string, string) {
	replacements := make(map[string]*[]string)
	var molecule string
	split := false

	utils.ReadStrings(fp, func(s string) {
		if split {
			molecule = s
		} else if s == "" {
			split = true
		} else {
			parts := strings.Fields(s)
			if existing, ok := replacements[parts[0]]; ok {
				*existing = append(*existing, parts[2])
			} else {
				newslice := make([]string, 0, 2)
				newslice = append(newslice, parts[2])
				replacements[parts[0]] = &newslice
			}

		}
	})
	return replacements, molecule
}

type replacementDayNineteen struct {
	orig, rep string
}

func invertReplacementsDayNineteen(r map[string]*[]string) []replacementDayNineteen {
	res := make([]replacementDayNineteen, 0, 50)
	for k, v := range r {
		for _, rep := range *v {
			res = append(res, replacementDayNineteen{orig: rep, rep: k})
		}
	}
	return res
}

func DayNineteenA(fp *bufio.Reader) string {
	h := maphash.Hash{}
	h.SetSeed(maphash.MakeSeed())
	replacements, molecule := readInputDayNineteen(fp)
	seen := make(map[uint64]struct{})

	for i := 0; i < len(molecule); i++ {
		var prev, r, post string
		prev = molecule[:i]
		if i+1 < len(molecule) && unicode.IsLower(rune(molecule[i+1])) {
			r = molecule[i : i+2]
			post = molecule[i+2:]
		} else {
			r = molecule[i : i+1]
			post = molecule[i+1:]
		}
		if reps, ok := replacements[r]; ok {
			for _, rep := range *reps {
				h.WriteString(prev)
				h.WriteString(rep)
				h.WriteString(post)
				seen[h.Sum64()] = struct{}{}
				h.Reset()
			}
		}
	}

	return strconv.Itoa(len(seen))
}

func DayNineteenB(fp *bufio.Reader) string {
	h := maphash.Hash{}
	h.SetSeed(maphash.MakeSeed())
	replacements, molecule := readInputDayNineteen(fp)
	ir := invertReplacementsDayNineteen(replacements)

	rand.Seed(time.Now().UnixNano())
	steps := 0
	for molecule != "e" {
		for _, r := range ir {
			slen := len(r.orig)
			idx := strings.Index(molecule, r.orig)
			if idx == -1 {
				continue
			}
			steps++
			parts := []string{
				molecule[:idx],
				r.rep,
				molecule[idx+slen:]}
			molecule = strings.Join(parts, "")
		}
	}
	return strconv.Itoa(steps)
}
