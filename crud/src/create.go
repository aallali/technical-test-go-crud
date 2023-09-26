package src

import (
	"fmt"
	"net/http"
	. "nuite/crud/src/helper"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"
)

func CreateUser(c *gin.Context) {
	var user User
	// parse the body into the User schema declared in helper/config/User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the input for rules set
	err := ValidateInput(user.Firstname, user.Lastname, user.Email, user.Phone, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create new row with data provided
	result, err := Db.Exec(`INSERT INTO users (firstname, lastname, email, phone) VALUES (?, ?, ?, ?);`,
		&user.Firstname, &user.Lastname, &user.Email, &user.Phone)

	if err != nil {
		// if the email is already used by another user, you cant use it again
		// Check if the error is a MySQL duplicate entry error (Error 1062)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			// Handle duplicate email error and return a custom error message
			msgErr := fmt.Errorf("Email '%s' is already in use", user.Email)
			c.JSON(http.StatusConflict, gin.H{"error": msgErr.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// retrieve the id given to this user we just created
	id, _ := result.LastInsertId()
	user.ID = int(id)

	// return the user data object
	c.JSON(http.StatusCreated, user)
	return
}
