package pomodoro

import (
	"time"
)

type Clock struct {
	SecondLasting uint64
	CurrentSecond uint64
	IsRuning      bool
	HasEnded      bool
	TickEmmiter   TickEmmiter
}

func (clk *Clock) Play() {
	if !clk.IsRuning {
		clk.IsRuning = true
		var second uint64
	out:
		for second = clk.CurrentSecond; second <= clk.SecondLasting; second++ {
			clk.CurrentSecond = second
			time.Sleep(time.Second / 10000)
			clk.TickEmmiter.Notify()
			if !clk.IsRuning {
				break out
			}
		}
	}
	if clk.CurrentSecond == clk.SecondLasting {
		clk.HasEnded = true
		clk.IsRuning = false
	}
}

func (clk *Clock) Pause() {
	if clk.IsRuning {
		clk.IsRuning = false
	}
}
