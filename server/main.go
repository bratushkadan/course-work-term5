package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"floral/auth"
	"floral/database"

	"floral/internal/app"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Foo interface {
	fmt.Stringer
}

type Bar struct {
	FirstName string
	LastName  string
}

func (b *Bar) String() string {
	return fmt.Sprintf("%s %s", b.FirstName, b.LastName)
}

func init() {
	b := &Bar{FirstName: "Dan", LastName: "Bratushka"}
	foobar(b)
}

func foobar(foo Foo) string {
	return foo.String()
}

var tokenSecret = os.Getenv("TOKEN_SECRET")

var tokenMaker auth.Maker

const TokenDuration = 30 * 24 * time.Hour // 30 days

const ErrInternalServerErrorJson = `{"error": "Internal server error."}`

func init() {
	if tokenSecret == "" {
		log.Fatal(`env "TOKEN_SECRET" must not be empty\n`)
	}

	var err error
	tokenMaker, err = auth.NewPasetoMaker(tokenSecret)
	if err != nil {
		log.Fatal(fmt.Errorf("error creating token maker: %w", err))
	}
}

func main() {
	_ = app.InitApp()

	router := mux.NewRouter()

	router.Use(jsonRespMiddleware)

	v1Router := router.PathPrefix("/v1").Subrouter()

	v1Auth := v1Router.PathPrefix("/auth").Subrouter()

	v1Auth.HandleFunc("/token", v1AuthTokenHandler).Methods("GET")
	v1Auth.HandleFunc("/token/verify", v1AuthTokenVerifyHandler).Methods("POST")

	v1User := v1Router.PathPrefix("/user").Subrouter()

	usersHandler := func(w http.ResponseWriter, r *http.Request) {
		users, err := GetUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{}`))
			return
		}
		if users == nil {
			users = []database.FloralUser{}
		}
		json.NewEncoder(w).Encode(users)
	}
	v1User.HandleFunc("", usersHandler).Methods("GET")
	v1User.HandleFunc("/", usersHandler).Methods("GET")
	v1User.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(struct {
				Error string `json:"error"`
			}{Error: `query parameter "id" has to be of type int32.`})
		}
		user, err := GetUser(int32(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(`{}`))
			return
		}
		json.NewEncoder(w).Encode(user)
	})

	router.HandleFunc("/ping", pingHandler).Methods("GET")

	router.HandleFunc("/todo/{id}", getTodoHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{}"))
	})

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		Ts int64 `json:"ts"`
	}{
		Ts: time.Now().UnixMilli(),
	})
}

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": Invalid ID}`))
		// http.Error(w, , http.StatusBadRequest)
		return
	}

	todo := Todo{
		ID:   id,
		Text: uuid.NewString(),
	}

	json.NewEncoder(w).Encode(todo)
}

var (
	ErrAuthUsernameNotProvided = errors.New(`query parameter "username" must be provided`)
	ErrAuthPasswordNotProvided = errors.New(`query parameter "password" must be provided`)
	ErrAuthTokenNotProvided    = errors.New(`query parameter "auth_token" must be provided`)
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

func v1AuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	errs := []error{}
	username, password := params.Get("username"), params.Get("password")
	if username == "" {
		errs = append(errs, ErrAuthUsernameNotProvided)
	}
	if password == "" {
		errs = append(errs, ErrAuthPasswordNotProvided)
	}
	if len(errs) > 0 {
		respondErrors(w, errs)
		return
	}

	token, err := tokenMaker.CreateToken(username, TokenDuration)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrInternalServerErrorJson))
		return
	}
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
	}{Token: token})
}
func v1AuthTokenVerifyHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	errs := []error{}
	token := params.Get("auth_token")
	if token == "" {
		errs = append(errs, ErrAuthTokenNotProvided)
	}
	if len(errs) > 0 {
		respondErrors(w, errs)
		return
	}

	_, err := tokenMaker.VerifyToken(token)
	json.NewEncoder(w).Encode(struct {
		IsValid bool `json:"is_valid"`
	}{IsValid: err == nil})
}

func jsonRespMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func respondErrors(w http.ResponseWriter, errors []error) {
	errs := make([]string, 0, len(errors))
	for _, err := range errors {
		errs = append(errs, err.Error())
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&ErrResponse{Errors: errs})
}
