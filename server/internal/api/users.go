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

func (*Impl) PostV1Users(c *gin.Context) {
	var user floralApi.PostV1UsersJSONRequestBody
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewJsonErr(err))
		return
	}

	hashedPassword, err := authn.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	userParams := database.CreateUserParams{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    hashedPassword,
		PhoneNumber: user.PhoneNumber,
	}

	createdUser, err := db.NewQueries().CreateUser(context.Background(), userParams)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	c.JSON(http.StatusCreated, floralApi.UserResponse{
		Id:          createdUser.ID,
		Email:       createdUser.Email,
		FirstName:   createdUser.FirstName,
		LastName:    createdUser.LastName,
		PhoneNumber: createdUser.PhoneNumber,
	})
}
func (*Impl) GetV1UsersId(c *gin.Context, id int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	user, err := db.NewQueries().GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, NewJsonErr(ErrNotFound))
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	c.JSON(http.StatusOK, floralApi.UserResponse{
		Id:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	})
}
