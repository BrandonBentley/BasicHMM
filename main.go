package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var obervablesPerCycle = 10
var runTimes = 100

func main() {
	correctCount := 0
	NoOutput := true
	if len(os.Args) > 1 {
		NoOutput = false
		num, err := strconv.Atoi(os.Args[1])
		if err == nil {
			runTimes = num
		} else {
			fmt.Println("NAN")
			os.Exit(1)
		}
	}
	if len(os.Args) > 2 {
		runTimes = 1000000
	}
	var state State
	state = NewSunState()
	bot := Bot{
		Records: map[string]StateRecord{},
	}
	if dataJSON, err := ioutil.ReadFile("dataset.json"); err == nil {
		json.Unmarshal(dataJSON, &bot)
	}

	counter := 0
	for counter != runTimes {
		obs := state.GetSideEffects()
		exp := state.State()
		guess := bot.GuessState(obs)
		bot.NewRecord(obs, exp)
		if NoOutput {
			if exp == guess {
				correctCount++
			}
			fmt.Printf("%-11v %-6v %-6v\n", obs, exp, exp == guess)
		}
		state = state.Transition()
		counter++
	}
	botData, _ := json.MarshalIndent(bot, "", "  ")
	ioutil.WriteFile("dataset.json", botData, 0644)
	if NoOutput {
		fmt.Printf("\n%26v: %v\n%26v: %v%%\n", "Total Data Points Recorded", bot.PointCount, "Current Accuracy", correctCount)
	}
}
