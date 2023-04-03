package db

import (
	"database/sql"
	"fmt"
)

type DB struct {
    *sql.DB
}

func ConnectDB(host string, port int, user, password, dbname string) (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// // CreateSchema creates the necessary tables in the database
// func CreateSchema(db *sql.DB) error {
// 	sqlFile, err := os.ReadFile("db/schema.sql")
// 	if err != nil {
// 	  return err
// 	}
// 	_, err = db.Exec(string(sqlFile))
// 	if err != nil {
// 	  return err
// 	}
// 	return nil
//   }