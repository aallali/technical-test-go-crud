package helper

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	var mysqlUser = os.Getenv("mysql_user")
	var mysqlPass = os.Getenv("mysql_pass")
	var mysqlHost = os.Getenv("mysql_host")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/myDb", mysqlUser, mysqlPass, mysqlHost)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	_, err = Db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INT(20) unsigned AUTO_INCREMENT,
		firstname VARCHAR(20),
		lastname VARCHAR(20),
		email VARCHAR(100) UNIQUE,
		phone VARCHAR(40),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

		PRIMARY KEY (id)
	);

	`)
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	Db.Close()
}
