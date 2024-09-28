package pomodoro

import (
	tea "github.com/charmbracelet/bubbletea"
)

var GlobalTickChanel chan tea.Msg = make(chan tea.Msg)
