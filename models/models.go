package models

import (
	"time"

	"github.com/hrhridoy/event-booking-API/db"
)

type Event struct {
	ID          int64
	Name        string `bindings:"required"`
	Description string `bindings:"required"`
	Location    string `bindings:"required"`
	DateTime    time.Time
	UserID      int
}

// All the events that are created based on the Events struct are stored in this variable.
// this variable Events is type of Slice of Events structs
var events = []Event{}

func (e *Event) Save() error {
	// Events will be stored on the DB Later.
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}
func GetAllEvents() ([]Event, error) {
	// Returns all the events stored in the Events variable for now.
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
