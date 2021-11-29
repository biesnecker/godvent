package twentyfifteen

import (
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

const (
	MAGIC_MISSLE int = 0
	DRAIN        int = 1
	SHIELD       int = 3
	POISON       int = 4
	RECHARGE     int = 5
)

var rpgSpells = map[int]int{
	MAGIC_MISSLE: 53,
	DRAIN:        73,
	SHIELD:       113,
	POISON:       173,
	RECHARGE:     229,
}

type rpgGameState struct {
	playerHp, playerMana              int
	bossHp, bossDamage                int
	totalManaSpent                    int
	shieldEff, poisonEff, rechargeEff int
	playerHasArmor                    bool
}

func readInputDayTwentyTwo(fp *bufio.Reader) (int, int) {
	var hp, damage int
	utils.ReadStrings(fp, func(s string) {
		parts := strings.SplitAfter(s, ": ")
		if n, err := strconv.Atoi(parts[1]); err != nil {
			log.Fatalln(err)
		} else {
			if strings.HasPrefix(parts[0], "Hit Points") {
				hp = n
			} else if strings.HasPrefix(parts[0], "Damage") {
				damage = n
			} else {
				panic("Bad input")
			}
		}
	})
	return hp, damage
}

func handleSpellEffects(gs *rpgGameState) {
	if gs.shieldEff > 0 {
		gs.playerHasArmor = true
		gs.shieldEff--
	} else {
		gs.playerHasArmor = false
	}
	if gs.poisonEff > 0 {
		gs.bossHp -= 3
		gs.poisonEff--
	}
	if gs.rechargeEff > 0 {
		gs.playerMana += 101
		gs.rechargeEff--
	}
}

func handleBossMaybeDead(gs *rpgGameState, minManaSpent *int) bool {
	if gs.bossHp <= 0 {
		if gs.totalManaSpent < *minManaSpent {
			*minManaSpent = gs.totalManaSpent
		}
		return true
	}
	return false
}

func simulateSpellBattle(gs rpgGameState, isPlayerTurn bool, hardMode bool, minManaSpent *int) {
	if gs.totalManaSpent > *minManaSpent {
		return
	}
	if isPlayerTurn && hardMode {
		gs.playerHp -= 1
		if gs.playerHp <= 0 {
			return
		}
	}
	handleSpellEffects(&gs)
	if handleBossMaybeDead(&gs, minManaSpent) {
		return
	}

	if isPlayerTurn {
		didAction := false

		// Player attack.
		for spell, cost := range rpgSpells {
			didCast := false
			if cost > gs.playerMana {
				continue
			}
			newState := gs
			if spell == MAGIC_MISSLE {
				newState.bossHp -= 4
				didCast = true
			} else if spell == DRAIN {
				newState.playerHp += 2
				newState.bossHp -= 2
				didCast = true
			} else if newState.shieldEff == 0 && spell == SHIELD {
				newState.shieldEff = 6
				didCast = true
			} else if newState.poisonEff == 0 && spell == POISON {
				newState.poisonEff = 6
				didCast = true
			} else if newState.rechargeEff == 0 && spell == RECHARGE {
				newState.rechargeEff = 5
				didCast = true
			}
			if didCast {
				newState.playerMana -= cost
				newState.totalManaSpent += cost
				simulateSpellBattle(newState, false, hardMode, minManaSpent)
			}
			didAction = didAction || didCast
		}

		if !didAction {
			// Player must do at least some action.
			return
		}
	} else {
		var armor int
		if gs.playerHasArmor {
			armor = 7
		}
		gs.playerHp -= utils.MaxInt(1, gs.bossDamage-armor)
		if gs.playerHp <= 0 {
			return
		}
		simulateSpellBattle(gs, true, hardMode, minManaSpent)
	}

}

func DayTwentyTwoA(fp *bufio.Reader) string {
	state := rpgGameState{
		playerHp:   50,
		playerMana: 500,
	}
	bossHp, bossDamage := readInputDayTwentyTwo(fp)
	state.bossHp = bossHp
	state.bossDamage = bossDamage
	minimumManaSpent := math.MaxInt64
	simulateSpellBattle(state, true, false, &minimumManaSpent)
	return strconv.Itoa(minimumManaSpent)
}

func DayTwentyTwoB(fp *bufio.Reader) string {
	state := rpgGameState{
		playerHp:   50,
		playerMana: 500,
	}
	bossHp, bossDamage := readInputDayTwentyTwo(fp)
	state.bossHp = bossHp
	state.bossDamage = bossDamage
	minimumManaSpent := math.MaxInt64
	simulateSpellBattle(state, true, true, &minimumManaSpent)
	return strconv.Itoa(minimumManaSpent)
}
