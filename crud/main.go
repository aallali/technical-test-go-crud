package main

import (
	. "nuite/crud/src"
	. "nuite/crud/src/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()
	defer CloseDB()

	r := gin.Default()

	r.POST("/users", CreateUser)
	r.PATCH("/users", UpdateUser)
	r.GET("/users", GetUsers)
	r.DELETE("/users/:id", DeleteUser)
	r.GET("/export", ExportToS3)
	r.Run(":1337")
}
