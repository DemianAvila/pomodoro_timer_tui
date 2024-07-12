package timer

type Timer struct {
	SecondsPassed 	uint64
	TimeLimit 		uint64
	Paused			bool
}

func (t *Timer) GetSecondsPassed() uint64{
	return t.SecondsPassed	
}

func (t *Timer) GetTimeLimit() uint64{
	return t.TimeLimit
}

func (t *Timer) IsPaused() bool {
	return t.Paused
}

func (t *Timer) Play() {
	t.Paused = false
}