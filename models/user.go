package models

import (
	"errors"
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return  err
	}
	defer stmt.Close()
	password, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_,err = stmt.Exec(u.Email, password)
	if err != nil {
		return  err
	}	
	return err

}


func (u *User) ValidateCredentials() error {
	query := `select id,password from users where email = $1`
	row := db.DB.QueryRow(query, u.Email)
	var retreivedPassword string
	err :=row.Scan(&u.Id,&retreivedPassword)
	if err != nil {
		return err
	}
	if utils.AreSame(u.Password, retreivedPassword) {
		return nil
	}
	return errors.New("invalid credentials")

}