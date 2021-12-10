package twentynineteen

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/types/queue"
	"github.com/biesnecker/godvent/types/set"
	"github.com/biesnecker/godvent/utils"
)

type nodeD6 struct {
	label     string
	orbits    []*nodeD6
	orbitedBy []*nodeD6
}

func getOrCreateNode(m map[string]*nodeD6, label string) *nodeD6 {
	if n, ok := m[label]; ok {
		return n
	} else {
		newnode := nodeD6{label: label}
		m[label] = &newnode
		return &newnode
	}
}

func readInputDaySix(fp *bufio.Reader) map[string]*nodeD6 {
	res := make(map[string]*nodeD6)
	utils.ReadStrings(fp, func(s string) {
		s = strings.ReplaceAll(s, ")", " ")
		var left, right string
		fmt.Sscanf(s, "%s %s", &left, &right)
		leftN := getOrCreateNode(res, left)
		rightN := getOrCreateNode(res, right)
		rightN.orbits = append(rightN.orbits, leftN)
		leftN.orbitedBy = append(leftN.orbitedBy, rightN)
	})
	return res
}

func countOrbits(cache map[string]int, n *nodeD6) int {
	if c, ok := cache[n.label]; ok {
		return c
	} else {
		total := 0
		for _, child := range n.orbits {
			total += countOrbits(cache, child) + 1
		}
		cache[n.label] = total
		return total
	}
}

func DaySixA(fp *bufio.Reader) string {
	input := readInputDaySix(fp)
	cache := make(map[string]int)
	for _, in := range input {
		countOrbits(cache, in)
	}

	total := 0
	for _, c := range cache {
		total += c
	}
	return strconv.Itoa(total)
}

func DaySixB(fp *bufio.Reader) string {
	input := readInputDaySix(fp)

	youN := getOrCreateNode(input, "YOU")
	sanN := getOrCreateNode(input, "SAN")

	if len(youN.orbits) != 1 || len(sanN.orbits) != 1 {
		log.Fatalln("YOU or SAN orbits more than one things")
	}
	start := youN.orbits[0]
	end := sanN.orbits[0]

	type step struct {
		n     *nodeD6
		steps int
	}

	q := queue.New()
	q.Push(step{n: start, steps: 0})

	seen := set.New()
	seen.Insert(start.label)

	minsteps := math.MaxInt

	for !q.Empty() {
		current := q.Pop().(step)
		if current.n == end {
			if current.steps < minsteps {
				minsteps = current.steps
			}
			continue
		}
		for _, v := range current.n.orbitedBy {
			if !seen.Contains(v.label) {
				q.Push(step{n: v, steps: current.steps + 1})
				seen.Insert(v.label)
			}
		}
		for _, v := range current.n.orbits {
			if !seen.Contains(v.label) {
				q.Push(step{n: v, steps: current.steps + 1})
				seen.Insert(v.label)
			}
		}
	}
	return strconv.Itoa(minsteps)
}
