package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ping", pingHandler).Methods("GET")

	router.HandleFunc("/todo/{id}", getTodoHandler).Methods("GET")

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
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo := Todo{
		ID:   id,
		Text: uuid.NewString(),
	}

	json.NewEncoder(w).Encode(todo)
}
