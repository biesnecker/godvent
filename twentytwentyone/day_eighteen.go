package twentytwentyone

import (
	"bufio"
	"fmt"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type snailfishNumber struct {
	left        *snailfishNumber
	right       *snailfishNumber
	parent      *snailfishNumber
	simpleValue int
	isSimple    bool
}

func sfnNewLiteral(parent *snailfishNumber, value int) *snailfishNumber {
	return &snailfishNumber{parent: parent, simpleValue: value, isSimple: true}
}

func sfnNewPair(parent, left, right *snailfishNumber) *snailfishNumber {
	p := &snailfishNumber{parent: parent, left: left, right: right}
	left.parent = p
	right.parent = p
	return p
}

func (sfn *snailfishNumber) String() string {
	if sfn.isSimple {
		return strconv.Itoa(sfn.simpleValue)
	} else {
		return fmt.Sprintf("[%s,%s]", sfn.left.String(), sfn.right.String())
	}
}

func (sfn *snailfishNumber) clone() *snailfishNumber {
	if sfn.isSimple {
		return sfnNewLiteral(nil, sfn.simpleValue)
	} else {
		left := sfn.left.clone()
		right := sfn.right.clone()
		return sfnNewPair(nil, left, right)
	}
}

func (sfn *snailfishNumber) magnitude() int {
	if sfn.isSimple {
		return sfn.simpleValue
	} else {
		return (3 * sfn.left.magnitude()) + (2 * sfn.right.magnitude())
	}
}

func (sfn *snailfishNumber) isLeftChild(other *snailfishNumber) bool {
	if sfn.isSimple {
		return false
	} else {
		return sfn.left == other
	}
}

func (sfn *snailfishNumber) isRightChild(other *snailfishNumber) bool {
	if sfn.isSimple {
		return false
	} else {
		return sfn.right == other
	}
}

func (sfn *snailfishNumber) findLeftNeighbor() *snailfishNumber {
	child := sfn
	current := sfn.parent
	for {
		if current == nil {
			// Got to the root.
			return nil
		} else if current.isLeftChild(child) {
			// The child is the left child of the parent, so the left neighbor
			// must be a child of the grandparent.
			child = current
			current = current.parent
		} else {
			// The child is the right child of the parent, so find the right-
			// most child of the left child.
			leftChild := current.left
			for !leftChild.isSimple {
				leftChild = leftChild.right
			}
			return leftChild
		}
	}
}

func (sfn *snailfishNumber) findRightNeighbor() *snailfishNumber {
	child := sfn
	current := sfn.parent
	for {
		if current == nil {
			// Got to the root.
			return nil
		} else if current.isRightChild(child) {
			// The child is the right child of the parent, so the right neighbor
			// must be a child of the grandparent.
			child = current
			current = current.parent
		} else {
			// The child is the left child of the parent, so find the left-
			// most child of the right child.
			rightChild := current.right
			for !rightChild.isSimple {
				rightChild = rightChild.left
			}
			return rightChild
		}
	}
}

func addSnailfishNumbers(left, right *snailfishNumber) *snailfishNumber {
	return sfnNewPair(nil, left, right)
}

func explodeSnailfishNumber(sfn *snailfishNumber, depth int) bool {
	if sfn.isSimple {
		// Simple values can't explode
		return false
	} else if depth < 4 {
		if explodeSnailfishNumber(sfn.left, depth+1) {
			return true
		} else if explodeSnailfishNumber(sfn.right, depth+1) {
			return true
		} else {
			return false
		}
	} else {
		// This node needs to explode.
		if !sfn.left.isSimple || !sfn.right.isSimple {
			panic("Trying to explode number with non-literal children")
		}
		leftVal := sfn.left.simpleValue
		leftNeighbor := sfn.findLeftNeighbor()
		if leftNeighbor != nil {
			leftNeighbor.simpleValue += leftVal
		}
		rightVal := sfn.right.simpleValue
		rightNeighbor := sfn.findRightNeighbor()
		if rightNeighbor != nil {
			rightNeighbor.simpleValue += rightVal
		}
		sfn.left = nil
		sfn.right = nil
		sfn.isSimple = true
		sfn.simpleValue = 0
		return true
	}
}

func splitSnailfishNumber(sfn *snailfishNumber) bool {
	if sfn.isSimple {
		if sfn.simpleValue > 9 {
			sfn.left = &snailfishNumber{
				parent: sfn, isSimple: true, simpleValue: sfn.simpleValue / 2}
			sfn.right = &snailfishNumber{
				parent: sfn, isSimple: true, simpleValue: (sfn.simpleValue + 1) / 2}
			sfn.isSimple = false
			sfn.simpleValue = 0
			return true
		} else {
			return false
		}
	} else {
		if splitSnailfishNumber(sfn.left) {
			return true
		} else if splitSnailfishNumber(sfn.right) {
			return true
		} else {
			return false
		}
	}
}

func reduceSnailfishNumber(sfn *snailfishNumber) *snailfishNumber {
	didSomething := true
	for didSomething {
		didSomething = false
		if explodeSnailfishNumber(sfn, 0) {
			didSomething = true
		} else if splitSnailfishNumber(sfn) {
			didSomething = true
		}
	}
	return sfn
}

func parseSnailfishNumber(s string) *snailfishNumber {
	var root, current *snailfishNumber
	var sideStack []bool
	for i := range s {
		c := s[i]
		switch c {
		case '[':
			if current == nil {
				// create the root
				root = &snailfishNumber{parent: nil}
				current = root
				sideStack = append(sideStack, true)
			} else {
				onLeft := sideStack[len(sideStack)-1]
				current.isSimple = false
				if onLeft {
					current.left = &snailfishNumber{parent: current}
					current = current.left
					sideStack[len(sideStack)-1] = false
					sideStack = append(sideStack, true)
				} else {
					current.right = &snailfishNumber{parent: current}
					current = current.right
					sideStack = append(sideStack, true)
				}
			}
		case ']':
			if !current.isSimple && (current.left == nil || current.right == nil) {
				panic("unbalanced")
			}
			sideStack = sideStack[:len(sideStack)-1]
			current = current.parent
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			onLeft := sideStack[len(sideStack)-1]
			current.isSimple = false
			if onLeft {
				current.left = &snailfishNumber{
					parent:      current,
					isSimple:    true,
					simpleValue: int(c - '0')}
				sideStack[len(sideStack)-1] = false
			} else {
				current.right = &snailfishNumber{
					parent:      current,
					isSimple:    true,
					simpleValue: int(c - '0')}
			}
		}
	}
	return root
}

func readInputD18(fp *bufio.Reader) []*snailfishNumber {
	res := make([]*snailfishNumber, 0, 128)
	utils.ReadStrings(fp, func(s string) {
		res = append(res, parseSnailfishNumber(s))
	})
	return res
}

func DayEighteenA(fp *bufio.Reader) string {
	nums := readInputD18(fp)

	sum := nums[0]

	for _, n := range nums[1:] {
		sum = addSnailfishNumbers(sum, n)
		sum = reduceSnailfishNumber(sum)
	}

	return strconv.Itoa(sum.magnitude())
}

func DayEighteenB(fp *bufio.Reader) string {
	nums := readInputD18(fp)

	max := math.MinInt

	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}
			icopy := nums[i].clone()
			jcopy := nums[j].clone()
			sum := reduceSnailfishNumber(addSnailfishNumbers(icopy, jcopy))
			mag := sum.magnitude()
			if mag > max {
				max = mag
			}
		}
	}
	return strconv.Itoa(max)
}
