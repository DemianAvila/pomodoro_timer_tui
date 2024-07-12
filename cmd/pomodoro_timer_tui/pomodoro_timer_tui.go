package main

import (
	"fmt"
	"github.com/DemianAvila/pomodoro_timer_tui/internal/timer"
)

func main() {
	var timer timer.Timer
	
	fmt.Println(timer.GetSecondsPassed())
	fmt.Println(timer.TimeLimit())
	fmt.Println(timer.IsPaused())
	timer.Play()
	fmt.Println(timer.IsPaused())

}