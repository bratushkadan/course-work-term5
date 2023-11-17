package db

import (
	"context"
	"database/sql"

	"floral/generated/database"

	"floral/config"

	_ "github.com/lib/pq"
)

// TODO: create pool of connections instead of opening a connection on earch

type User struct {
	Id       int
	Username string
}

func GetUsers() ([]database.FloralUser, error) {
	dbConn, err := sql.Open("postgres", config.App.Postgres.Connstr)
	if err != nil {
		return nil, err
	}

	defer dbConn.Close()

	db := database.New(dbConn)

	return db.GetUsers(context.Background())
}
func GetUser(id int32) (*database.FloralUser, error) {
	dbConn, err := sql.Open("postgres", config.App.Postgres.Connstr)
	if err != nil {
		return nil, err
	}

	defer dbConn.Close()

	db := database.New(dbConn)

	user, err := db.GetUser(context.Background(), id)

	return &user, err
}
