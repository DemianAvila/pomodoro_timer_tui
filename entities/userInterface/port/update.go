package UserInterfacePort

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	TimerPort "github.com/DemianAvila/pomodoro_timer_tui/entities/timer/port"
	PomodoroPort "github.com/DemianAvila/pomodoro_timer_tui/entities/pomodoro/port"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ":
			ClockInterfaceHandeler(&m)
		}
			return m, nil
	}

	return m, nil
}

func ClockInterfaceHandeler (m *Model) int {
	//IF THE CLOCK IS PAUSED AND MILISECONDS ARE NOT 0, PLAY
	if TimerPort.IsPaused(&m.Timer) && 
	PomodoroPort.MillisecondsTillFinishCycle(&m.Pomodoro, &m.Timer) > 0 {
		TimerPort.Play(&m.Timer)
		fmt.Println("play timer")
		return 0
	}
	//IF THE CLOCK IS PAUSED AND MILISECONDS ARE 0, CHANGE CYCLE
	if TimerPort.IsPaused(&m.Timer) && 
	PomodoroPort.MillisecondsTillFinishCycle(&m.Pomodoro, &m.Timer) == 0 {
		PomodoroPort.PomodoroNextCycle (&m.Pomodoro)
		fmt.Println("change cycle")
		return 0
	}
	//IF THE CLOCK IS PLAY, PAUSE
	if !TimerPort.IsPaused(&m.Timer) {
		TimerPort.Pause(&m.Timer)
		fmt.Println("pause")
		return 0
	}
	return 0
}