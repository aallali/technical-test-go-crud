package src

import (
	"net/http"
	. "nuite/crud/src/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

// return user by id if given in query params
// return all list of users if no id provided
// TODO: set pagination in case of big data
func GetUsers(c *gin.Context) {
	var id, foundId = c.GetQuery("id")
	if foundId {
		idInt, _ := strconv.Atoi(id)
		var userIdDb, _ = GetUserByID(Db, idInt)

		if userIdDb == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ErrUserNotFound.Error()})
			return
		}
		c.JSON(http.StatusOK, userIdDb)

	} else {
		rows, err := Db.Query("SELECT id, firstname, lastname, email, phone, created_at FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Created_at); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			users = append(users, user)
		}
		c.JSON(http.StatusOK, users)
	}

}
