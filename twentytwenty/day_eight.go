package twentytwenty

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/types/bitvector"
	"github.com/biesnecker/godvent/utils"
)

type instrD8 struct {
	op  int
	arg int
}

func readInputDayEight(fp *bufio.Reader) []*instrD8 {
	var prog []*instrD8
	utils.ReadStrings(fp, func(s string) {
		var opstring string
		var op, arg int
		fmt.Sscanf(s, "%s %d", &opstring, &arg)
		switch opstring {
		case "acc":
			op = 0
		case "jmp":
			op = 1
		case "nop":
			op = 2
		default:
			log.Fatalln("Unknown op: ", opstring)
		}
		prog = append(prog, &instrD8{op: op, arg: arg})
	})
	return prog
}

func DayEightA(fp *bufio.Reader) string {
	prog := readInputDayEight(fp)
	pc := 0
	acc := 0
	seen := bitvector.New(len(prog))
	for {
		if seen.Check(pc) {
			return strconv.Itoa(acc)
		}
		i := prog[pc]
		seen.Set(pc)
		switch i.op {
		case 0: // acc
			acc += i.arg
			pc++
		case 1: // jmp
			pc += i.arg
		case 2: // nop
			pc++
		default:
			log.Fatalln("Unknown op ", i.op)
		}
	}
}

func DayEightB(fp *bufio.Reader) string {
	prog := readInputDayEight(fp)
	seen := bitvector.New(len(prog))
	for idx := range prog {
		var oldArg int
		if prog[idx].op == 1 {
			prog[idx].op = 2
			oldArg = 1
		} else if prog[idx].op == 2 {
			prog[idx].op = 1
			oldArg = 2
		} else {
			continue
		}
		pc := 0
		acc := 0
		seen.Clear()
		for {
			if pc == len(prog) {
				return strconv.Itoa(acc)
			}
			if seen.Check(pc) {
				break
			}

			seen.Set(pc)
			i := prog[pc]
			switch i.op {
			case 0: // acc
				acc += i.arg
				pc++
			case 1: // jmp
				pc += i.arg
			case 2: // nop
				pc++
			default:
				log.Fatalln("Unknown op ", i.op)
			}
		}
		prog[idx].op = oldArg
	}
	return ""
}
