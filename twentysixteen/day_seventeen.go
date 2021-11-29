package twentysixteen

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type positionD17 struct {
	pos  types.Coord
	path []byte
}

func getDoorState(salt string, path []byte) [4]bool {
	h := crypto.MD5.New()
	h.Write([]byte(salt))
	h.Write(path)
	hashBytes := make([]byte, 0, 16)
	hashBytes = h.Sum(hashBytes)
	hash := hex.EncodeToString(hashBytes)
	var res [4]bool
	for i := 0; i < 4; i++ {
		b := hash[i] - 'a'
		res[i] = b > 0 && b < 6
	}
	return res
}

func findSolutionDaySeventeen(salt string) (string, string) {
	q := make(chan *positionD17, 1000)
	q <- &positionD17{pos: types.Coord{}, path: make([]byte, 0)}

	min := math.MaxInt64
	max := math.MinInt64
	var minPath, maxPath string

complete:
	for {
		select {
		case p := <-q:
			if p.pos.X == 3 && p.pos.Y == 3 {
				if len(p.path) > max {
					max = len(p.path)
					maxPath = string(p.path)
				}
				if len(p.path) < min {
					min = len(p.path)
					minPath = string(p.path)
				}
				continue
			}

			isOpen := getDoorState(salt, p.path)
			for i := 0; i < 4; i++ {
				if !isOpen[i] {
					continue
				}

				var pos types.Coord
				var dir byte
				switch i {
				case 0:
					pos = p.pos.Down()
					if !pos.IsInBounds(0, 0, 3, 3) {
						continue
					}
					dir = 'U'
				case 1:
					pos = p.pos.Up()
					if !pos.IsInBounds(0, 0, 3, 3) {
						continue
					}
					dir = 'D'
				case 2:
					pos = p.pos.Left()
					if !pos.IsInBounds(0, 0, 3, 3) {
						continue
					}
					dir = 'L'
				case 3:
					pos = p.pos.Right()
					if !pos.IsInBounds(0, 0, 3, 3) {
						continue
					}
					dir = 'R'
				}
				newPath := make([]byte, len(p.path), len(p.path)+1)
				copy(newPath, p.path)
				newPath = append(newPath, dir)
				q <- &positionD17{pos: pos, path: newPath}
			}
		default:
			break complete
		}
	}
	return minPath, maxPath
}

func DaySeventeenA(fp *bufio.Reader) string {
	salt := utils.ReadSingleString(fp)
	min, _ := findSolutionDaySeventeen(salt)
	return min
}

func DaySeventeenB(fp *bufio.Reader) string {
	salt := utils.ReadSingleString(fp)
	_, max := findSolutionDaySeventeen(salt)
	return strconv.Itoa(len(max))
}
