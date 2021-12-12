package twentytwentyone

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/biesnecker/godvent/utils"
)

type node struct {
	label string
	next  []string
	big   bool
}

func copySeen(m map[string]bool) map[string]bool {
	newm := make(map[string]bool)
	for k, v := range m {
		newm[k] = v
	}
	return newm
}

func copyPath(old []string) []string {
	new := make([]string, len(old))
	copy(new, old)
	return new
}

func getNode(m map[string]*node, label string) *node {
	if n, ok := m[label]; ok {
		return n
	} else {
		big := unicode.IsUpper(rune(label[0]))
		n := node{label: label, big: big}
		m[label] = &n
		return &n
	}
}

func readInputDayTwelve(fp *bufio.Reader) map[string]*node {
	res := make(map[string]*node)

	utils.ReadStrings(fp, func(s string) {
		s = strings.ReplaceAll(s, "-", " ")
		var left, right string
		fmt.Sscanf(s, "%s %s", &left, &right)
		leftN := getNode(res, left)
		rightN := getNode(res, right)

		leftN.next = append(leftN.next, right)
		rightN.next = append(rightN.next, left)
	})

	return res
}

func getPath(
	allNodes map[string]*node,
	currentNode *node,
	path []string,
	seen map[string]bool,
	allPaths map[string]bool) {

	for _, next := range currentNode.next {
		nn := allNodes[next]
		if !nn.big && seen[nn.label] {
			continue
		}

		if currentNode.label == "end" {
			pj := strings.Join(path, "-")
			allPaths[pj] = true
			continue
		}

		newseen := copySeen(seen)
		if !nn.big {
			newseen[nn.label] = true
		}
		newpath := copyPath(path)
		newpath = append(newpath, nn.label)

		getPath(allNodes, nn, newpath, newseen, allPaths)

	}

}

func DayTwelveA(fp *bufio.Reader) string {
	input := readInputDayTwelve(fp)

	allPaths := make(map[string]bool)

	getPath(input, input["start"], []string{"start"}, map[string]bool{"start": true}, allPaths)
	return strconv.Itoa(len(allPaths))
}

func getPath2(
	allNodes map[string]*node,
	currentNode *node,
	path []string,
	seen map[string]bool,
	doneTwice bool,
	allPaths map[string]bool) {

	for _, next := range currentNode.next {
		nn := allNodes[next]
		newDoneTwice := doneTwice
		if !nn.big && seen[nn.label] {
			if doneTwice || nn.label == "start" {
				continue
			} else {
				newDoneTwice = true
			}
		}

		if currentNode.label == "end" {
			pj := strings.Join(path, "-")
			allPaths[pj] = true
			continue
		}

		newseen := copySeen(seen)
		if !nn.big {
			newseen[nn.label] = true
		}
		newpath := copyPath(path)
		newpath = append(newpath, nn.label)

		getPath2(allNodes, nn, newpath, newseen, newDoneTwice, allPaths)

	}

}

func DayTwelveB(fp *bufio.Reader) string {
	input := readInputDayTwelve(fp)

	allPaths := make(map[string]bool)

	getPath2(input, input["start"], []string{"start"}, map[string]bool{"start": true}, false, allPaths)
	return strconv.Itoa(len(allPaths))
}
