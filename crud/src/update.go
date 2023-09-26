package src

import (
	"fmt"
	"net/http"
	. "nuite/crud/src/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func UpdateUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fieldsToUpdate := map[string]interface{}{}

	// fill the map with the User json object
	fieldsToUpdate["id"] = user.ID
	fieldsToUpdate["firstname"] = user.Firstname
	fieldsToUpdate["lastname"] = user.Lastname
	fieldsToUpdate["email"] = user.Email
	fieldsToUpdate["phone"] = user.Phone

	// validate the fields
	isNotValid := ValidateUpdateUserInput(fieldsToUpdate)

	// return error if any
	if isNotValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": isNotValid.Error()})
		return
	}

	// search for user with this id if exists
	var userIdDb, _ = GetUserByID(Db, user.ID)

	// return error if not
	if userIdDb == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrUserNotFound.Error()})
		return
	}

	// remove the id from the Map, because it gonna be used to generate the query with correspond colums to update
	delete(fieldsToUpdate, "id")

	// if the field provided to update contain same existent value, so no need to update it, said so, we remove it from the field Map
	if userIdDb.Firstname == fieldsToUpdate["firstname"] {
		delete(fieldsToUpdate, "firstname")
	}
	if userIdDb.Lastname == fieldsToUpdate["lastname"] {
		delete(fieldsToUpdate, "lastname")
	}
	if userIdDb.Email == fieldsToUpdate["email"] {
		delete(fieldsToUpdate, "email")
	}
	if userIdDb.Phone == fieldsToUpdate["phone"] {
		delete(fieldsToUpdate, "phone")
	}
	// if there is some fields containing new values other than one in Db, row have to be updated after checking if email (if provided) not already in use.
	if len(fieldsToUpdate) > 0 {
		_, err := UpdateUserByID(user.ID, fieldsToUpdate)
		if err != nil {
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
		c.JSON(http.StatusCreated, "Document updated succefully")
		return
	}
	// if all data is the same and nothing has really changed, so no need to make an update query to DB
	c.JSON(http.StatusAlreadyReported, "No fields to update!")
	return
}
