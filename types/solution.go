package types

import "bufio"

type SolutionFunc func(*bufio.Reader) string

type Solution struct {
	Input    string
	Solution SolutionFunc
}
