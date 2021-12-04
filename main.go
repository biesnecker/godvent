package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/biesnecker/godvent/twentyfifteen"
	"github.com/biesnecker/godvent/twentynineteen"
	"github.com/biesnecker/godvent/twentyseventeen"
	"github.com/biesnecker/godvent/twentysixteen"
	"github.com/biesnecker/godvent/twentytwenty"
	"github.com/biesnecker/godvent/twentytwentyone"
	"github.com/biesnecker/godvent/types"
)

func getAllSolutions() map[string]types.Solution {
	allSolutions := make(map[string]types.Solution)

	yearSolutions := []map[string]types.Solution{
		twentyfifteen.GetSolutions(),
		twentysixteen.GetSolutions(),
		twentyseventeen.GetSolutions(),
		twentynineteen.GetSolutions(),
		twentytwenty.GetSolutions(),
		twentytwentyone.GetSolutions()}

	for idx := range yearSolutions {
		for problemName, solution := range yearSolutions[idx] {
			allSolutions[problemName] = solution
		}
	}
	return allSolutions
}

func showUsage(progname *string, solutions *map[string]types.Solution) {
	fmt.Printf("Usage: %s [problem]\n\nAvailable problems:\n", *progname)
	for k := range *solutions {
		fmt.Printf("    %s\n", k)
	}
}

func main() {
	solutions := getAllSolutions()
	args := os.Args
	if len(args) != 2 {
		showUsage(&args[0], &solutions)
		return
	}
	problem := args[1]

	solution, ok := solutions[problem]
	if !ok {
		fmt.Printf("Unknown problem: %s\n\n", problem)
		showUsage(&args[0], &solutions)
		return
	}

	fp, err := os.Open(solution.Input)
	if err != nil {
		fmt.Printf("Error opening input file %s: %s\n", solution.Input, err)
		return
	}
	defer fp.Close()

	r := bufio.NewReader(fp)

	println(solution.Solution(r))
}
