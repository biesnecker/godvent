package twentyseventeen

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
	"github.com/biesnecker/godvent/utils/search"
)

type nodeD12 struct {
	id    int
	edges []*nodeD12
}

func (n *nodeD12) GetNext(interface{}) []search.Searchable {
	res := make([]search.Searchable, 0, len(n.edges))
	for i := range n.edges {
		res = append(res, n.edges[i])
	}
	return res
}

func (n *nodeD12) GetRepr(interface{}) interface{} {
	return n.id
}

func lookupOrInsertNode(nodes map[int]*nodeD12, id int) *nodeD12 {
	if node, found := nodes[id]; found {
		return node
	} else {
		newNode := nodeD12{id: id}
		nodes[id] = &newNode
		return &newNode
	}
}

func readInputDayTwelve(fp *bufio.Reader) map[int]*nodeD12 {
	nodes := make(map[int]*nodeD12)

	utils.ReadStrings(fp, func(s string) {
		parts := strings.Split(s, " <-> ")
		leftNode := lookupOrInsertNode(nodes, utils.ReadInt(parts[0]))
		rightParts := strings.Split(parts[1], ", ")
		for i := range rightParts {
			rightNode := lookupOrInsertNode(nodes, utils.ReadInt(rightParts[i]))
			leftNode.edges = append(leftNode.edges, rightNode)
			rightNode.edges = append(rightNode.edges, leftNode)
		}
	})

	return nodes
}

func DayTwelveA(fp *bufio.Reader) string {
	nodes := readInputDayTwelve(fp)

	nodeZero := nodes[0]

	bfs := search.NewBFSGenerator(nodeZero, nil)
	cnt := 0
	bfs.ForEach(func(elem search.Searchable, _ interface{}) bool {
		cnt++
		return true
	})

	return strconv.Itoa(cnt)
}

func DayTwelveB(fp *bufio.Reader) string {
	nodes := readInputDayTwelve(fp)

	groups := 0
	for len(nodes) > 0 {
		groups++

		var starting *nodeD12

		// Just get the first one that is returned, doesn't matter which.
		for _, v := range nodes {
			starting = v
			break
		}

		bfs := search.NewBFSGenerator(starting, nil)
		bfs.ForEach(func(elem search.Searchable, _ interface{}) bool {
			delete(nodes, elem.(*nodeD12).id)
			return true
		})
	}

	return strconv.Itoa(groups)
}
