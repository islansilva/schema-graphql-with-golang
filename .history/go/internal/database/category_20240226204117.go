package database

import ("database/sql"
		"github.com/google/uuid")


type Category struct {
	db *sql.db
	ID string
	Name string
	Description String
}


func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}


func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err:= c.db.Exec
}