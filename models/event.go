package models

import (
	"time"

	"github.com/Jane-Mwangi/GoEventApi/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    ` binding:"required"`
	Location    string    ` binding:"required"`
	DateTime    time.Time ` binding:"required"`
	UserID      int64     `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := "INSERT INTO events(name, description, location, dateTime, user_id) VALUES(?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	defer stmt.Close()
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
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

func GetEvent(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?,description = ?,location = ?,dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err

}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}


func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}


	defer stmt.Close()

	stmt.Exec(e.ID, userId)
	return err
}