package twentytwentyone

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

type squareD4 struct {
	value  int
	marked bool
}

type boardD4 struct {
	squares [5][5]squareD4
	won     bool
}

// Returns true if the mark results in the board winning.
func (b *boardD4) mark(value int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.squares[i][j].value == value {
				b.squares[i][j].marked = true

				rowWin := true
				columnWin := true
				for x := 0; x < 5; x++ {
					if rowWin && !b.squares[i][x].marked {
						rowWin = false
					}
					if columnWin && !b.squares[x][j].marked {
						columnWin = false
					}
				}

				if rowWin || columnWin {
					b.won = true
					return true
				} else {
					return false
				}
			}
		}
	}
	return false
}

func (b *boardD4) score(multiplier int) int {
	var total int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.squares[i][j].marked {
				total += b.squares[i][j].value
			}
		}
	}
	return total * multiplier
}

func readInputDayFour(fp *bufio.Reader) ([]int, []boardD4) {
	var nums []int
	var boards []boardD4

	lines := utils.ReadStringsAsSlice(fp)
	first := lines[0]

	for _, n := range strings.Split(first, ",") {
		nums = append(nums, utils.ReadInt(n))
	}

	lines = lines[2:]
	for len(lines) > 0 {
		var newboard boardD4
		boardLines := lines[:5]
		for i, line := range boardLines {
			var a, b, c, d, e int
			_, err := fmt.Sscanf(
				line,
				"%d %d %d %d %d",
				&a, &b, &c, &d, &e,
			)

			if err != nil {
				log.Fatalln(err)
			}

			newboard.squares[i][0].value = a
			newboard.squares[i][1].value = b
			newboard.squares[i][2].value = c
			newboard.squares[i][3].value = d
			newboard.squares[i][4].value = e

		}
		boards = append(boards, newboard)
		lines = lines[5:]
		if len(lines) > 0 {
			lines = lines[1:]
		}
	}
	return nums, boards
}

func DayFourA(fp *bufio.Reader) string {
	nums, boards := readInputDayFour(fp)
	for _, num := range nums {
		for bid := range boards {
			if boards[bid].mark(num) {
				return strconv.Itoa(boards[bid].score(num))
			}
		}
	}
	return ""
}

func DayFourB(fp *bufio.Reader) string {
	nums, boards := readInputDayFour(fp)
	boardsLeft := len(boards)
	for _, num := range nums {
		for bid := range boards {
			if boards[bid].won {
				continue
			}
			if boards[bid].mark(num) {
				if boardsLeft == 1 {
					return strconv.Itoa(boards[bid].score(num))
				} else {
					boardsLeft--
				}
			}
		}
	}
	return ""
}
