package database

import "database/sql"


type Category struct {
	db *sql.db
	ID string
	Name string
	Description String
}


func NewCategory(db *sql.DB) 