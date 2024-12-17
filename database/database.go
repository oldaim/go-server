package database

import "database/sql"

var DB *sql.DB

func Initialize() error {

	db, err := InitDB()

	if err != nil {
		return err
	}

	DB = db

	return nil

}
