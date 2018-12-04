package main

import (
	"math/rand"
)

type Bot struct {
	PointCount int                       `json:"count"`
	Records    map[string]map[string]int `json:"records"`
}

func NewBot() *Bot {
	return &Bot{
		Records: map[string]map[string]int{},
	}
}

func NewIntMap() map[string]int {
	m := map[string]int{}
	for _, v := range possibleStates {
		m[v.Name] = 0
	}
	return m
}

func (b *Bot) NewRecord(sideEffects string, state string) {
	record, ok := b.Records[sideEffects]
	if !ok {
		record = NewIntMap()
	}
	temp := record[state]
	temp++
	record[state] = temp
	b.Records[sideEffects] = record
	b.PointCount++
}

func (b *Bot) GuessState(sideEffects string) string {
	if _, ok := b.Records[sideEffects]; ok {
		maxVal := -1
		maxKey := ""
		allZero := true
		for i, v := range b.Records[sideEffects] {
			if v > maxVal {
				maxVal = v
				maxKey = i
			}
			if v != 0 {
				allZero = false
			}
		}
		if !allZero {
			return maxKey
		}
	}
	return possibleStates[rand.Intn(1000)%len(possibleStates)].Name
}

// type State interface {
// 	GetSideEffects() string
// 	Transition() State
// 	State() string
// }

type State struct {
	Name        string
	SideEffects []string
}

func (s *State) GetSideEffects() string {
	side := ""
	for i := 0; i < obervablesPerCycle; i++ {
		side += s.SideEffects[rand.Intn(1000)%10]
	}
	return side
}

func (s *State) Transition() State {
	// Determine which state to move to based on transition probability
	return possibleStates[rand.Intn(1000)%len(possibleStates)]
}

func (s *State) State() string {
	return s.Name
}

func NewState() State {
	return State{
		Name:        "sun",
		SideEffects: []string{"c", "w", "w", "w", "w", "w", "w", "s", "s", "s"}, // 1 c, 6 w, 3 s
	}
}
