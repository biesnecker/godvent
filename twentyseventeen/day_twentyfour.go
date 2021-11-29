package twentyseventeen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
	"github.com/biesnecker/godvent/utils/search"
)

type component struct {
	head, tail int
}

type stateD24 struct {
	headIndex, tailIndex map[int][]*component
}

type stepD24 struct {
	components []int
}

func (s *stepD24) Copy() *stepD24 {
	newComponents := make([]int, len(s.components))
	copy(newComponents, s.components)
	return &stepD24{components: newComponents}
}

func (s *stepD24) GetNext(s interface{}) []search.Searchable {
	state := s.(*stateD24)
	var next []search.Searchable

	return next
}

func (s *stepD24) GetRepr(interface{}) interface{} {
	parts := make([]string, 0, len(s.components))
	for i := range s.components {
		parts = append(parts, strconv.Itoa(s.components[i]))
	}
	return strings.Join(parts, "_")
}

func readInputDay24(fp *bufio.Reader) *stateD24 {
	headIndex := make(map[int][]*component)
	tailIndex := make(map[int][]*component)

	utils.ReadStrings(fp, func(s string) {
		var head, tail int
		fmt.Sscanf(s, "%d/%d", &head, &tail)
		c := &component{head: head, tail: tail}
		headIndex[head] = append(headIndex[head], c)
		tailIndex[tail] = append(tailIndex[tail], c)
	})

	return &stateD24{headIndex: headIndex, tailIndex: tailIndex}
}

func DayTwentyFourA(fp *bufio.Reader) string {
	return ""
}

func DayTwentyFourB(fp *bufio.Reader) string {
	return ""
}
