package main

import (
	"math/rand"
)

type Bot struct {
	PointCount int                    `json:"count"`
	Records    map[string]StateRecord `json:"records"`
}

func (b *Bot) NewRecord(sideEffects string, state string) {
	record, ok := b.Records[sideEffects]
	if !ok {
		record = StateRecord{}
	}
	if state == "sun" {
		record.Sun++
	} else {
		record.Rain++
	}
	b.Records[sideEffects] = record
	b.PointCount++
}

func (b *Bot) GuessState(sideEffects string) string {
	if _, ok := b.Records[sideEffects]; ok {
		if b.Records[sideEffects].Sun > b.Records[sideEffects].Rain {
			return "sun"
		}
		return "rain"
	}

	if rand.Intn(100)%2 == 0 {
		return "sun"
	}
	return "rain"
}

type StateRecord struct {
	Sun  int
	Rain int
}

type State interface {
	GetSideEffects() string
	Transition() State
	State() string
}

type RainState struct {
	Name        string
	SideEffects []string
}

type SunState struct {
	Name        string
	SideEffects []string
}

func (s *RainState) GetSideEffects() string {
	side := ""
	for i := 0; i < obervablesPerCycle; i++ {
		side += s.SideEffects[rand.Intn(1000)%10]
	}
	return side
}

func (s *SunState) GetSideEffects() string {
	side := ""
	for i := 0; i < obervablesPerCycle; i++ {
		side += s.SideEffects[rand.Intn(1000)%10]
	}
	return side
}

func (s *RainState) Transition() State {
	val := rand.Intn(1000) % 10
	if val < 4 {
		return NewRainState()
	}
	return NewSunState()
}

func (s *SunState) Transition() State {
	val := rand.Intn(1000) % 10
	if val < 6 {
		return NewRainState()
	}
	return NewRainState()
}

func (s *RainState) State() string {
	return "rain"
}

func (s *SunState) State() string {
	return "sun"
}

func NewRainState() State {
	return &RainState{
		Name:        "rain",
		SideEffects: []string{"w", "c", "c", "c", "c", "s", "s", "s", "s", "s"}, // 1 w, 4 c, 5 s
	}
}

func NewSunState() State {
	return &SunState{
		Name:        "sun",
		SideEffects: []string{"c", "w", "w", "w", "w", "w", "w", "s", "s", "s"}, // 1 c, 6 w, 3 s
	}
}
