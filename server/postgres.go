package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"floral/database"

	_ "github.com/lib/pq"
)

const (
	host               = "c-c9qjij9qbq5d7msf6plp.rw.mdb.yandexcloud.net"
	port               = 6432
	user               = "bratushkadan"
	dbname             = "common"
	sslmode            = "verify-full"
	targetSessionAttrs = "read-write"
)

var (
	password               = os.Getenv("POSTGRES_PASSWORD")
	dbCertificateAuthority = os.Getenv("PGSSLROOTCERT")
)
var connStr string

func init() {
	if dbCertificateAuthority == "" {
		log.Fatal("Env PGSSLROOTCERT has to be provided.")
	}
}

func init() {
	if password == "" {
		log.Fatalf(`env variable "POSTGRES_PASSWORD" must be provided.`)
	}
	connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&target_session_attrs=%s", user, password, host, port, dbname, sslmode, targetSessionAttrs)
}

type User struct {
	Id       int
	Username string
}

func GetUsers() ([]database.FloralUser, error) {
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	db := database.New(dbConn)

	return db.GetUsers(context.Background())
}
func GetUser(id int32) (database.FloralUser, error) {
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	db := database.New(dbConn)

	return db.GetUser(context.Background(), id)
}
