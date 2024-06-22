package server

import (
	"bufio"
	"fmt"
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
	uses        uint64
	lastUsed    string
}

func (k *Key) getString() string {
	return fmt.Sprintf("ID: %s\nName: %s\nDescription: %s\nCreated: %s\nExpires: %s\nSingle Use: %v\nUses: %v\nLast Used: %s",
		k.id.String(), k.name, k.description, k.created, k.expires, k.singleUse, k.uses, k.lastUsed)
}

func createKey(n string, d string, exp string, single bool) (Key, error) {
	// Auto-generate a unique ID
	id, err := uuid.NewRandom()
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
		"never",
	}

	// Return nil if no error has been encountered
	return key, nil
}

func (s *Server) readKeyFromFile(id string) (Key, error) {
	s.Log("Attempting to read key: " + id)
	f, err := os.OpenFile(s.keyDir+id, os.O_RDONLY, 0640)
	if err != nil {
		return Key{}, err
	}

	defer f.Close()

	// Start reading strings separated by newlines
	reader := bufio.NewReader(f)
	readstr, err := reader.ReadString('\n')
	if err != nil {
		return Key{}, err
	}
	kid, err := uuid.Parse(readstr) // Parse the string into a UUID
	if err != nil {
		return Key{}, err
	}

	n, err := reader.ReadString('\n') // Read the key name
	if err != nil {
		return Key{}, err
	}

	d, err := reader.ReadString('\n') // Read the key description
	if err != nil {
		return Key{}, err
	}

	created, err := reader.ReadString('\n') // Read the date key was created
	if err != nil {
		return Key{}, err
	}

	exp, err := reader.ReadString('\n') // Read the expiry date
	if err != nil {
		return Key{}, err
	}

	singleStr, err := reader.ReadString('\n') // Read single use
	if err != nil {
		return Key{}, err
	}
	single, err := strconv.ParseBool(singleStr) // Parse to bool
	if err != nil {
		return Key{}, err
	}

	usesStr, err := reader.ReadString('\n') // Read number of uses
	if err != nil {
		return Key{}, err
	}
	uses, err := strconv.ParseUint(usesStr, 10, 64) // Parse to int
	if err != nil {
		return Key{}, err
	}

	last, err := reader.ReadString('\n')
	if err != nil {
		return Key{}, err
	}

	k := Key{
		kid,
		n,
		d,
		created,
		exp,
		single,
		uses,
		last,
	}
	s.Log("Read key: " + k.getString())
	return k, nil
}

func (s *Server) writeKey(k *Key) error {
	s.Log("Attempting to write key:\n" + k.getString())
	id := k.id.String()
	name := k.name
	description := k.description
	created := k.created
	expires := k.expires
	singleUse := strconv.FormatBool(k.singleUse)
	uses := strconv.FormatUint(k.uses, 10)
	lastUsed := k.lastUsed

	// Create the file for working
	f, err := os.OpenFile(s.keyDir+id, os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		return err
	}

	defer f.Close()

	// Parse everything into a string
	parsed := id + "\n" + name + "\n" + description + "\n" + created + "\n" + expires + "\n" + singleUse + "\n" +
		uses + "\n" + lastUsed + "\n"

	// Write to the file
	n, err := f.WriteString(parsed)
	if err != nil {
		return err
	}
	s.Log(strconv.Itoa(n) + " bytes written.")

	// Return nil if no error has been encountered
	return nil
}
