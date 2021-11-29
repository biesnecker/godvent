package twentyseventeen

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

type balancedProg struct {
	label           string
	weight          int
	supportedWeight int
	supportedBy     *balancedProg
	supports        []*balancedProg
}

func (bp *balancedProg) setWeights() int {
	for _, child := range bp.supports {
		bp.supportedWeight += child.setWeights()
	}
	return bp.weight + bp.supportedWeight
}

func (bp *balancedProg) getTotalWeight() int {
	return bp.weight + bp.supportedWeight
}

func (bp *balancedProg) getUnbalancedChild() *balancedProg {
	if len(bp.supports) == 0 {
		return nil
	}
	for i := range bp.supports {
		foundMatch := false
		for j := range bp.supports {
			if i == j {
				continue
			}
			if bp.supports[i].getTotalWeight() == bp.supports[j].getTotalWeight() {
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			return bp.supports[i]
		}
	}
	return nil
}

func newBalancedProg(label string, weight int) *balancedProg {
	return &balancedProg{label: label, weight: weight}
}

func findTreeRoot(progs map[string]*balancedProg) *balancedProg {
	for _, p := range progs {
		if p.supportedBy == nil {
			return p
		}
	}
	return nil
}

func lookupOrInsertProg(
	progs map[string]*balancedProg,
	label string,
	weight int,
) *balancedProg {
	if found, ok := progs[label]; !ok {
		newprog := newBalancedProg(label, weight)
		progs[label] = newprog
		return newprog
	} else {
		if found.weight == 0 && weight != 0 {
			found.weight = weight
		}
		return found
	}
}

func readInputDaySeven(fp *bufio.Reader) map[string]*balancedProg {
	res := make(map[string]*balancedProg)
	utils.ReadStrings(fp, func(s string) {
		var label string
		var weight int
		if strings.Contains(s, "->") {
			parts := strings.Split(s, " -> ")
			if _, err := fmt.Sscanf(parts[0], "%s (%d)", &label, &weight); err != nil {
				log.Fatalln(err)
			}
			supporting := lookupOrInsertProg(res, label, weight)

			for _, supportedLabel := range strings.Split(parts[1], ", ") {
				supported := lookupOrInsertProg(res, supportedLabel, 0)
				supported.supportedBy = supporting
				supporting.supports = append(supporting.supports, supported)
			}
		} else {
			if _, err := fmt.Sscanf(s, "%s (%d)", &label, &weight); err != nil {
				log.Fatalln(err)
			}
			lookupOrInsertProg(res, label, weight)
		}
	})
	return res
}

func DaySevenA(fp *bufio.Reader) string {
	return findTreeRoot(readInputDaySeven(fp)).label
}

func DaySevenB(fp *bufio.Reader) string {
	tree := readInputDaySeven(fp)
	root := findTreeRoot(tree)
	root.setWeights()

	for {
		unbalanced := root.getUnbalancedChild()
		if unbalanced == nil {
			// Then root is the problem node.
			break
		} else {
			root = unbalanced
		}
	}

	unbalancedParent := root.supportedBy
	neededWeight := (unbalancedParent.supportedWeight - root.getTotalWeight()) /
		(len(unbalancedParent.supports) - 1)

	return strconv.Itoa(root.weight + (neededWeight - root.getTotalWeight()))
}
