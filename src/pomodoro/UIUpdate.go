package pomodoro

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case Tick:
		return m, m.GetTick
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ":
			if !m.Clock.HasEnded {
				if m.Clock.IsRuning {
					go m.Clock.Pause()
				} else {
					go m.Clock.Play()
				}
			}
			return m, nil
		}
	}
	return m, nil
}
