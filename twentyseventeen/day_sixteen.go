package twentyseventeen

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

type opD16 int

const (
	SPIN opD16 = iota
	EXCHANGE
	PARTNER
)

type instrD16 struct {
	op   opD16
	a, b byte
}

func startingPosition() []byte {
	return []byte("abcdefghijklmnop")
}

func readInputDaySixteen(fp *bufio.Reader) []instrD16 {
	var res []instrD16
	for _, cmd := range strings.Split(utils.ReadSingleString(fp), ",") {
		switch cmd[0] {
		case 's':
			var offset int
			fmt.Sscanf(cmd, "s%d", &offset)
			res = append(res, instrD16{op: SPIN, a: byte(offset)})
		case 'x':
			var posA, posB int
			fmt.Sscanf(cmd, "x%d/%d", &posA, &posB)
			res = append(res, instrD16{op: EXCHANGE, a: byte(posA), b: byte(posB)})
		case 'p':
			var a, b byte
			fmt.Sscanf(cmd, "p%c/%c", &a, &b)
			res = append(res, instrD16{op: PARTNER, a: a, b: b})
		}
	}
	return res
}

func dance(progs []byte, moves []instrD16) []byte {
	for _, move := range moves {
		switch move.op {
		case SPIN:
			buffer := make([]byte, len(progs))
			for i := range progs {
				buffer[i] = progs[(i+(len(progs)-int(move.a)))%len(progs)]
			}
			progs = buffer
		case EXCHANGE:
			progs[move.a], progs[move.b] = progs[move.b], progs[move.a]
		case PARTNER:
			posA := bytes.IndexByte(progs, move.a)
			posB := bytes.IndexByte(progs, move.b)
			progs[posA], progs[posB] = progs[posB], progs[posA]
		}
	}
	return progs
}

func loopCount(progs []byte, moves []instrD16) int {
	state := make([]byte, len(progs))
	copy(state, progs)
	count := 0
	for {
		state = dance(state, moves)
		count++
		if bytes.Equal(progs, state) {
			return count
		}
	}
}

func danceN(progs []byte, moves []instrD16, n int) string {
	loops := loopCount(progs, moves)
	state := make([]byte, len(progs))
	copy(state, progs)
	copy(progs, state)
	for i := 0; i < n%loops; i++ {
		state = dance(state, moves)
	}
	return string(state)
}

func DaySixteenA(fp *bufio.Reader) string {
	moves := readInputDaySixteen(fp)
	return string(dance(startingPosition(), moves))
}

func DaySixteenB(fp *bufio.Reader) string {
	progs := startingPosition()
	moves := readInputDaySixteen(fp)
	return danceN(progs, moves, 1000000000)
}
