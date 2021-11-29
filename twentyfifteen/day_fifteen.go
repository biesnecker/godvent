package twentyfifteen

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type cookieIngredient struct {
	capacity, durability, flavor, texture, calories int
}

func readIngredientsDayFifteen(fp *bufio.Reader) []cookieIngredient {
	ingredients := make([]cookieIngredient, 0, 5)
	utils.ReadStrings(fp, func(s string) {
		var name string
		var capacity, durability, flavor, texture, calories int
		fmt.Sscanf(
			s,
			"%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&name, &capacity, &durability, &flavor, &texture, &calories)
		ingredients = append(ingredients,
			cookieIngredient{
				capacity:   capacity,
				durability: durability,
				flavor:     flavor,
				texture:    texture,
				calories:   calories})
	})
	return ingredients
}

func scoreCookie(ingredients []cookieIngredient, amounts []int) (int, int) {
	if len(ingredients) != len(amounts) {
		log.Fatalln("len(ingredients) != len(amounts")
	}
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	calories := 0
	for i := 0; i < len(ingredients); i++ {
		capacity += ingredients[i].capacity * amounts[i]
		durability += ingredients[i].durability * amounts[i]
		flavor += ingredients[i].flavor * amounts[i]
		texture += ingredients[i].texture * amounts[i]
		calories += ingredients[i].calories * amounts[i]
	}
	if capacity < 1 || durability < 1 || flavor < 1 || texture < 1 {
		return 0, 0
	}
	score := capacity * durability * flavor * texture
	return score, calories
}

func compareCookie(
	ingredients []cookieIngredient,
	challenger []int,
	baseScore, baseCalories, requiredCalories int,
) bool {
	challengerScore, challengerCalories := scoreCookie(ingredients, challenger)

	baseDelta := math.Abs(float64(requiredCalories - baseCalories))
	challengerDelta := math.Abs(float64(requiredCalories - challengerCalories))

	if requiredCalories == 0 || baseDelta == challengerDelta {
		if challengerScore > baseScore {
			return true
		} else {
			return false
		}
	} else if baseDelta < challengerDelta {
		return false
	} else {
		return true
	}
}

func findOptimalCookie(ingredients []cookieIngredient, amountTotal, requiredCalories int) int {
	step := 16

	nIngredients := len(ingredients)

	currentCookie := make([]int, nIngredients)
	for i := 0; i < nIngredients; i++ {
		currentCookie[i] = amountTotal / nIngredients
	}
	if amountTotal%nIngredients != 0 {
		currentCookie[0] += amountTotal % nIngredients
	}

	var currentScore int
	var currentCalories int

	for step > 0 {
		currentScore, currentCalories = scoreCookie(ingredients, currentCookie)
		foundBetter := false
		var nextGenBestCookie []int
		for i := range currentCookie {
			for j := range currentCookie {
				if i == j {
					continue
				}
				if currentCookie[i] < step || currentCookie[j]+step > 100 {
					continue
				}
				newCookie := make([]int, nIngredients)
				copy(newCookie, currentCookie)
				newCookie[i] -= step
				newCookie[j] += step
				if compareCookie(
					ingredients,
					newCookie,
					currentScore,
					currentCalories,
					requiredCalories) {
					nextGenBestCookie = newCookie
					foundBetter = true
				}
			}
		}
		if foundBetter {
			currentCookie = nextGenBestCookie
		} else {
			step /= 2
		}
	}
	return currentScore
}

func DayFifteenA(fp *bufio.Reader) string {
	ingredients := readIngredientsDayFifteen(fp)
	return strconv.Itoa(findOptimalCookie(ingredients, 100, 0))
}

func DayFifteenB(fp *bufio.Reader) string {
	ingredients := readIngredientsDayFifteen(fp)
	return strconv.Itoa(findOptimalCookie(ingredients, 100, 500))
}
