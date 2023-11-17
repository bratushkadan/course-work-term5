package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppCfg struct {
	Token struct {
		Secret string
	}

	Postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		Dbname   string

		Connstr string
	}
}

var App = AppCfg{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	} else {
		log.Print("Loaded .env")
	}

	App.Token.Secret = os.Getenv("TOKEN_SECRET")

	if App.Token.Secret == "" {
		log.Fatal(`env "TOKEN_SECRET" must not be empty\n`)
	}
}

func init() {
	App.Postgres.Host = "localhost"
	App.Postgres.Port = 5432
	App.Postgres.User = "bratushkadan"
	App.Postgres.Dbname = "floral"

	var (
		password   = os.Getenv("POSTGRES_PASSWORD")
		env_host   = os.Getenv("POSTGRES_HOST")
		env_port   = os.Getenv("POSTGRES_PORT")
		env_user   = os.Getenv("POSTGRES_USER")
		env_dbname = os.Getenv("POSTGRES_DB")
	)

	if password == "" {
		log.Fatalf(`env variable "POSTGRES_PASSWORD" must be provided.`)
	}
	App.Postgres.Password = password
	if env_host != "" {
		App.Postgres.Host = env_host
	}
	if env_port != "" {
		env_port_int, err := strconv.Atoi(env_port)
		if err != nil {
			log.Fatalf(`env variable "POSTGRES_PORT" must be of type int, provided POSTGRES_PORT="%s". | err = %v`, env_port, err)
		}
		App.Postgres.Port = env_port_int
	}
	if env_user != "" {
		App.Postgres.User = env_user
	}
	if env_dbname != "" {
		App.Postgres.Dbname = env_dbname
	}

	App.Postgres.Connstr = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		App.Postgres.User,
		App.Postgres.Password,
		App.Postgres.Host,
		App.Postgres.Port,
		App.Postgres.Dbname,
	)

	fmt.Printf("Postgres cfg %v\n", App.Postgres)
}
