package twentysixteen

import "github.com/biesnecker/godvent/types"

func GetSolutions() map[string]types.Solution {
	return map[string]types.Solution{
		"year_2016_day_one_a":         {Input: "input/2016/day_one.txt", Solution: DayOneA},
		"year_2016_day_one_b":         {Input: "input/2016/day_one.txt", Solution: DayOneB},
		"year_2016_day_two_a":         {Input: "input/2016/day_two.txt", Solution: DayTwoA},
		"year_2016_day_two_b":         {Input: "input/2016/day_two.txt", Solution: DayTwoB},
		"year_2016_day_three_a":       {Input: "input/2016/day_three.txt", Solution: DayThreeA},
		"year_2016_day_three_b":       {Input: "input/2016/day_three.txt", Solution: DayThreeB},
		"year_2016_day_four_a":        {Input: "input/2016/day_four.txt", Solution: DayFourA},
		"year_2016_day_four_b":        {Input: "input/2016/day_four.txt", Solution: DayFourB},
		"year_2016_day_five_a":        {Input: "input/2016/day_five.txt", Solution: DayFiveA},
		"year_2016_day_five_b":        {Input: "input/2016/day_five.txt", Solution: DayFiveB},
		"year_2016_day_six_a":         {Input: "input/2016/day_six.txt", Solution: DaySixA},
		"year_2016_day_six_b":         {Input: "input/2016/day_six.txt", Solution: DaySixB},
		"year_2016_day_seven_a":       {Input: "input/2016/day_seven.txt", Solution: DaySevenA},
		"year_2016_day_seven_b":       {Input: "input/2016/day_seven.txt", Solution: DaySevenB},
		"year_2016_day_eight_a":       {Input: "input/2016/day_eight.txt", Solution: DayEightA},
		"year_2016_day_eight_b":       {Input: "input/2016/day_eight.txt", Solution: DayEightB},
		"year_2016_day_nine_a":        {Input: "input/2016/day_nine.txt", Solution: DayNineA},
		"year_2016_day_nine_b":        {Input: "input/2016/day_nine.txt", Solution: DayNineB},
		"year_2016_day_ten_a":         {Input: "input/2016/day_ten.txt", Solution: DayTenA},
		"year_2016_day_ten_b":         {Input: "input/2016/day_ten.txt", Solution: DayTenB},
		"year_2016_day_eleven_a":      {Input: "input/2016/day_eleven.txt", Solution: DayElevenA},
		"year_2016_day_eleven_b":      {Input: "input/2016/day_eleven.txt", Solution: DayElevenB},
		"year_2016_day_twelve_a":      {Input: "input/2016/day_twelve.txt", Solution: DayTwelveA},
		"year_2016_day_twelve_b":      {Input: "input/2016/day_twelve.txt", Solution: DayTwelveB},
		"year_2016_day_thirteen_a":    {Input: "input/2016/day_thirteen.txt", Solution: DayThirteenA},
		"year_2016_day_thirteen_b":    {Input: "input/2016/day_thirteen.txt", Solution: DayThirteenB},
		"year_2016_day_fourteen_a":    {Input: "input/2016/day_fourteen.txt", Solution: DayFourteenA},
		"year_2016_day_fourteen_b":    {Input: "input/2016/day_fourteen.txt", Solution: DayFourteenB},
		"year_2016_day_fifteen_a":     {Input: "input/2016/day_fifteen.txt", Solution: DayFifteenA},
		"year_2016_day_fifteen_b":     {Input: "input/2016/day_fifteen.txt", Solution: DayFifteenB},
		"year_2016_day_sixteen_a":     {Input: "input/2016/day_sixteen.txt", Solution: DaySixteenA},
		"year_2016_day_sixteen_b":     {Input: "input/2016/day_sixteen.txt", Solution: DaySixteenB},
		"year_2016_day_seventeen_a":   {Input: "input/2016/day_seventeen.txt", Solution: DaySeventeenA},
		"year_2016_day_seventeen_b":   {Input: "input/2016/day_seventeen.txt", Solution: DaySeventeenB},
		"year_2016_day_eighteen_a":    {Input: "input/2016/day_eighteen.txt", Solution: DayEighteenA},
		"year_2016_day_eighteen_b":    {Input: "input/2016/day_eighteen.txt", Solution: DayEighteenB},
		"year_2016_day_nineteen_a":    {Input: "input/2016/day_nineteen.txt", Solution: DayNineteenA},
		"year_2016_day_nineteen_b":    {Input: "input/2016/day_nineteen.txt", Solution: DayNineteenB},
		"year_2016_day_twenty_a":      {Input: "input/2016/day_twenty.txt", Solution: DayTwentyA},
		"year_2016_day_twenty_b":      {Input: "input/2016/day_twenty.txt", Solution: DayTwentyB},
		"year_2016_day_twentyone_a":   {Input: "input/2016/day_twentyone.txt", Solution: DayTwentyOneA},
		"year_2016_day_twentyone_b":   {Input: "input/2016/day_twentyone.txt", Solution: DayTwentyOneB},
		"year_2016_day_twentytwo_a":   {Input: "input/2016/day_twentytwo.txt", Solution: DayTwentyTwoA},
		"year_2016_day_twentytwo_b":   {Input: "input/2016/day_twentytwo.txt", Solution: DayTwentyTwoB},
		"year_2016_day_twentythree_a": {Input: "input/2016/day_twentythree.txt", Solution: DayTwentyThreeA},
		"year_2016_day_twentythree_b": {Input: "input/2016/day_twentythree.txt", Solution: DayTwentyThreeB},
		"year_2016_day_twentyfour_a":  {Input: "input/2016/day_twentyfour.txt", Solution: DayTwentyFourA},
		"year_2016_day_twentyfour_b":  {Input: "input/2016/day_twentyfour.txt", Solution: DayTwentyFourB},
		"year_2016_day_twentyfive_a":  {Input: "input/2016/day_twentyfive.txt", Solution: DayTwentyFiveA},
		"year_2016_day_twentyfive_b":  {Input: "input/2016/day_twentyfive.txt", Solution: DayTwentyFiveB},
	}
}
