package twentyfifteen

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/biesnecker/godvent/utils"
)

type intKey [2]int

func parseInputDayThirteen(handler func(string, string, int)) func(string) {
	return func(s string) {
		var nameOne, nameTwo, op string
		var delta int
		fmt.Sscanf(s, "%s would %s %d happiness units by sitting next to %s",
			&nameOne, &op, &delta, &nameTwo)
		if op == "lose" {
			delta = -delta
		}
		nameTwo = strings.TrimRightFunc(
			nameTwo,
			func(r rune) bool { return !unicode.IsLetter(r) })
		handler(nameOne, nameTwo, delta)
	}
}

func buildHappinessMap(happiness map[intKey]int, numIds *int) func(string, string, int) {
	nameIds := make(map[string]int)
	nextId := 0
	return func(nameOne string, nameTwo string, delta int) {
		var ok bool
		var nameIdOne, nameIdTwo int
		if nameIdOne, ok = nameIds[nameOne]; !ok {
			nameIds[nameOne] = nextId
			nameIdOne = nextId
			nextId++
			*numIds++
		}

		if nameIdTwo, ok = nameIds[nameTwo]; !ok {
			nameIds[nameTwo] = nextId
			nameIdTwo = nextId
			nextId++
			*numIds++
		}
		happiness[intKey{nameIdOne, nameIdTwo}] = delta
	}
}

func calculateMaximumHappiness(happiness map[intKey]int, numIds int) int {
	type perm struct {
		firstId, prevId, total int
		mask                   uint16
	}

	maxHappiness := math.MinInt32
	fullMask := (uint16(1) << numIds) - 1

	q := make(chan *perm, 500000)

	q <- &perm{firstId: -1, prevId: -1, total: 0, mask: 0}

	finished := false

	for {
		if finished {
			break
		}
		select {
		case p := <-q:
			if p.mask == fullMask {
				// Calculate the effect of the first and last people on each
				// other.
				p.total += happiness[intKey{p.firstId, p.prevId}]
				p.total += happiness[intKey{p.prevId, p.firstId}]
				if p.total > maxHappiness {
					maxHappiness = p.total
				}
				break
			}

			for i := 0; i < numIds; i++ {
				check := uint16(1) << i
				if p.mask&check > 0 {
					continue // Already added here.
				}

				newTotal := p.total
				newFirstId := p.firstId
				newPrevId := p.prevId

				if newPrevId != -1 {
					newTotal += happiness[intKey{p.prevId, i}]
					newTotal += happiness[intKey{i, p.prevId}]
				} else {
					newFirstId = i
				}
				newPrevId = i

				q <- &perm{
					firstId: newFirstId,
					prevId:  newPrevId,
					total:   newTotal,
					mask:    p.mask | check}
			}
		default:
			finished = true
		}
	}
	return maxHappiness
}

func DayThirteenA(fp *bufio.Reader) string {

	happiness := make(map[intKey]int)
	numIds := 0
	utils.ReadStrings(
		fp,
		parseInputDayThirteen(
			buildHappinessMap(happiness, &numIds)))

	return strconv.Itoa(calculateMaximumHappiness(happiness, numIds))
}

func DayThirteenB(fp *bufio.Reader) string {
	happiness := make(map[intKey]int)
	numIds := 0
	utils.ReadStrings(
		fp,
		parseInputDayThirteen(
			buildHappinessMap(happiness, &numIds)))

	myId := numIds
	for i := 0; i < numIds; i++ {
		happiness[intKey{myId, i}] = 0
		happiness[intKey{i, myId}] = 0
	}

	return strconv.Itoa(calculateMaximumHappiness(happiness, numIds+1))
}
