package pomodoro

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case Tick:
		var secPercentage float64 = 1 / float64(m.Clock.SecondLasting)
		cmd := m.Progress.IncrPercent(secPercentage)
		return m, tea.Batch(m.GetTick, cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

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
			} else {
				m.NextCycle()
			}
			return m, nil
		}
	}
	return m, nil
}
