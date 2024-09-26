package userUI

import (
	"strconv"
)

func (m Model) View() string {

	var sec string = strconv.FormatUint(m.Clock.CurrentSecond, 10)
	var paused string = strconv.FormatBool(m.Clock.IsRuning)

	return sec + "   " + paused

}
