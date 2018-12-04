package main

import (
	"encoding/json"
	"io/ioutil"
)

var heavyUP int

func GenerateStates() {
	var sg StateGen
	j, _ := ioutil.ReadFile("data/state_proto.json")
	json.Unmarshal(j, &sg)
	allPossible := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	if sg.NumChoices > 10 {
		sg.NumChoices = 10
	}
	lengthOfStates := len(sg.States)
	genStates := GeneratedStates{
		Obs:       allPossible[:sg.NumChoices],
		StateKeys: map[string]int{},
		States:    make([]State, lengthOfStates),
	}
	for i, v := range sg.States {
		genStates.StateKeys[v] = i
	}
	for i, v := range sg.States {
		heavyUP = i
		genStates.States[i] = NewState(v, lengthOfStates, genStates.Obs)

	}
	sj, _ := json.MarshalIndent(genStates, "", "  ")
	ioutil.WriteFile("data/states.json", sj, 0644)
}

type StateGen struct {
	NumChoices int      `json:"numChoices"`
	States     []string `json:"states"`
}

type GeneratedStates struct {
	Obs       []string       `json:"obs"`
	States    []State        `json:"states"`
	StateKeys map[string]int `json:"stateKeys"`
}
