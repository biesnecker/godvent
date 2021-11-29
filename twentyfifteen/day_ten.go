package twentyfifteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type looksaynode struct {
	repetitions, value int
	next               *looksaynode
}

func stringToLookSayList(s string) *looksaynode {
	fakeHead := &looksaynode{}
	tail := fakeHead
	for _, c := range s {
		cval := int(c - '0')
		if cval == tail.value {
			tail.repetitions++
		} else {
			tail.next = &looksaynode{repetitions: 1, value: cval, next: nil}
			tail = tail.next
		}
	}
	return fakeHead.next
}

func stepLookSayList(old *looksaynode) (*looksaynode, int) {
	length := 0

	newHead := &looksaynode{}
	newTail := newHead

	for {
		if old == nil {
			break
		}
		if old.repetitions == newTail.value {
			newTail.repetitions++
		} else {
			length += newTail.repetitions
			newTail.next = &looksaynode{repetitions: 1, value: old.repetitions, next: nil}
			newTail = newTail.next
		}
		if old.value == newTail.value {
			newTail.repetitions++
		} else {
			length += newTail.repetitions
			newTail.next = &looksaynode{repetitions: 1, value: old.value}
			newTail = newTail.next
		}
		old = old.next
	}
	// Get the last one.
	length += newTail.repetitions

	return newHead.next, length
}

func DayTenA(fp *bufio.Reader) string {
	input := utils.ReadSingleString(fp)

	initial := stringToLookSayList(input)
	length := 0

	for i := 0; i < 40; i++ {
		initial, length = stepLookSayList(initial)
	}
	return strconv.Itoa(length)
}

func DayTenB(fp *bufio.Reader) string {
	input := utils.ReadSingleString(fp)

	initial := stringToLookSayList(input)
	length := 0

	for i := 0; i < 50; i++ {
		initial, length = stepLookSayList(initial)
	}
	return strconv.Itoa(length)
}
