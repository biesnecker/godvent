package twentysixteen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

type inputD10 struct {
	value  int
	target *botD10
}

type botD10 struct {
	botId, hiValue, loValue int
	hiTarget, loTarget      interface{}
	hiToOutput, loToOutput  bool
}

type outputD10 struct {
	value int
}

type stateD10 struct {
	inputs  map[int]*inputD10
	bots    map[int]*botD10
	outputs map[int]*outputD10
}

func insertOrLookupInput(state *stateD10, inputId, targetBotId int) *inputD10 {
	if in, found := state.inputs[inputId]; found {
		return in
	}
	in := &inputD10{
		value:  inputId,
		target: insertOrLookupBot(state, targetBotId)}
	state.inputs[inputId] = in
	return in
}

func insertOrLookupBot(state *stateD10, botId int) *botD10 {
	if bot, found := state.bots[botId]; found {
		return bot
	}
	bot := &botD10{botId: botId}
	state.bots[botId] = bot
	return bot
}

func insertOrLookupOutput(state *stateD10, outputId int) *outputD10 {
	if out, found := state.outputs[outputId]; found {
		return out
	}
	out := &outputD10{}
	state.outputs[outputId] = out
	return out
}

func assignValueToBot(bot *botD10, value int) {
	if bot.loValue == 0 {
		bot.loValue = value
	} else if bot.loValue > value {
		bot.hiValue = bot.loValue
		bot.loValue = value
	} else {
		bot.hiValue = value
	}
}

func propogateFromBot(b *botD10) {
	if b.hiValue == 0 || b.loValue == 0 {
		return
	}
	if b.hiTarget == nil || b.loTarget == nil {
		return
	}

	if b.hiToOutput {
		target := b.hiTarget.(*outputD10)
		target.value = b.hiValue
	} else {
		target := b.hiTarget.(*botD10)
		assignValueToBot(target, b.hiValue)
		propogateFromBot(target)
	}

	if b.loToOutput {
		target := b.loTarget.(*outputD10)
		target.value = b.loValue
	} else {
		target := b.loTarget.(*botD10)
		assignValueToBot(target, b.loValue)
		propogateFromBot(target)
	}
}

func readInputDayTen(fp *bufio.Reader) *stateD10 {
	state := &stateD10{
		inputs:  make(map[int]*inputD10),
		bots:    make(map[int]*botD10),
		outputs: make(map[int]*outputD10)}
	utils.ReadStrings(fp, func(s string) {
		if strings.HasPrefix(s, "value") {
			var value, botId int
			fmt.Sscanf(s, "value %d goes to bot %d", &value, &botId)
			insertOrLookupInput(state, value, botId)
		} else {
			// bot 127 gives low to bot 118 and high to bot 142
			var sourceBotId, loTargetId, hiTargetId int
			var loTargetType, hiTargetType string
			fmt.Sscanf(s, "bot %d gives low to %s %d and high to %s %d",
				&sourceBotId,
				&loTargetType, &loTargetId,
				&hiTargetType, &hiTargetId)

			hiToOutput := hiTargetType == "output"
			loToOutput := loTargetType == "output"

			sourceBot := insertOrLookupBot(state, sourceBotId)

			if hiToOutput {
				out := insertOrLookupOutput(state, hiTargetId)
				sourceBot.hiTarget = out
				sourceBot.hiToOutput = true
			} else {
				out := insertOrLookupBot(state, hiTargetId)
				sourceBot.hiTarget = out
				sourceBot.hiToOutput = false
			}

			if loToOutput {
				out := insertOrLookupOutput(state, loTargetId)
				sourceBot.loTarget = out
				sourceBot.loToOutput = true
			} else {
				out := insertOrLookupBot(state, loTargetId)
				sourceBot.loTarget = out
				sourceBot.loToOutput = false
			}
		}
	})
	return state
}

func runSimulationDayTen(state *stateD10) {
	for _, v := range state.inputs {
		assignValueToBot(v.target, v.value)
		propogateFromBot(v.target)
	}
}

func DayTenA(fp *bufio.Reader) string {
	state := readInputDayTen(fp)
	runSimulationDayTen(state)
	for _, bot := range state.bots {
		if bot.hiValue == 61 && bot.loValue == 17 {
			return strconv.Itoa(bot.botId)
		}
	}
	return ""
}

func DayTenB(fp *bufio.Reader) string {
	state := readInputDayTen(fp)
	runSimulationDayTen(state)
	product := 1
	for k, output := range state.outputs {
		if k == 0 || k == 1 || k == 2 {
			product *= output.value
		}
	}
	return strconv.Itoa(product)
}
