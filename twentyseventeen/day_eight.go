package twentyseventeen

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func findSolutionDayEight(fp *bufio.Reader) (int, int) {
	registers := make(map[string]int)
	maxRegisterValue := math.MinInt64
	utils.ReadStrings(fp, func(s string) {
		var regA, op, regB, cmp string
		var opArg, cmpArg int
		fmt.Sscanf(s, "%s %s %d if %s %s %d",
			&regA, &op, &opArg, &regB, &cmp, &cmpArg)

		regBValue := registers[regB]
		satisfied := false
		switch cmp {
		case "==":
			satisfied = regBValue == cmpArg
		case ">":
			satisfied = regBValue > cmpArg
		case "<":
			satisfied = regBValue < cmpArg
		case "<=":
			satisfied = regBValue <= cmpArg
		case ">=":
			satisfied = regBValue >= cmpArg
		case "!=":
			satisfied = regBValue != cmpArg
		default:
			log.Fatalf("Unknown comparison operator: %s (%s)\n", cmp, s)
		}
		if satisfied {
			switch op {
			case "inc":
				registers[regA] += opArg
			case "dec":
				registers[regA] -= opArg
			default:
				log.Fatalf("Unknown op: %s\n", op)
			}
			if registers[regA] > maxRegisterValue {
				maxRegisterValue = registers[regA]
			}
		}
	})
	maxValue := math.MinInt64
	for _, v := range registers {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue, maxRegisterValue
}

func DayEightA(fp *bufio.Reader) string {
	max, _ := findSolutionDayEight(fp)
	return strconv.Itoa(max)
}

func DayEightB(fp *bufio.Reader) string {
	_, maxEver := findSolutionDayEight(fp)
	return strconv.Itoa(maxEver)
}
