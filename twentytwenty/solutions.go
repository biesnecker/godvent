package twentytwenty

import "github.com/biesnecker/godvent/types"

func GetSolutions() map[string]types.Solution {
	return map[string]types.Solution{
		"year_2020_day_one_a": {Input: "input/2020/day_one.txt", Solution: DayOneA},
		"year_2020_day_one_b": {Input: "input/2020/day_one.txt", Solution: DayOneB},
		"year_2020_day_two_a": {Input: "input/2020/day_two.txt", Solution: DayTwoA},
		"year_2020_day_two_b": {Input: "input/2020/day_two.txt", Solution: DayTwoB},
	}
}