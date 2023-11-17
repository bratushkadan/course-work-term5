package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"floral/config"

	api "floral/generated/api"

	apiImpl "floral/internal/api"
	"floral/internal/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

var tokenMaker auth.Maker

const TokenDuration = 30 * 24 * time.Hour // 30 days

const ErrInternalServerErrorJson = `{"error": "Internal server error."}`

func init() {
	var err error
	tokenMaker, err = auth.NewPasetoMaker(config.App.Token.Secret)
	if err != nil {
		log.Fatal(fmt.Errorf("error creating token maker: %w", err))
	}
}

func main() {
	var floralApi = &apiImpl.Impl{}

	r := gin.Default()
	r.Use(cors.Default())
	api.RegisterHandlers(r, floralApi)

	r.Run()
}

var (
	ErrAuthUsernameNotProvided = errors.New(`query parameter "username" must be provided`)
	ErrAuthPasswordNotProvided = errors.New(`query parameter "password" must be provided`)
	ErrAuthTokenNotProvided    = errors.New(`query parameter "auth_token" must be provided`)
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

// func getTodoHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	w.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(`{"error": Invalid ID}`))
// 		// http.Error(w, , http.StatusBadRequest)
// 		return
// 	}

// 	todo := Todo{
// 		ID:   id,
// 		Text: uuid.NewString(),
// 	}

// 	json.NewEncoder(w).Encode(todo)
// }

// func v1AuthTokenHandler(w http.ResponseWriter, r *http.Request) {
// 	params := r.URL.Query()
// 	errs := []error{}
// 	username, password := params.Get("username"), params.Get("password")
// 	if username == "" {
// 		errs = append(errs, ErrAuthUsernameNotProvided)
// 	}
// 	if password == "" {
// 		errs = append(errs, ErrAuthPasswordNotProvided)
// 	}
// 	if len(errs) > 0 {
// 		respondErrors(w, errs)
// 		return
// 	}

// 	token, err := tokenMaker.CreateToken(username, TokenDuration)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(ErrInternalServerErrorJson))
// 		return
// 	}
// 	json.NewEncoder(w).Encode(struct {
// 		Token string `json:"token"`
// 	}{Token: token})
// }
// func v1AuthTokenVerifyHandler(w http.ResponseWriter, r *http.Request) {
// 	params := r.URL.Query()
// 	errs := []error{}
// 	token := params.Get("auth_token")
// 	if token == "" {
// 		errs = append(errs, ErrAuthTokenNotProvided)
// 	}
// 	if len(errs) > 0 {
// 		respondErrors(w, errs)
// 		return
// 	}

// 	_, err := tokenMaker.VerifyToken(token)
// 	json.NewEncoder(w).Encode(struct {
// 		IsValid bool `json:"is_valid"`
// 	}{IsValid: err == nil})
// }

// func jsonRespMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func respondErrors(w http.ResponseWriter, errors []error) {
// 	errs := make([]string, 0, len(errors))
// 	for _, err := range errors {
// 		errs = append(errs, err.Error())
// 	}
// 	w.WriteHeader(http.StatusBadRequest)
// 	json.NewEncoder(w).Encode(&ErrResponse{Errors: errs})
// }
