package database

type Course struct {
	db				*sql.DB
	ID				int
	Name			string
	Description		string
	CategoryID		string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}
