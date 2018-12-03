package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	NoOutput := true
	if len(os.Args) > 1 {
		NoOutput = false
	}
	var state State
	state = NewSunState()
	bot := Bot{
		Records: map[string]StateRecord{},
	}
	if dataJSON, err := ioutil.ReadFile("dataset.json"); err == nil {
		json.Unmarshal(dataJSON, &bot)
	}

	runTimes := 1000
	counter := 0
	for counter != runTimes {
		obs := state.GetSideEffects()
		exp := state.State()
		guess := bot.GuessState(obs)
		bot.NewRecord(obs, exp)
		if NoOutput {
			fmt.Printf("%-6v %-6v %-6v\n", obs, exp, exp == guess)
		}
		state = state.Transition()
		counter++
	}
	botData, _ := json.MarshalIndent(bot, "", "  ")
	ioutil.WriteFile("dataset.json", botData, 0644)
}
