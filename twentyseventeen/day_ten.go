package twentyseventeen

import (
	"bufio"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/twentyseventeen/knothash"
	"github.com/biesnecker/godvent/utils"
)

func DayTenA(fp *bufio.Reader) string {
	input := utils.ReadSingleString(fp)
	if lens, err := utils.ParseDelimitedByteString(input, ","); err != nil {
		log.Fatalln(err)
	} else {
		knot := knothash.KnotHashSparse(lens, 1)
		return strconv.Itoa(int(knot[0]) * int(knot[1]))
	}
	return ""
}

func DayTenB(fp *bufio.Reader) string {
	input := []byte(utils.ReadSingleString(fp))
	return knothash.KnotHashString(input)
}
