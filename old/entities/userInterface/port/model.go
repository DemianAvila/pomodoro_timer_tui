package UserInterfacePort

import (
	TimerPort "github.com/DemianAvila/pomodoro_timer_tui/entities/timer/port"
	PomodoroPort "github.com/DemianAvila/pomodoro_timer_tui/entities/pomodoro/port"
)

type Model struct {
	Timer TimerPort.Timer
	Pomodoro PomodoroPort.Pomodoro
	ClockString string
}



func InitialModel() Model {

	var stages = map[string] PomodoroPort.PomodoroStage{}
	stages["workCycle"] = PomodoroPort.PomodoroStage {
		Name: "Work Cycle",
		DurationMilli: 5000,
		CyclesDone: 0,
	}
	stages["shortRestCycle"] = PomodoroPort.PomodoroStage {
		Name: "Short Rest Cycle",
		DurationMilli: 300000,
		CyclesDone: 0,
	}
	stages["longRestCycle"] = PomodoroPort.PomodoroStage {
		Name: "Long Rest Cycle",
		DurationMilli: 900000,
		CyclesDone: 0,
	}

	return Model{
		Timer: TimerPort.Timer{
			Milliseconds: 0,
			Paused:       true,
		},
		Pomodoro: PomodoroPort.Pomodoro{
			Stages: stages,
			WorkCyclesUntilLongRest: 4,
			CurrentStage: stages["workCycle"],
		},
	}
}
