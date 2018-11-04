package lock

// Unlock the door and keep unlocked until locked again
func Unlock() {
	// TO DO
}

// Lock the door
func Lock() {
	// TO DO
}

// UnlockTemp unlocks the door for a period of time (in seconds) and then locks it again
func UnlockTemp(duration int) {
	Unlock()
	// TO DO: WAIT for <duration>
	Lock()
}
