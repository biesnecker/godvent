package twentyfifteen

import (
	"bufio"
	"log"
	"strconv"
	"unicode"

	"github.com/biesnecker/godvent/utils"
)

func addDigitToTotal(total int, digit rune) int {
	total *= 10
	total += int(digit - '0')
	return total
}

func calculateFinalNumberAndClear(total *int, isNegative *bool) int {
	t := *total
	if *isNegative {
		t *= -1
	}
	*total = 0
	*isNegative = false
	return t
}

func DayTwelveA(fp *bufio.Reader) string {
	total := 0
	currentNumber := 0
	isNegative := false
	utils.ReadChars(fp, func(c rune, _ int) bool {
		if c == '-' {
			isNegative = true
		} else if unicode.IsDigit(c) {
			currentNumber = addDigitToTotal(currentNumber, c)
		} else if currentNumber > 0 {
			total += calculateFinalNumberAndClear(&currentNumber, &isNegative)
		}
		return true
	})
	return strconv.Itoa(total)
}

func DayTwelveB(fp *bufio.Reader) string {
	currentNumber := 0
	isNegative := false
	lastRedChar := rune(0)

	type node struct {
		tp    rune
		total int
		isRed bool
		prev  *node
	}

	stackBottom := &node{tp: 0, total: 0, isRed: false, prev: nil}
	stackTop := stackBottom

	utils.ReadChars(fp, func(c rune, _ int) bool {
		switch c {
		case '-':
			isNegative = true
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			currentNumber *= 10
			currentNumber += int(c - '0')
		default:
			stackTop.total += calculateFinalNumberAndClear(&currentNumber, &isNegative)
			switch c {
			case '{', '[':
				nst := &node{tp: c, total: 0, isRed: false, prev: stackTop}
				stackTop = nst
			case '}', ']':
				if (c == '}' && stackTop.tp != '{') || (c == ']' && stackTop.tp != '[') {
					log.Fatalln("Mismatched bracket.")
				}
				if !stackTop.isRed || stackTop.tp != '{' {
					stackTop.prev.total += stackTop.total
				}
				stackTop = stackTop.prev
			case ':':
				if lastRedChar == 0 {
					lastRedChar = ':'
				} else {
					lastRedChar = 0
				}
			case '"':
				if lastRedChar == ':' {
					lastRedChar = '"'
				} else if lastRedChar == 'd' {
					stackTop.isRed = true
					lastRedChar = 0
				} else {
					lastRedChar = 0
				}
			case 'r':
				if lastRedChar == '"' {
					lastRedChar = 'r'
				} else {
					lastRedChar = 0
				}
			case 'e':
				if lastRedChar == 'r' {
					lastRedChar = 'e'
				} else {
					lastRedChar = 0
				}
			case 'd':
				if lastRedChar == 'e' {
					lastRedChar = 'd'
				} else {
					lastRedChar = 0
				}
			}
		}
		return true
	})
	return strconv.Itoa(stackBottom.total)
}
