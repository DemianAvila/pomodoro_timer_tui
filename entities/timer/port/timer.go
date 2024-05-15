package TimerPort

import (
	"fmt"
	"strconv"
	"time"
)

type Timer struct {
	Milliseconds uint64
	Paused       bool
}

type TimerPresentation struct {
	Hours   int64
	Minutes int64
	Seconds int64
}

func ResetTimer(c *Timer) {
	c.Paused = true
	c.Milliseconds = 0
}

func Pause(c *Timer) {
	c.Paused = true
}

func Play(c *Timer) {
	c.Paused = false
	go CountingController(c)

}

func IsPaused(c *Timer) bool {
	return c.Paused
}

func AddMillisecond(c *Timer) {
	c.Milliseconds += 1
}

func CountingController(c *Timer) {
	for !IsPaused(c) {
		time.Sleep(1 * time.Millisecond)
		AddMillisecond(c)
	}
}

func TimerFormat(c *Timer) TimerPresentation {
	var Hours, Minutes, Seconds uint64
	dur := time.Duration(c.Milliseconds) * time.Millisecond

	Hours = uint64(time.Duration.Hours(dur))
	Minutes = uint64(time.Duration.Minutes(dur)) - (Hours * 60)
	Seconds = uint64(time.Duration.Seconds(dur)) - (Hours * 3600) - (Minutes * 60)

	return TimerPresentation{
		Hours:   int64(Hours),
		Minutes: int64(Minutes),
		Seconds: int64(Seconds),
	}
}

func clockFormat(ammount string) string {
	if len(ammount) == 1 {
		return fmt.Sprintf("0%s", ammount)
	} else {
		return ammount
	}
}

func GetTimerString(c *Timer) string {
	f := TimerFormat(c)
	hours := strconv.FormatInt(f.Hours, 10)
	minutes := strconv.FormatInt(f.Minutes, 10)
	seconds := strconv.FormatInt(f.Seconds, 10)
	return fmt.Sprintf(
		"%s:%s:%s",
		clockFormat(hours),
		clockFormat(minutes),
		clockFormat(seconds),
	)
}
