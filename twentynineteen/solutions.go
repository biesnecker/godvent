package twentynineteen

import "github.com/biesnecker/godvent/types"

func GetSolutions() map[string]types.Solution {
	return map[string]types.Solution{
		"year_2019_day_one_a":   {Input: "input/2019/day_one.txt", Solution: DayOneA},
		"year_2019_day_one_b":   {Input: "input/2019/day_one.txt", Solution: DayOneB},
		"year_2019_day_two_a":   {Input: "input/2019/day_two.txt", Solution: DayTwoA},
		"year_2019_day_two_b":   {Input: "input/2019/day_two.txt", Solution: DayTwoB},
		"year_2019_day_three_a": {Input: "input/2019/day_three.txt", Solution: DayThreeA},
		"year_2019_day_three_b": {Input: "input/2019/day_three.txt", Solution: DayThreeB},
		"year_2019_day_four_a":  {Input: "input/2019/day_four.txt", Solution: DayFourA},
		"year_2019_day_four_b":  {Input: "input/2019/day_four.txt", Solution: DayFourB},
		"year_2019_day_five_a":  {Input: "input/2019/day_five.txt", Solution: DayFiveA},
		"year_2019_day_five_b":  {Input: "input/2019/day_five.txt", Solution: DayFiveB},
		"year_2019_day_six_a":   {Input: "input/2019/day_six.txt", Solution: DaySixA},
		"year_2019_day_six_b":   {Input: "input/2019/day_six.txt", Solution: DaySixB},
		"year_2019_day_seven_a": {Input: "input/2019/day_seven.txt", Solution: DaySevenA},
		"year_2019_day_seven_b": {Input: "input/2019/day_seven.txt", Solution: DaySevenB},
		"year_2019_day_eight_a": {Input: "input/2019/day_eight.txt", Solution: DayEightA},
		"year_2019_day_eight_b": {Input: "input/2019/day_eight.txt", Solution: DayEightB},
		"year_2019_day_nine_a":  {Input: "input/2019/day_nine.txt", Solution: DayNineA},
		"year_2019_day_nine_b":  {Input: "input/2019/day_nine.txt", Solution: DayNineB},
	}
}
