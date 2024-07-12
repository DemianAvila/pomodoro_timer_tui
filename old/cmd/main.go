package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	TimerPort "github.com/DemianAvila/pomodoro_timer_tui/entities/timer/port"
	UserInterfacePort "github.com/DemianAvila/pomodoro_timer_tui/entities/userInterface/port"
	EventListeners "github.com/DemianAvila/pomodoro_timer_tui/events"
)

func main() {

	p := tea.NewProgram(UserInterfacePort.InitialModel())
	p.Run()
}

func testCountTo10() {
	timer := TimerPort.Timer{
		Milliseconds: 0,
		Paused:       true,
	}

	TimerPort.Play(&timer)

	finish_clock := false
	var prev_clock, next_clock string

	for !finish_clock {
		next_clock = TimerPort.GetTimerString(&timer)
		if timer.Milliseconds >= 10000 {
			finish_clock = true
		}

		if EventListeners.ChangeSeconds(prev_clock, next_clock) {
			fmt.Println(next_clock)
		}

		prev_clock = next_clock
	}
}
