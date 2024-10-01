package pomodoro

import (	"github.com/charmbracelet/bubbles/progress"
)

func (m Model) View() string {

	/*var sec string = strconv.FormatUint(m.Clock.CurrentSecond, 10)
	var paused string = strconv.FormatBool(m.Clock.IsRuning)

	return sec + "   " + paused*/
	//return strconv.FormatUint(m.Clock.CurrentSecond, 10)
	var reminingSec uint64 = m.Clock.SecondLasting - m.Clock.CurrentSecond
	return StringTimeToASCII(SecToTimeFormat(reminingSec)) + "\n" +
	
}
