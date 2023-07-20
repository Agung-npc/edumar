package conn

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	. "github.com/dwirobbin/edumar-backend/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DBConnection() (*sql.DB, error) {
	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	mustGetenv := func(key string) string {
		envName := os.Getenv(key)
		if envName == "" && key == "DB_PASS" {
			envName = ""
			if key != "DB_PASS" {
				panic("value in key " + key + " not found")
			}
		}
		return envName
	}

	dbDriver := mustGetenv("DB_DRIVER")
	dbUser := mustGetenv("DB_USER")
	dbPass := mustGetenv("DB_PASS")
	dbHost := mustGetenv("DB_HOST")
	dbPort := mustGetenv("DB_PORT")
	dbName := mustGetenv("DB_NAME")

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	db, err := sql.Open(dbDriver, dataSourceName)
	PanicIfError(err)

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	_, err = db.Exec(query)
	PanicIfError(err)

	log.Printf("Successfully created database: %s.db", dbName)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	PanicIfError(err)

	log.Println("Successfully connected to database")

	return db, nil
}
