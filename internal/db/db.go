package db

import (
	"database/sql"
	"fmt"
)

const (
    host     = "localhost"
    port     = 5455
    user     = "postgresUser"
    password = "postgresPW"
    dbname   = "postgresDB"
)

func Connect() (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
	// open database
    db, err := sql.Open("postgres", psqlconn)
	
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}