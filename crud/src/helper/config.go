package helper

type User struct {
	ID         int     `json:"id"`
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Created_at *string `json:"created_at,omitempty"`
}
