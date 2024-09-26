package userUI

import (
	"pomodoro_timer_tui/src/observer"
	"pomodoro_timer_tui/src/pomodoro"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Clock             pomodoro.Clock
	Cycles            []pomodoro.Cycle
	CurrentCycle      pomodoro.Cycle
	ModelTickObserver observer.TickObserver
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

func (m *Model) ChangeCycle(c pomodoro.Cycle) {
	m.CurrentCycle = c
}

func (m *Model) GetWorkCycle() pomodoro.Cycle {
	return m.Cycles[0]
}

func (m *Model) GetShortRestCycle() pomodoro.Cycle {
	return m.Cycles[1]
}

func (m *Model) GetLongRestCycle() pomodoro.Cycle {
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

func (m Model) Init() tea.Cmd {
	return nil
}

func ModelInitialization() *Model {

	var initialClock pomodoro.Clock = pomodoro.Clock{
		SecondLasting: pomodoro.AvailableCycles[0].LastingInSeconds,
		CurrentSecond: 0,
		IsRuning:      false,
		HasEnded:      false,
		TickEmmiter:   observer.TickEmmiter{},
	}

	initialClock.TickEmmiter.RegistObs("updateTick")

	return &Model{
		Clock:        initialClock,
		Cycles:       pomodoro.AvailableCycles,
		CurrentCycle: pomodoro.AvailableCycles[0],
	}
}
