package server

import (
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Key is a struct for holding the data of each key
type Key struct {
	id          uuid.UUID
	name        string
	description string
	created     string
	expires     string
	singleUse   bool
	uses        int
	lastUsed    string
}

func createKey(n string, d string, exp string, single bool) (Key, error) {
	// Auto-generate a unique ID
	id, err := uuid.NewUUID()
	if err != nil {
		return Key{}, err
	}
	created := time.Now().Format(time.UnixDate)
	key := Key{
		id,
		n,
		d,
		created,
		exp,
		single,
		0,
		created,
	}

	// Return nil if no error has been encountered
	return key, nil
}

func writeKey(k *Key) error {
	id := k.id.String()
	name := k.name
	description := k.description
	created := k.created
	expires := k.expires
	singleUse := strconv.FormatBool(k.singleUse)
	uses := strconv.Itoa(k.uses)
	lastUsed := k.lastUsed

	// Create the file for working
	f, err := os.OpenFile(id, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Parse everything into a string
	parsed := id + "\n" + name + "\n" + description + "\n" + created + "\n" + expires + "\n" + singleUse + "\n" +
		uses + "\n" + lastUsed

	// Write to the file
	_, err = f.WriteString(parsed)
	if err != nil {
		return err
	}

	// Return nil if no error has been encountered
	return nil
}
