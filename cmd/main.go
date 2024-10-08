package main

import (
	"fmt"
	"os"
	"pomodoro_timer_tui/src/pomodoro"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	p := tea.NewProgram(pomodoro.ModelInitialization())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
