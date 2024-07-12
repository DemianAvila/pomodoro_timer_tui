package UserInterfacePort

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
	TimerPort "github.com/DemianAvila/pomodoro_timer_tui/entities/timer/port"
	EventListeners "github.com/DemianAvila/pomodoro_timer_tui/events"
)


func (m Model) View() string {
	doc := strings.Builder{}
	finish_clock := false
	var prev_clock, next_clock, question string

	for !finish_clock {
		next_clock = TimerPort.GetTimerString(&m.Timer)
		if m.Timer.Milliseconds >= 10000 {
			finish_clock = true
		}

		if EventListeners.ChangeSeconds(prev_clock, next_clock) {
			question = lipgloss.NewStyle().
				Width(50).
				Align(lipgloss.Center).
				Render(next_clock)
				return ""
		}

		prev_clock = next_clock
		dialogBoxStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(2, 2).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
			
	
		ui := lipgloss.JoinVertical(lipgloss.Center, question)

		dialog := lipgloss.Place(2, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(" "),
		)

		doc.WriteString(dialog + "\n\n")
		

		return doc.String()
	}
	
	return doc.String()
	
	
	
	
}

