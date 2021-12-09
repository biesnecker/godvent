package twentyseventeen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type substateD25 struct {
	write     int
	move      int
	nextState byte
}

type stateD25 struct {
	onZero substateD25
	onOne  substateD25
}

func readInputDay25(fp *bufio.Reader) (byte, int, map[byte]stateD25) {
	res := make(map[byte]stateD25)
	lines := utils.ReadStringsAsSlice(fp)
	var initialState byte
	var iterations int
	fmt.Sscanf(lines[0], "Begin in state %c.", &initialState)
	fmt.Sscanf(lines[1], "Perform a diagnostic checksum after %d steps.", &iterations)

	lines = lines[3:]

	for len(lines) > 0 {
		var stateId byte
		var dir string
		fmt.Sscanf(lines[0], "In state %c:", &stateId)
		state := stateD25{}
		// - Write the value 1.
		// - Move one slot to the right.
		// - Continue with state B.
		fmt.Sscanf(lines[2], "- Write the value %d.", &state.onZero.write)
		fmt.Sscanf(lines[3], "- Move one slot to the %s", &dir)
		if dir == "left." {
			state.onZero.move = -1
		} else {
			state.onZero.move = 1
		}
		fmt.Sscanf(lines[4], "- Continue with state %c.", &state.onZero.nextState)

		fmt.Sscanf(lines[6], "- Write the value %d.", &state.onOne.write)
		fmt.Sscanf(lines[7], "- Move one slot to the %s", &dir)
		if dir == "left." {
			state.onOne.move = -1
		} else {
			state.onOne.move = 1
		}
		fmt.Sscanf(lines[8], "- Continue with state %c.", &state.onOne.nextState)
		lines = lines[9:]
		res[stateId] = state
		if len(lines) > 0 {
			lines = lines[1:]
		}
	}
	return initialState, iterations, res
}

func DayTwentyFiveA(fp *bufio.Reader) string {
	state, iterations, stateMap := readInputDay25(fp)

	mem := make(map[int]bool)
	step := 0
	loc := 0
	for step < iterations {
		step++
		cv := mem[loc]
		var instr substateD25
		if cv {
			instr = stateMap[state].onOne
		} else {
			instr = stateMap[state].onZero
		}
		if instr.write == 1 {
			mem[loc] = true
		} else {
			delete(mem, loc)
		}
		loc += instr.move
		state = instr.nextState
	}
	return strconv.Itoa(len(mem))
}

func DayTwentyFiveB(fp *bufio.Reader) string {
	return "There is no 2017 Day 25 Part B puzzle"
}
