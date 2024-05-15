package UserInterfacePort

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(22)

	return style.Render("Hello, kitty")
}
