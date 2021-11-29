package twentyseventeen

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"unicode"

	"github.com/biesnecker/godvent/utils"
)

const (
	UP int = iota
	RIGHT
	DOWN
	LEFT
)

func readInputDayNineteen(fp *bufio.Reader) ([][]byte, int) {
	var res [][]byte
	var startPos int
	utils.ReadStringsWithIndexNoTrim(fp, func(i int, s string) {
		bslice := []byte(strings.TrimSuffix(s, "\n"))
		res = append(res, bslice)
		if i == 0 {
			startPos = bytes.IndexByte(bslice, '|')
		}
	})
	return res, startPos
}

func findSolutionDayNineteen(grid [][]byte, startPos int) (string, int) {
	x := 0
	y := startPos
	dir := DOWN
	steps := 1 // starts at one because initial step down counts
	finished := false
	var result []byte

	for !finished {
		// Can we continue in the same direction as before?
		newX, newY := x, y
		switch dir {
		case UP:
			newX--
		case DOWN:
			newX++
		case LEFT:
			newY--
		case RIGHT:
			newY++
		}
		if newX == len(grid) || newY == len(grid[x]) || grid[newX][newY] == ' ' {
			// Nope, we've hit a blank space. Need to figure out how to turn.
			newX, newY = x, y
			switch dir {
			case UP, DOWN:
				if newY-1 >= 0 && grid[newX][newY-1] != ' ' {
					dir = LEFT
					newY = newY - 1
				} else if newY+1 < len(grid[newX]) && grid[newX][newY+1] != ' ' {
					dir = RIGHT
					newY = newY + 1
				} else {
					finished = true
				}
			case LEFT, RIGHT:
				if newX+1 < len(grid) && grid[newX+1][newY] != ' ' {
					dir = DOWN
					newX = newX + 1
				} else if newX-1 >= 0 && grid[newX-1][newY] != ' ' {
					dir = UP
					newX = newX - 1
				} else {
					finished = true
				}
			}
		}
		if !finished {
			x, y = newX, newY
			steps++
			if unicode.IsLetter(rune(grid[newX][newY])) {
				result = append(result, grid[newX][newY])
			}
		}

	}

	return string(result), steps
}

func DayNineteenA(fp *bufio.Reader) string {
	grid, startPos := readInputDayNineteen(fp)
	solution, _ := findSolutionDayNineteen(grid, startPos)
	return solution
}

func DayNineteenB(fp *bufio.Reader) string {
	grid, startPos := readInputDayNineteen(fp)
	_, steps := findSolutionDayNineteen(grid, startPos)
	return strconv.Itoa(steps)
}
