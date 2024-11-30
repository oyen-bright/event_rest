package models

import (
	"fmt"
	"log"
	"time"

	"github.com/oyen-bright/event_REST/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (event *Event) Save() error {

	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?,?,?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	id, _ := result.LastInsertId()
	event.ID = id
	return err

}

func (event *Event) Upate() error {

	query := `
	UPDATE events SET name = ?, description = ?, location = ?, dateTime = ? WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		log.Println("Error updating event", err)
		return err
	}
	id, err := result.LastInsertId()
	event.ID = id

	fmt.Println("Event saved successfully", result)

	return err

}

func (event *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}
func GetAllEvents() ([]Event, error) {

	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEvent(ID int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(ID)

	var event Event

	err = row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) Register(userID int64) error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, userID)
	return err
}

func (event *Event) Unregister(userID int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, userID)
	return err
}
