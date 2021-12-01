package twentytwenty

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/types/queue"
	"github.com/biesnecker/godvent/types/set"
	"github.com/biesnecker/godvent/utils"
)

type bagD7 struct {
	name        string
	contains    map[string]int
	containedBy []string
}

func findBagByName(m map[string]*bagD7, name string) *bagD7 {
	if bag, ok := m[name]; ok {
		return bag
	} else {
		b := &bagD7{
			name:        name,
			contains:    make(map[string]int),
			containedBy: make([]string, 0),
		}
		m[name] = b
		return b
	}
}

func readInputDaySeven(fp *bufio.Reader) map[string]*bagD7 {
	res := make(map[string]*bagD7)

	utils.ReadStrings(fp, func(s string) {
		s = s[:len(s)-1]
		leftRight := strings.SplitAfterN(s, "contain ", 2)
		left := strings.TrimSpace(leftRight[0][:strings.LastIndex(leftRight[0], " bag")])

		leftBag := findBagByName(res, left)
		rights := strings.Split(leftRight[1], ", ")
		for _, right := range rights {
			right = right[:strings.LastIndex(right, " bag")]
			var count int
			var worda, wordb string
			fmt.Sscanf(right, "%d %s %s", &count, &worda, &wordb)
			right = fmt.Sprintf("%s %s", worda, wordb)
			rightBag := findBagByName(res, right)

			leftBag.contains[right] = count
			rightBag.containedBy = append(rightBag.containedBy, left)
		}
	})

	return res
}

func DaySevenA(fp *bufio.Reader) string {
	input := readInputDaySeven(fp)
	seen := set.New()
	q := queue.New()
	q.Push(findBagByName(input, "shiny gold"))

	for !q.Empty() {
		current := q.Pop().(*bagD7)
		for _, b := range current.containedBy {
			if !seen.Contains(b) {
				seen.Insert(b)
				q.Push(findBagByName(input, b))
			}
		}
	}
	return strconv.Itoa(seen.Len())
}

func countBags(m map[string]*bagD7, bag *bagD7, multiplier int) int {
	count := multiplier
	for k, v := range bag.contains {
		bag := findBagByName(m, k)
		count += countBags(m, bag, v*multiplier)
	}
	return count
}

func DaySevenB(fp *bufio.Reader) string {
	input := readInputDaySeven(fp)

	start := findBagByName(input, "shiny gold")
	total := countBags(input, start, 1)

	// Subtract one because we don't want to count the starting bag.
	return strconv.Itoa(total - 1)
}
