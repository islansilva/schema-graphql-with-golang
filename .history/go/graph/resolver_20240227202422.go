package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import(
	"example/GraphQL/internal/database"
)

type Resolver struct{
	CategoryDB *database.Category
	CoursesDB *database.Course
}
