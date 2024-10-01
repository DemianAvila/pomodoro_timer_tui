package pomodoro

type Cycle struct {
	Name             string
	LastingInSeconds uint64
	Counter          uint16
}

var AvailableCycles []Cycle

func init() {

	AvailableCycles = []Cycle{
		{
			Name:             "WorkCycle",
			LastingInSeconds: 25 * 60 * 60,
			Counter:          0,
		},
		{
			Name:             "ShortRestCycle",
			LastingInSeconds: 5 * 60 * 60,
			Counter:          0,
		},
		{
			Name:             "LongRestCycle",
			LastingInSeconds: 15 * 60 * 60,
			Counter:          0,
		},
	}
}
