package pomodoro

type Observer interface {
	Update()
}

type TickObserver struct {
	ID            string
	PendingUpdate bool
}

type Tick struct {
	Tick bool
}

func (to *TickObserver) GetObserverID() string {
	return to.ID
}

func (to *TickObserver) IsUpdatePending() bool {
	return to.PendingUpdate
}

func (to *TickObserver) SetPendingUp() {
	to.PendingUpdate = true
}

func (to *TickObserver) ResetPendingUp() {
	to.PendingUpdate = false
}

func (to *TickObserver) Update() {
	go func() {
		GlobalTickChanel <- Tick{Tick: true}
	}()
}
