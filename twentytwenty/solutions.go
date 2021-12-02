package twentytwenty

import "github.com/biesnecker/godvent/types"

func GetSolutions() map[string]types.Solution {
	return map[string]types.Solution{
		"year_2020_day_one_a":    {Input: "input/2020/day_one.txt", Solution: DayOneA},
		"year_2020_day_one_b":    {Input: "input/2020/day_one.txt", Solution: DayOneB},
		"year_2020_day_two_a":    {Input: "input/2020/day_two.txt", Solution: DayTwoA},
		"year_2020_day_two_b":    {Input: "input/2020/day_two.txt", Solution: DayTwoB},
		"year_2020_day_three_a":  {Input: "input/2020/day_three.txt", Solution: DayThreeA},
		"year_2020_day_three_b":  {Input: "input/2020/day_three.txt", Solution: DayThreeB},
		"year_2020_day_four_a":   {Input: "input/2020/day_four.txt", Solution: DayFourA},
		"year_2020_day_four_b":   {Input: "input/2020/day_four.txt", Solution: DayFourB},
		"year_2020_day_five_a":   {Input: "input/2020/day_five.txt", Solution: DayFiveA},
		"year_2020_day_five_b":   {Input: "input/2020/day_five.txt", Solution: DayFiveB},
		"year_2020_day_six_a":    {Input: "input/2020/day_six.txt", Solution: DaySixA},
		"year_2020_day_six_b":    {Input: "input/2020/day_six.txt", Solution: DaySixB},
		"year_2020_day_seven_a":  {Input: "input/2020/day_seven.txt", Solution: DaySevenA},
		"year_2020_day_seven_b":  {Input: "input/2020/day_seven.txt", Solution: DaySevenB},
		"year_2020_day_eight_a":  {Input: "input/2020/day_eight.txt", Solution: DayEightA},
		"year_2020_day_eight_b":  {Input: "input/2020/day_eight.txt", Solution: DayEightB},
		"year_2020_day_nine_a":   {Input: "input/2020/day_nine.txt", Solution: DayNineA},
		"year_2020_day_nine_b":   {Input: "input/2020/day_nine.txt", Solution: DayNineB},
		"year_2020_day_ten_a":    {Input: "input/2020/day_ten.txt", Solution: DayTenA},
		"year_2020_day_ten_b":    {Input: "input/2020/day_ten.txt", Solution: DayTenB},
		"year_2020_day_eleven_a": {Input: "input/2020/day_eleven.txt", Solution: DayElevenA},
		"year_2020_day_eleven_b": {Input: "input/2020/day_eleven.txt", Solution: DayElevenB},
		"year_2020_day_twelve_a": {Input: "input/2020/day_twelve.txt", Solution: DayTwelveA},
		"year_2020_day_twelve_b": {Input: "input/2020/day_twelve.txt", Solution: DayTwelveB},
	}
}
