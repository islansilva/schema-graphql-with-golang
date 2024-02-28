package database

type Course struct {
	db				*sql.DB
	ID				int
	Name			string
	Description		string
	CategoryID		string
}

func NewCourse(db *sql.DB) *Category {
	return &Category{db: db}
}


func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err:= c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: &description}, nil
}