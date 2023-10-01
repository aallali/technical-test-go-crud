package src

import (
	"fmt"
	"net/http"
	. "nuite/crud/src/helper"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

var fileName string

func users_table_to_excel() (err error) {
	currentTime := time.Now()

	fileName = "Users-export-" + currentTime.Format("01-02-2006") + ".xlsx"
	xlsx := excelize.NewFile()

	// Create a new sheet and set its name
	sheetName := "Sheet1" // write in first file
	// xlsx.NewSheet(sheetName)
	rows, err := Db.Query("SELECT id, firstname, lastname, email, phone, created_at FROM users")
	if err != nil {
		return
	}
	defer rows.Close()
	rowIndex := 0 // Start from row 0 (first line in excel file)

	// Iterate through the query result and stream data to Excel
	headers := [6]string{"ID", "first name", "last name", "e-mail", "phone number", "added date"}
	for colIndex, cellValue := range headers {
		cellName := excelize.ToAlphaString(colIndex) + fmt.Sprint(rowIndex+1)
		xlsx.SetCellValue(sheetName, cellName, cellValue)
	}
	rowIndex++
	for rows.Next() {
		var id int
		var firstname, lastname, email, phone, created_at string

		err = rows.Scan(&id, &firstname, &lastname, &email, &phone, &created_at)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		cellValues := []interface{}{id, firstname, lastname, email, phone}
		// Write data to the Excel sheet row by row
		// NOTE : that i am streaming one line at time from Database to avoid loading all data to RAM while the program is runing (in case there is bug data)
		for colIndex, cellValue := range cellValues {
			cellName := excelize.ToAlphaString(colIndex) + fmt.Sprint(rowIndex+1)
			xlsx.SetCellValue(sheetName, cellName, cellValue)
		}

		rowIndex++
	}
	// Save the Excel file
	if err = xlsx.SaveAs(fileName); err != nil {
		fmt.Println("Error saving Excel file:", err)
		return
	}

	return
}

func uploadFileToS3() {
	fmt.Printf("Uploading (%s) to S3 Bucket ...\n", fileName)
}

func ExportToS3(c *gin.Context) {
	err := users_table_to_excel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	uploadFileToS3()
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Data exported in (%s.xls) to S3 Bucket.", fileName)})
}
