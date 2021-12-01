package twentytwentyone

import "github.com/biesnecker/godvent/types"

func GetSolutions() map[string]types.Solution {
	return map[string]types.Solution{
		"year_2021_day_one_a": {Input: "input/2021/day_one.txt", Solution: DayOneA},
		"year_2021_day_one_b": {Input: "input/2021/day_one.txt", Solution: DayOneB},
	}
}
