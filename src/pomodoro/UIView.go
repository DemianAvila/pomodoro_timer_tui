package pomodoro

import "github.com/charmbracelet/lipgloss"

func PrintClock(m *Model) string {
	var reminingSec uint64 = m.Clock.SecondLasting - m.Clock.CurrentSecond
	return StringTimeToASCII(SecToTimeFormat(reminingSec)) + "\n"
}

func PrintProgressBar(m *Model) string {
	return m.Progress.View() + "\n\n"
}

func PrintClockAndCycleStatus(m *Model) string {
	var status string
	if m.Clock.HasEnded {
		status = m.CurrentCycle.Name +
			" has ended, press \"enter\" or \"space\"" +
			" to move to the next cycle"
	} else {
		if m.Clock.IsRuning {
			status = "▶️  - Your " +
				m.CurrentCycle.Name +
				" is running."

			timesDots := m.Clock.CurrentSecond % 3
			for i := 0; i < int(timesDots)+1; i++ {
				status = status + " ."
			}
		} else {
			status = "⏸️  - The cycle is paused" +
				" press \"enter\" or \"space\"" +
				" to resume cycle."
		}
	}

	return status
}

func GetRenderedString(styleName string, strToRender string, stylesMap map[string]lipgloss.Style) string {
	style, ok := stylesMap[styleName]
	if ok {
		return style.Render(strToRender)
	} else {
		return ""
	}
}

func (m Model) View() string {

	/*var sec string = strconv.FormatUint(m.Clock.CurrentSecond, 10)
	var paused string = strconv.FormatBool(m.Clock.IsRuning)

	return sec + "   " + paused*/
	//return strconv.FormatUint(m.Clock.CurrentSecond, 10)
	var Styles map[string]lipgloss.Style = map[string]lipgloss.Style{
		"headerStyle": lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 0).
			Width(m.ScreenSize.W).
			Align(lipgloss.Center),
		"workCycleStyle": lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#FF2D00")).
			Padding(0, 0).
			Width(m.ScreenSize.W).
			Align(lipgloss.Center),
		"shortCycleStyle": lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#16FEFE")).
			Padding(0, 0).
			Width(m.ScreenSize.W).
			Align(lipgloss.Center),
		"longCycleStyle": lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#00ED2C")).
			Padding(0, 0).
			Width(m.ScreenSize.W).
			Align(lipgloss.Center),
		"fullScreen": lipgloss.NewStyle().
			Width(m.ScreenSize.W).
			Height(m.ScreenSize.H),
		"statusBox": lipgloss.NewStyle().
			Width(m.ScreenSize.W).
			Align(lipgloss.Center),
	}

	var cycleHeader string
	if m.IsWorkCycle() {
		cycleHeader = GetRenderedString("workCycleStyle", m.CurrentCycle.Name, Styles)
	} else if m.IsShortRest() {
		cycleHeader = GetRenderedString("shortCycleStyle", m.CurrentCycle.Name, Styles)
	} else {
		cycleHeader = GetRenderedString("longCycleStyle", m.CurrentCycle.Name, Styles)
	}

	var returnString string = GetRenderedString("headerStyle", "🍅  -- POMODORO TIMER --  🍅", Styles) +
		"\n" +
		cycleHeader +
		"\n\n" +
		PrintClock(&m) +
		GetRenderedString("statusBox", PrintProgressBar(&m), Styles) +
		"\n" +
		GetRenderedString("statusBox", PrintClockAndCycleStatus(&m), Styles) +
		"\n\n\nPress \"q\" or \"ctrl+c\" to exit program" +
		"\nPress \"s\" to skip cycle\n"

	return GetRenderedString("fullScreen", returnString, Styles)
}
