package server

import "time"

// Key is a struct for holding the data of each key
type Key struct {
	id          uint
	name        string
	description string
	created     time.Time
	expires     time.Time
	maxUses     uint
	uses        uint
	lastUsed    time.Time
}
