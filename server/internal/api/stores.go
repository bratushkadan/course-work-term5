package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	floralApi "floral/generated/api"
	"floral/generated/database"
	"floral/internal/db"
	authn "floral/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (*Impl) GetV1Stores(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	stores, err := db.NewQueries().GetStores(ctx)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}
	jsonStores := make([]floralApi.StoreResponse, 0, len(stores))
	for _, store := range stores {
		jsonStores = append(jsonStores, floralApi.StoreResponse{
			Id:          store.ID,
			Name:        store.Name,
			Email:       store.Email,
			PhoneNumber: store.PhoneNumber,
			Created:     store.Created.Time.UnixMilli(),
		})
	}
	c.JSON(http.StatusOK, jsonStores)
}
func (*Impl) PostV1Stores(c *gin.Context) {
	var store floralApi.PostV1StoresJSONRequestBody
	err := json.NewDecoder(c.Request.Body).Decode(&store)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewJsonErr(err))
		return
	}

	hashedPassword, err := authn.HashPassword(store.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	storeParams := database.AddStoreParams{
		Name:        store.Name,
		Email:       store.Email,
		Password:    hashedPassword,
		PhoneNumber: store.PhoneNumber,
	}

	createdStore, err := db.NewQueries().AddStore(context.Background(), storeParams)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, floralApi.StoreResponse{
		Id:          createdStore.ID,
		Email:       createdStore.Email,
		Name:        createdStore.Name,
		PhoneNumber: createdStore.PhoneNumber,
		Created:     createdStore.Created.Time.UnixMilli(),
	})
}
func (*Impl) GetV1StoresId(c *gin.Context, id int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	store, err := db.NewQueries().GetStore(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, NewJsonErr(ErrNotFound))
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}
	c.JSON(http.StatusOK, floralApi.StoreResponse{
		Id:          store.ID,
		Name:        store.Name,
		Email:       store.Email,
		PhoneNumber: store.PhoneNumber,
		Created:     store.Created.Time.UnixMilli(),
	})
}
