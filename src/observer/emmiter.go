package observer

type Emmiter interface {
	RegistObs(o Observer)
	DeregistObs(o Observer)
	Notify()
}

type TickEmmiter struct {
	TickObservers []TickObserver
}

func (te *TickEmmiter) ObsExists(to TickObserver) bool {
	for _, observers := range te.TickObservers {
		if observers.ID == to.ID {
			return true
		}
	}
	return false
}

func (te *TickEmmiter) RegistObs(ID string) TickObserver {
	var new_obs TickObserver = TickObserver{
		ID:            ID,
		PendingUpdate: false,
	}

	if !te.ObsExists(new_obs) {
		te.TickObservers = append(te.TickObservers, new_obs)
		return new_obs
	}

	return TickObserver{}
}

func (te *TickEmmiter) DeregistObs(ID string) string {
	for index, observer := range te.TickObservers {
		if observer.ID == ID {
			te.TickObservers = append(te.TickObservers[:index], te.TickObservers[index+1:]...)
			return observer.ID
		}
	}
	return ""
}

func (te *TickEmmiter) GetObserver(ID string) TickObserver {
	for _, observer := range te.TickObservers {
		if observer.ID == ID {
			return observer
		}
	}
	return TickObserver{}
}

func (te *TickEmmiter) Notify() {
	for index := range te.TickObservers {
		te.TickObservers[index].Update()
	}
}
