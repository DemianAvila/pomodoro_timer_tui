package pomodoro

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gen2brain/beeep"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case Tick:
		var secPercentage float64 = 1 / float64(m.Clock.SecondLasting)
		cmd := m.Progress.IncrPercent(secPercentage)
		if m.Clock.HasEnded {
			beeep.Beep(beeep.DefaultFreq, 2000)
			beeep.Notify("Pomodoro Timer -- "+m.CurrentCycle.Name+" has ended", "", "")
		}
		return m, tea.Batch(m.GetTick, cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

	case tea.WindowSizeMsg:
		m.ScreenSize.W = msg.Width
		m.ScreenSize.H = msg.Height

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
				cmd := m.Progress.DecrPercent(1)
				return m, tea.Batch(m.GetTick, cmd)

			}
			return m, nil
		case "s":
			m.Clock = Clock{}
			m.NextCycle()
			cmd := m.Progress.DecrPercent(1)
			return m, tea.Batch(m.GetTick, cmd)

		}
	}
	return m, nil
}
