package twentyfifteen

import (
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

type rpgItem struct {
	cost, damage, armor int
}

type rpgPlayer struct {
	hp, damage, armor int
}

var weapons = []rpgItem{
	{cost: 8, damage: 4, armor: 0},
	{cost: 10, damage: 5, armor: 0},
	{cost: 25, damage: 6, armor: 0},
	{cost: 40, damage: 7, armor: 0},
	{cost: 74, damage: 8, armor: 0}}

var armors = []rpgItem{
	{cost: 0, damage: 0, armor: 0},
	{cost: 13, damage: 0, armor: 1},
	{cost: 31, damage: 0, armor: 2},
	{cost: 53, damage: 0, armor: 3},
	{cost: 75, damage: 0, armor: 4},
	{cost: 102, damage: 0, armor: 5}}

var rings = []rpgItem{
	{cost: 0, damage: 0, armor: 0},
	{cost: 0, damage: 0, armor: 0},
	{cost: 25, damage: 1, armor: 0},
	{cost: 50, damage: 2, armor: 0},
	{cost: 100, damage: 3, armor: 0},
	{cost: 20, damage: 0, armor: 1},
	{cost: 40, damage: 0, armor: 2},
	{cost: 80, damage: 0, armor: 3}}

func simulateBattle(player *rpgPlayer, boss *rpgPlayer) bool {
	youHp := player.hp
	bossHp := boss.hp

	for {
		bossHp -= utils.MaxInt(1, player.damage-boss.armor)
		if bossHp <= 0 {
			return true
		}
		youHp -= utils.MaxInt(1, boss.damage-player.armor)
		if youHp <= 0 {
			return false
		}
	}
}

func chooseItems(
	boss *rpgPlayer, armor int, weapon int, ringOne int, ringTwo int,
	cost int, minCostToWin *int, maxCostToLose *int,
) {
	if armor == -1 {
		for i, a := range armors {
			chooseItems(boss, i, weapon, ringOne, ringTwo,
				cost+a.cost, minCostToWin, maxCostToLose)
		}
		return
	} else if weapon == -1 {
		for i, w := range weapons {
			chooseItems(boss, armor, i, ringOne, ringTwo,
				cost+w.cost, minCostToWin, maxCostToLose)
		}
		return
	} else if ringOne == -1 {
		for i, r := range rings {
			if i == ringTwo {
				continue
			}
			chooseItems(boss, armor, weapon, i, ringTwo,
				cost+r.cost, minCostToWin, maxCostToLose)
		}
		return
	} else if ringTwo == -1 {
		for i, r := range rings {
			if i == ringOne {
				continue
			}
			chooseItems(boss, armor, weapon, ringOne, i,
				cost+r.cost, minCostToWin, maxCostToLose)
		}
		return
	}

	// We've now chosen everything, it's time to fight.
	player := rpgPlayer{
		hp: 100,
		armor: weapons[weapon].armor + armors[armor].armor +
			rings[ringOne].armor + rings[ringTwo].armor,
		damage: weapons[weapon].damage + armors[armor].damage +
			rings[ringOne].damage + rings[ringTwo].damage}

	if simulateBattle(&player, boss) {
		if cost < *minCostToWin {
			*minCostToWin = cost
		}
	} else {
		if cost > *maxCostToLose {
			*maxCostToLose = cost
		}
	}

}

func DayTwentyOneA(fp *bufio.Reader) string {
	minCostToWin := math.MaxInt32
	maxCostToLose := math.MinInt32
	var boss rpgPlayer
	utils.ReadStrings(fp, func(s string) {
		parts := strings.SplitAfter(s, ": ")
		if n, err := strconv.Atoi(parts[1]); err != nil {
			log.Fatalln(err)
		} else {
			if strings.HasPrefix(parts[0], "Hit Points") {
				boss.hp = n
			} else if strings.HasPrefix(parts[0], "Damage") {
				boss.damage = n
			} else if strings.HasPrefix(parts[0], "Armor") {
				boss.armor = n
			} else {
				panic("Bad input")
			}
		}

	})
	chooseItems(&boss, -1, -1, -1, -1, 0, &minCostToWin, &maxCostToLose)
	return strconv.Itoa(minCostToWin)
}

func DayTwentyOneB(fp *bufio.Reader) string {
	minCostToWin := math.MaxInt32
	maxCostToLose := math.MinInt32
	var boss rpgPlayer
	utils.ReadStrings(fp, func(s string) {
		parts := strings.SplitAfter(s, ": ")
		if n, err := strconv.Atoi(parts[1]); err != nil {
			log.Fatalln(err)
		} else {
			if strings.HasPrefix(parts[0], "Hit Points") {
				boss.hp = n
			} else if strings.HasPrefix(parts[0], "Damage") {
				boss.damage = n
			} else if strings.HasPrefix(parts[0], "Armor") {
				boss.armor = n
			} else {
				panic("Bad input")
			}
		}

	})
	chooseItems(&boss, -1, -1, -1, -1, 0, &minCostToWin, &maxCostToLose)
	return strconv.Itoa(maxCostToLose)
}
