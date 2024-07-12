package EventListeners

func ChangeSeconds(prevClock string, nextClock string) bool {
	return prevClock != nextClock
}
