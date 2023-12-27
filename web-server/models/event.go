package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/learning-webserver/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
	UserId      int64     `json:"user_id"`
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

func (e *Event) GetEvent(id int64) (*Event, error) {

	query := "SELECT * FROM events where id = ?"

	row := db.DB.QueryRow(query, id)

	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.Created_At, &e.Updated_At, &e.UserId)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Event) UpdateEvent(id int64, userId int64) (int64, error) {
	query := "UPDATE events SET name = ?, description = ?, location = ?, user_id = ? where id = ? and user_id = ?"

	stmt, err := db.DB.Prepare(query)

	// err = row.Scan()

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(&e.Name, &e.Description, &e.Location, &e.UserId, id)

	if err != nil {
		log.Fatal(err)
	}

	id, err = result.RowsAffected()

	if id == 0 {
		return 0, errors.New("You are not authorized to delete this event")
	}

	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events;"

	rows, err := db.DB.Query(query)

	fmt.Println(rows)

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

		// eve := map[string]interface{}{
		// 	"id":          event.ID,
		// 	"name":        event.Name,
		// 	"description": event.Description,
		// 	"location":    event.Location,
		// 	"createdat":   event.Created_At,
		// 	"updatedat":   event.Updated_At,
		// 	"user_id":     event.UserId,
		// }

		// newEvent := mapToEvent(eve)

		// event.Created_At, _ = time.Parse("2006-01-02 15:04:05", event.Created_At.Format("2006-01-02 15:04:05"))
		// event.Updated_At, _ = time.Parse("2006-01-02 15:04:05", event.Updated_At.Format("2006-01-02 15:04:05"))

		events = append(events, event)
	}

	return events, nil
}

// func mapToEvent(m map[string]interface{}) Event {
// 	return Event{
// 		ID:          m["id"].(int64),
// 		Name:        m["name"].(string),
// 		Description: m["description"].(string),
// 		Location:    m["location"].(string),
// 		Created_At:  m["createdat"].(time.Time),
// 		Updated_At:  m["updatedat"].(time.Time),
// 		UserId:      m["user_id"].(int64),
// 	}
// }

func (e *Event) DeleteEvent(id int64, userId int64) (int64, error) {
	query := "DELETE FROM events where id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(id, userId)

	if err != nil {
		log.Fatal(err)
	}

	id, err = result.RowsAffected()

	if id == 0 {
		return 0, errors.New("You are not authorized to delete this event")
	}

	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}
