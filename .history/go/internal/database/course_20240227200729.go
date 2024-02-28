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


func (c *Course) Create(name string, description string, categoryID string) (Category, error) {
	id := uuid.New().String()
	_, err:= c.db.Exec("INSERT INTO courses (id, name, description, categoryid) VALUES ($1, $2, $3, $4)", id, name, description, categoryI)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: &description}, nil
}