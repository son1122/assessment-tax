package model

// User represents a user in the system
type User struct {
	ID   int
	Name string
}

// GetUsers is a hypothetical function that fetches users from a datastore
func GetUsers() ([]User, error) {
	//db.DB.QueryRow()

	return []User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Doe"},
	}, nil
}
