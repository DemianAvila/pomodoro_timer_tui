package pomodoro

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Clock        Clock
	Cycles       []Cycle
	CurrentCycle Cycle
}

func (m *Model) IsWorkCycle() bool {
	return m.CurrentCycle.Name == "WorkCycle"
}

func (m *Model) IsShortRest() bool {
	return m.CurrentCycle.Name == "ShortRestCycle"
}

func (m *Model) IsLongRest() bool {
	return m.CurrentCycle.Name == "LongRestCycle"
}

func (m *Model) ChangeCycle(c Cycle) {
	m.CurrentCycle = c
}

func (m *Model) GetWorkCycle() Cycle {
	return m.Cycles[0]
}

func (m *Model) GetShortRestCycle() Cycle {
	return m.Cycles[1]
}

func (m *Model) GetLongRestCycle() Cycle {
	return m.Cycles[2]
}

func (m *Model) IsLastWorkCycle() bool {
	return m.GetWorkCycle().Counter == 4
}

func (m *Model) NextCycle() {
	if m.IsLongRest() || m.IsShortRest() {
		m.ChangeCycle(m.GetWorkCycle())
	} else {
		if m.IsLastWorkCycle() {
			m.ChangeCycle(m.GetLongRestCycle())
		} else {
			m.ChangeCycle(m.GetShortRestCycle())
		}
	}
}

func (m *Model) GetTick() tea.Msg {
	getTick := <-GlobalTickChanel
	return getTick

}

func (m Model) Init() tea.Cmd {
	return m.GetTick
}

func ModelInitialization() *Model {

	var initialClock Clock = Clock{
		SecondLasting: AvailableCycles[0].LastingInSeconds,
		CurrentSecond: 0,
		IsRuning:      false,
		HasEnded:      false,
	}

	initialClock.TickEmmiter.RegistObs("updateTick")

	return &Model{
		Clock:        initialClock,
		Cycles:       AvailableCycles,
		CurrentCycle: AvailableCycles[0],
	}
}
