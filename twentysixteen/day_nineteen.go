package twentysixteen

import (
	"bufio"
	"container/list"
	"container/ring"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayNineteen(fp *bufio.Reader) int {
	s := utils.ReadSingleString(fp)
	if n, err := strconv.Atoi(s); err != nil {
		log.Fatalln(err)
	} else {
		return n
	}
	return 0
}

func findSolutionDayNineteenA(fp *bufio.Reader) int {
	elves := readInputDayNineteen(fp)

	r := ring.New(elves)

	rfiller := r
	for i := 0; i < elves; i++ {
		rfiller.Value = i
		rfiller = rfiller.Next()
	}

	for {
		if r == r.Next() {
			// Answer is one-indexed
			return r.Value.(int) + 1
		}
		r.Unlink(1)
		r = r.Next()
	}
}

func findSolutionNineteenB(fp *bufio.Reader) int {
	elves := readInputDayNineteen(fp)

	left := list.New()
	right := list.New()

	for i := 0; i < elves; i++ {
		if i < elves/2 {
			left.PushBack(i)
		} else {
			right.PushFront(i)
		}
	}

	for right.Len() > 0 {
		if left.Len() > right.Len() {
			left.Remove(left.Back())
		} else {
			right.Remove(right.Back())
		}

		// Rotate
		leftFront := left.Front()
		leftValue := leftFront.Value
		left.Remove(leftFront)
		right.PushFront(leftValue)

		rightBack := right.Back()
		rightValue := rightBack.Value
		right.Remove(rightBack)
		left.PushBack(rightValue)
	}
	return left.Front().Value.(int) + 1
}

func DayNineteenA(fp *bufio.Reader) string {
	solution := findSolutionDayNineteenA(fp)
	return strconv.Itoa(solution)
}

func DayNineteenB(fp *bufio.Reader) string {
	solution := findSolutionNineteenB(fp)
	return strconv.Itoa(solution)
}
