package app

import (
	"database/sql"
	"fmt"
	"go_fiber/helper"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {

	var (	
		username = envVariable("DB_USERNAME")
		password = envVariable("DB_PASSWORD")
		host = envVariable("DB_HOST")
		port = envVariable("DB_PORT")
		db_name = envVariable("DB_NAME")
	)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", username, password, host, port, db_name)

	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

// func CreateDatabase(db *sql.DB){
// 	query := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id INT AUTO_INCREMENT PRIMARY KEY,
// 		name VARCHAR(255) NOT NULL,
// 		email VARCHAR(255) NOT NULL,
// 		password VARCHAR(255) NOT NULL,
// 		address VARCHAR(255) NOT NULL,
// 		phone VARCHAR(255) NOT NULL,
// 		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
// 		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
// 	)
// 	`
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		panic(err)
// 	}	

// }

// func AddColumnDB(db *sql.DB)  {
// 	query := `
// 		ALTER TABLE users 
// 		ADD COLUMN verifyCode VARCHAR(255) NOT NULL, 
// 		ADD COLUMN isVerify VARCHAR(255) NOT NULL;
// 	`

// 	_, err := db.Exec(query)
// 	if err != nil {
// 		panic(err)
// 	}	
// }

func envVariable(key string) string {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	return os.Getenv(key)
}