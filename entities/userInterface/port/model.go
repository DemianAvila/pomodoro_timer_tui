package UserInterfacePort

import (
	TimerPort "github.com/DemianAvila/pomodoro_timer_tui/entities/timer/port"
)

type Model struct {
	Timer TimerPort.Timer
}

func InitialModel() Model {
	return Model{
		Timer: TimerPort.Timer{
			Milliseconds: 0,
			Paused:       true,
		},
	}
}
