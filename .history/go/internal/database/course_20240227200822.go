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
	_, err:= c.db.Exec("INSERT INTO courses (id, name, description, categoryid) VALUES ($1, $2, $3, $4)", id, name, description, categoryID)

	if err != nil {
		return Category{}, err
	}

	return Course{ID: id, Name: name, Description: &description, CategoryID: categoryID}, nil
}



func (c *Course) FindAll()  ([]Courses, error) {
	rows, err := c.db.Query("SELECT id, name, description, categoryid FROM courses")
	if err!=nil {
		return nil, err
	}

	defer rows.Close()

	categories:=[]Category{}

	for rows.Next () {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		categories = append(categories, Category{ID: id, Name: name, Description: &description})
	}

	return categories, nil


}