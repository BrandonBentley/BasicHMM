package main

import (
	"math/rand"
)

type Bot struct {
	PointCount int              `json:"count"`
	Records    map[string][]int `json:"records"`
}

func NewBot() *Bot {
	return &Bot{
		Records: map[string][]int{},
	}
}

func (b *Bot) NewRecord(sideEffects string, state string) {
	record, ok := b.Records[sideEffects]
	if !ok {
		record = make([]int, len(stateData.States))
	}
	record[stateData.StateKeys[state]]++
	b.Records[sideEffects] = record
	b.PointCount++
}

func (b *Bot) GuessState(sideEffects string) string {
	if _, ok := b.Records[sideEffects]; ok {
		maxVal := -1
		maxIndex := 0
		allZero := true
		for i, v := range b.Records[sideEffects] {
			if v > maxVal {
				maxVal = v
				maxIndex = i
			}
			if v != 0 {
				allZero = false
			}
		}
		if !allZero {
			return stateData.States[maxIndex].Name
		}
	}
	return stateData.States[rand.Intn(1000)%len(stateData.States)].Name
}

// type State interface {
// 	GetSideEffects() string
// 	Transition() State
// 	State() string
// }

type State struct {
	Name          string
	SideEffects   []string
	TransitionSet []int
}

func (s *State) GetSideEffects() string {
	side := ""
	for i := 0; i < obervablesPerCycle; i++ {
		side += s.SideEffects[rand.Intn(1000)%lengthOfStateSlices]
	}
	return side
}

func (s *State) Transition() State {
	// Determine which state to move to based on transition probability
	return stateData.States[s.TransitionSet[rand.Intn(1000)%lengthOfStateSlices]]
}

func (s *State) State() string {
	return s.Name
}

func NewState(name string, numStates int, obsSet []string) State {
	state := State{
		Name:          name,
		SideEffects:   make([]string, lengthOfStateSlices),
		TransitionSet: make([]int, lengthOfStateSlices),
	}
	lengthObs := len(obsSet)
	rand.Seed(rand.Int63())
	for i, _ := range state.SideEffects {
		if heavyUP < len(obsSet) && i < 2 {
			state.SideEffects[i] = obsSet[heavyUP]
		} else {
			state.SideEffects[i] = obsSet[rand.Intn(rand.Int())%lengthObs]
		}
		state.TransitionSet[i] = rand.Intn(rand.Int()) % numStates

	}
	// sort.Strings(state.SideEffects)
	// sort.Ints(state.TransitionSet)
	return state
}
