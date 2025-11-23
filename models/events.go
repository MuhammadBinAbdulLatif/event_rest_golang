package models

import (
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int 
	Name        string `binding:"required"`
	Description string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	Location    string `binding:"required"`
	UserId int64
}

func (e Event) Save() error {
	query := `insert into events (name, description, location, datetime, user_id) values ($1, $2, $3, $4, $5)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	return nil
}


func GetAllEvents() ([]Event, error) {
	query := `select * from events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}


func GetEventByID(id string) (Event, error) {
	query := `select * from events where id = `+id
	row := db.DB.QueryRow(query)
	event := Event{}
	row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	return event, nil
}


func (event Event) UpdateEvent() error {
	query := `
	Update events
	set name=$2, description = $3, location = $4, dateTime = $5
	where id=$1`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return  err
	}
	defer stmt.Close()
	_,err =stmt.Exec(event.ID, event.Name, event.Description, event.Location, event.DateTime)
	// return an error
return  err
}


func (event Event) Delete() error {
	query:= "Delete from events where id =$1"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return  err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return  err
}


func (event Event) Register(userId int64) error {
	query := `insert into registrations (event_id, user_id) values($1, $2)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, userId)
	return err
}


func (event Event) Unregister(userId int64) error{
	query := `delete from registrations where event_id = $1 and user_id = $2`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, userId)
	return err
}