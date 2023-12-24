package models

import (
	"fmt"
	"log"
	"time"

	"github.com/learning-webserver/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Created_At  time.Time
	Updated_At  time.Time
	UserId      int `binding:"required"`
}

// var events []Event = []Event{}

func (e *Event) Save() (int64, error) {
	query := `
		INSERT INTO events(name, description, location, user_id) VALUES (?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.UserId)

	if err != nil {
		return 0, err
	}

	// events = append(events, *e)

	id, err := result.LastInsertId()
	fmt.Println("ID", id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events;"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Created_At, &event.Updated_At, &event.UserId)

		if err != nil {
			return nil, err
		}

		event.Created_At, _ = time.Parse("2006-01-02 15:04:05", event.Created_At.Format("2006-01-02 15:04:05"))
		event.Updated_At, _ = time.Parse("2006-01-02 15:04:05", event.Updated_At.Format("2006-01-02 15:04:05"))

		events = append(events, event)
	}

	return events, nil
}
