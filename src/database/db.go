package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CreateConnection() *sql.DB {
	fmt.Println("[Database] Creating Connection...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DB_User"),
		Passwd:               os.Getenv("DB_Password"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_Host") + ":" + os.Getenv("DB_Port"),
		DBName:               os.Getenv("DB_Name"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("[Database] Connection Created")

	return db
}
