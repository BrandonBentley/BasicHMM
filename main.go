package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

var obervablesPerCycle = 10
var runTimes = 100
var stateData GeneratedStates
var lengthOfStateSlices = 20

func main() {
	correctCount := 0
	NoOutput := true
	if len(os.Args) > 1 {
		if os.Args[1] == "gen" {
			GenerateStates()
			os.Exit(0)
		}
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
	data, _ := ioutil.ReadFile("data/states.json")
	err := json.Unmarshal(data, &stateData)
	if err != nil {
		fmt.Println("Couldn't Initialize States")
		os.Exit(1)
	}
	var state State
	state = stateData.States[0]
	// obervablesPerCycle = len(stateData.States) * len(stateData.Obs)
	// if obervablesPerCycle < 10 {
	// 	obervablesPerCycle = 10
	// } else if obervablesPerCycle > 20 {
	// 	obervablesPerCycle = 20
	// }
	bot := NewBot()
	if dataJSON, err := ioutil.ReadFile("dataset.json"); err == nil {
		json.Unmarshal(dataJSON, &bot)
	}

	counter := 0
	for counter != runTimes {
		obs := orderString(state.GetSideEffects())
		exp := state.State()
		guess := bot.GuessState(obs)
		bot.NewRecord(obs, exp)
		if NoOutput {
			if exp == guess {
				correctCount++
			}
			fmt.Printf("%-25v %-15v %-6v\n", obs, exp, exp == guess)
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

func orderString(s string) string {
	sorted := ""
	stringSlice := []string{}
	for _, v := range s {
		stringSlice = append(stringSlice, string(v))
	}
	sort.Strings(stringSlice)
	for _, v := range stringSlice {
		sorted += v
	}
	return sorted
}
