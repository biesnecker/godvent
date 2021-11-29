package twentysixteen

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

const (
	SWAP_BY_POS byte = iota
	SWAP_BY_LETTER
	ROTATE_LEFT
	ROTATE_RIGHT
	ROTATE_BY_LETTER
	REVERSE
	MOVE
)

func rotateBytesLeft(original []byte, amount int) []byte {
	return rotateBytes(original, len(original)-amount)
}

func rotateBytesByLetter(original []byte, target byte) []byte {
	x := bytes.IndexByte(original, target) + 1
	if x > 4 {
		x++
	}
	return rotateBytes(original, x)
}

func rotateBytes(original []byte, amount int) []byte {
	scratch := make([]byte, len(original))
	for i, b := range original {
		idx := (i + amount) % len(original)
		scratch[idx] = b
	}
	return scratch
}

func reverseBytes(original []byte, start, end int) []byte {
	scratch := make([]byte, len(original))
	for i := range original {
		if i < start || i > end {
			scratch[i] = original[i]
		} else {
			scratch[i] = original[end-(i-start)]
		}
	}
	return scratch
}

func moveBytes(original []byte, from, to int) []byte {
	if from == to {
		return original
	}
	scratch := make([]byte, len(original))
	var offset int
	for i := range original {
		if i == from {
			if offset == 0 {
				offset++
				scratch[i] = original[i+offset]
			} else {
				scratch[i] = original[i+offset]
				offset++
			}
		} else if i == to {
			offset--
			scratch[i] = original[from]
		} else {
			scratch[i] = original[i+offset]
		}
	}
	return scratch
}

func readInputDayTwentyOne(fp *bufio.Reader) []byte {
	var res []byte
	utils.ReadStrings(fp, func(s string) {
		parts := strings.Split(s, " ")
		var op, a, b byte
		switch parts[0] {
		case "swap":
			switch parts[1] {
			case "position":
				op = SWAP_BY_POS
				a = utils.ReadByte(parts[2])
				b = utils.ReadByte(parts[5])
			case "letter":
				op = SWAP_BY_LETTER
				a = parts[2][0]
				b = parts[5][0]
			}
		case "rotate":
			switch parts[1] {
			case "left":
				op = ROTATE_LEFT
				a = utils.ReadByte(parts[2])
				b = 0
			case "right":
				op = ROTATE_RIGHT
				a = utils.ReadByte(parts[2])
				b = 0
			case "based":
				op = ROTATE_BY_LETTER
				a = parts[6][0]
				b = 0
			}
		case "reverse":
			op = REVERSE
			a = utils.ReadByte(parts[2])
			b = utils.ReadByte(parts[4])
		case "move":
			op = MOVE
			a = utils.ReadByte(parts[2])
			b = utils.ReadByte(parts[5])
		}
		res = append(res, op, a, b)
	})
	return res
}

func findSolutionDayTwentyOne(prog []byte, input string) []byte {
	ib := []byte(input)

	for pc := 0; pc < len(prog); pc += 3 {
		op := prog[pc]
		a := prog[pc+1]
		b := prog[pc+2]

		switch op {
		case SWAP_BY_POS:
			ib[a], ib[b] = ib[b], ib[a]
		case SWAP_BY_LETTER:
			x := bytes.IndexByte(ib, a)
			y := bytes.IndexByte(ib, b)
			ib[x], ib[y] = ib[y], ib[x]
		case ROTATE_LEFT:
			ib = rotateBytesLeft(ib, int(a))
		case ROTATE_RIGHT:
			ib = rotateBytes(ib, int(a))
		case ROTATE_BY_LETTER:
			ib = rotateBytesByLetter(ib, a)
		case REVERSE:
			ib = reverseBytes(ib, int(a), int(b))
		case MOVE:
			ib = moveBytes(ib, int(a), int(b))
		}
	}

	return ib
}

func DayTwentyOneA(fp *bufio.Reader) string {
	return string(findSolutionDayTwentyOne(readInputDayTwentyOne(fp), "abcdefgh"))
}

func DayTwentyOneB(fp *bufio.Reader) string {
	target := []byte("fbgdceah")
	original := []byte("abcdefgh")
	candidate := make([]byte, len(original))
	prog := readInputDayTwentyOne(fp)
	utils.Permutations(len(original), func(ns []int) bool {
		for i, x := range ns {
			candidate[i] = original[x]
		}
		return !bytes.Equal(findSolutionDayTwentyOne(prog, string(candidate)), target)
	})
	return string(candidate)
}
