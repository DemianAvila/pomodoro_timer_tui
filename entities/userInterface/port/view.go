package UserInterfacePort

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
	TimerPort "github.com/DemianAvila/pomodoro_timer_tui/entities/timer/port"
)


func (m Model) View() string {
	doc := strings.Builder{}
	if TimerPort.IsPaused(&m.Timer) {
		dialogBoxStyle := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#874BFD")).
				Padding(2, 2).
				BorderTop(true).
				BorderLeft(true).
				BorderRight(true).
				BorderBottom(true)
				
		question := lipgloss.NewStyle().
					Width(50).
					Align(lipgloss.Center).
					Render(TimerPort.GetTimerString(&m.Timer))
		ui := lipgloss.JoinVertical(lipgloss.Center, question)

		dialog := lipgloss.Place(2, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(" "),
		)

		doc.WriteString(dialog + "\n\n")
		

		return doc.String()
	}
	/*for !TimerPort.IsPaused(m.Timer){
		if 
		doc := strings.Builder{}	
		dialogBoxStyle := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#874BFD")).
				Padding(2, 2).
				BorderTop(true).
				BorderLeft(true).
				BorderRight(true).
				BorderBottom(true)
				
		question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Are you sure you want to eat marmalade?")
		ui := lipgloss.JoinVertical(lipgloss.Center, question)

		dialog := lipgloss.Place(2, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(" "),
		)

		doc.WriteString(dialog + "\n\n")
		

		return doc.String()
	}*/
	
	return ""
	
}

