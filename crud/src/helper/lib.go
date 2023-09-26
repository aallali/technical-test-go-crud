package helper

import (
	"database/sql"
	"fmt"

	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// updateUserByID updates a user row in the database based on the provided fields.
// id: The ID of the user to update.
// fieldsToUpdate: A map where keys are field names (e.g., "firstname", "lastname") and values are new values for those fields.
func UpdateUserByID(id int, fieldsToUpdate map[string]interface{}) (interface{}, error) {
	// Create a SQL update query with placeholders for the fields to update
	updateQuery := "UPDATE users SET "
	var updateValues []interface{}

	for fieldName, fieldValue := range fieldsToUpdate {
		updateQuery += fmt.Sprintf("%s = ?, ", fieldName)
		updateValues = append(updateValues, fieldValue)
	}

	// Remove the trailing ", " from the query
	updateQuery = strings.TrimSuffix(updateQuery, ", ")

	// Add the WHERE clause to specify the user by ID
	updateQuery += " WHERE id = ?"
	updateValues = append(updateValues, id)

	// Execute the update query with the provided values
	res, err := Db.Exec(updateQuery, updateValues...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func GetUserByID(db *sql.DB, userID int) (*User, error) {
	// Define the SQL query to retrieve a user by ID
	query := "SELECT id, firstname, lastname, email, phone FROM users WHERE id = ?"

	// Query the database to fetch the user by ID
	row := db.QueryRow(query, userID)

	// Create a User struct to store the query result
	var user User

	// Scan the query result into the User struct
	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no rows were found, return a custom error or handle it as needed.
			return nil, fmt.Errorf("User with ID %d not found", userID)
		}
		// Handle other errors as needed.
		return nil, err
	}

	return &user, nil
}
