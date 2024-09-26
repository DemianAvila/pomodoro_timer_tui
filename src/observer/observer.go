package observer

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Observer interface {
	Update()
}

type TickObserver struct {
	ID            string
	PendingUpdate bool
}

func (to *TickObserver) GetObserverID() string {
	return to.ID
}

func (to *TickObserver) IsUpdatePending() bool {
	return to.PendingUpdate
}

func (to *TickObserver) SetPendingUp() {
	to.PendingUpdate = true
}

func (to *TickObserver) ResetPendingUp() {
	to.PendingUpdate = false
}

type Tick struct {
	tick bool
}

func (to *TickObserver) Update() tea.Msg {
	return Tick{
		tick: true,
	}
}
