package keyboard

import (
	"log"

	"github.com/MarinX/keylogger"
)

// List the keyboards on the computer
func List() []string {
	return keylogger.FindAllKeyboardDevices()
}

// Keyboard represents a keyboard with its associated informations
type Keyboard struct {
	Channel  chan keylogger.InputEvent
	Location string
	Buffer   []string
	Index    int
}

// Read strokes from the keyboard
func Read(location string) (Keyboard, bool) {
	kb, err := keylogger.New(location)
	if err != nil {
		log.Printf("couldn't open keyboard '%s': %v\n", location, err)
		return Keyboard{}, false
	}
	return Keyboard{
		Channel:  kb.Read(),
		Location: location,
	}, true
}
