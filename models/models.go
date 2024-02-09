package models

import "time"

type Event struct {
	ID          int
	Name        string `bindings:"required"`
	Description string `bindings:"required"`
	Location    string `bindings:"required"`
	DateTime    time.Time
	UserID      int
}

// All the events that are created based on the Events struct are stored in this variable.
// this variable Events is type of Slice of Events structs
var events = []Event{}

func (e *Event) Save() {
	// Events will be stored on the DB Later.
	events = append(events, *e)
}
func GetAllEvents() []Event {
	// Returns all the events stored in the Events variable for now.
	return events
}
