package twentytwentyone

import "github.com/biesnecker/godvent/types"

func GetSolutions() map[string]types.Solution {
	return map[string]types.Solution{
		"year_2021_day_one_a":   {Input: "input/2021/day_one.txt", Solution: DayOneA},
		"year_2021_day_one_b":   {Input: "input/2021/day_one.txt", Solution: DayOneB},
		"year_2021_day_two_a":   {Input: "input/2021/day_two.txt", Solution: DayTwoA},
		"year_2021_day_two_b":   {Input: "input/2021/day_two.txt", Solution: DayTwoB},
		"year_2021_day_three_a": {Input: "input/2021/day_three.txt", Solution: DayThreeA},
		"year_2021_day_three_b": {Input: "input/2021/day_three.txt", Solution: DayThreeB},
		"year_2021_day_four_a":  {Input: "input/2021/day_four.txt", Solution: DayFourA},
		"year_2021_day_four_b":  {Input: "input/2021/day_four.txt", Solution: DayFourB},
		"year_2021_day_five_a":  {Input: "input/2021/day_five.txt", Solution: DayFiveA},
		"year_2021_day_five_b":  {Input: "input/2021/day_five.txt", Solution: DayFiveB},
	}
}
