package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	conn, err := sql.Open("postgres", "postgres://postgres:<yourpostgresPassword>@localhost:5432/test_project?sslmode=disable")
	if err != nil {
		panic(err)
	}
	DB = conn
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)
}



func CreateTables() {
	createEventsTable :=`CREATE TABLE IF NOT EXISTS events (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime TIMESTAMP NOT NULL,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
`
	

createUserTable := `
create table if not exists users (
id serial primary key,
email Text not null unique,
password Text not null
)`

createRegistrationsTable := `
create table if not exists registrations (
id serial primary key,
event_id integer , 
user_id integer,
foreign key(event_id) references events(id),
foreign key(user_id) references users(id)
)`
_, err := DB.Exec(createUserTable)
if err != nil{
	print(err.Error())
	panic("Could not create users talbe")
}
_, err =DB.Exec(createEventsTable)
	if err != nil{
		panic("Could not create events table")

	}
_, err =DB.Exec(createRegistrationsTable)
	if err != nil{
		panic("Could not create registrations table")

	}
}
