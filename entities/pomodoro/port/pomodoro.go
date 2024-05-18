package PomodoroPort

import (
		TimerPort "github.com/DemianAvila/pomodoro_timer_tui/entities/timer/port"
)

type PomodoroStage struct {
	Name string
	DurationMilli uint64
	CyclesDone int
}

type Pomodoro struct {
	Stages map[string] PomodoroStage
	WorkCyclesUntilLongRest int 
	CurrentStage PomodoroStage
}

func PomodoroNextCycle (p *Pomodoro){
	//IF CURRENT CYCLE IS A WORK ONE
	if p.CurrentStage.Name =="Work Cycle"{
		//IF THE CURRENT CYCLE EQUALS THE EXPECTED FOR LONG REST, PASS TO LONG REST
		if p.CurrentStage.CyclesDone == (p.WorkCyclesUntilLongRest-1){
			p.CurrentStage = p.Stages["longRestCycle"]
			AddCycleCounter (p.Stages["longRestCycle"])
		} else {
			p.CurrentStage = p.Stages["shortRestCycle"]
			AddCycleCounter (p.Stages["shortRestCycle"])
		}
	}
	//IF CURRENT CYCLE IS A REST ONE; AUTOMATICALLY JUMP TO A WORK ONE
	if p.CurrentStage.Name == "Short Rest Cycle" || p.CurrentStage.Name == "Long Rest Cycle" {
		p.CurrentStage = p.Stages["workCycle"]
		if p.CurrentStage.Name == "Long Rest Cycle" {
			ResetCycles (p)
		}
		AddCycleCounter (p.Stages["workCycle"])
	}
}

func ResetCycles (p *Pomodoro){
	for _, v := range p.Stages {
        v.CyclesDone = 0
    }
}

func AddCycleCounter (ps PomodoroStage){
	ps.CyclesDone++
}

func MillisecondsTillFinishCycle(p *Pomodoro, t *TimerPort.Timer) uint64{
	return p.CurrentStage.DurationMilli - t.Milliseconds
}