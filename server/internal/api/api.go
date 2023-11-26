package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"floral/internal/auth"
	"floral/internal/db"
	authn "floral/pkg/auth"

	floralApi "floral/generated/api"
	"floral/generated/database"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound             = errors.New("not found")
	ErrBadCredentials       = errors.New("bad credentials")
	ErrFailedToGenAuthToken = errors.New("failed to generate auth token")
	ErrInternalServerError  = errors.New("internal server error")
)

type Impl struct {
}

func (*Impl) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, floralApi.PingResponse{Ts: time.Now().UnixMilli()})
}
func (api *Impl) GetPing(c *gin.Context) {
	api.GetHealth(c)
}
func (*Impl) GetV1AuthTokenStore(c *gin.Context, params floralApi.GetV1AuthTokenStoreParams)   {}
func (*Impl) PostV1AuthTokenStore(c *gin.Context, params floralApi.PostV1AuthTokenStoreParams) {}
func (api *Impl) GetV1AuthTokenUser(c *gin.Context, params floralApi.GetV1AuthTokenUserParams) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	creds, err := db.NewQueries().GetUserCreds(ctx, params.Email)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	if authn.MatchPassword(creds.Password, params.Password) != nil {
		c.JSON(http.StatusBadRequest, NewJsonErr(ErrBadCredentials))
		return
	}

	token, err := auth.Token.CreateToken(params.Email, creds.ID, auth.BaseTokenDuration)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrFailedToGenAuthToken))
		return
	}
	c.JSON(http.StatusOK, floralApi.AuthTokenResponse{Token: token})
}
func (*Impl) PostV1AuthTokenUser(c *gin.Context, params floralApi.PostV1AuthTokenUserParams) {
	_, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, floralApi.AuthTokenValidationResponse{
			Valid: false,
		})
		return
	}

	c.JSON(http.StatusOK, floralApi.AuthTokenValidationResponse{
		Valid: true,
	})
}

func (*Impl) GetV1Cart(c *gin.Context, params floralApi.GetV1CartParams) {

	c.JSON(http.StatusOK, []struct{}{})
}

func (*Impl) GetV1Products(c *gin.Context, params floralApi.GetV1ProductsParams) {
	c.JSON(http.StatusOK, []struct{}{})
}

func (*Impl) GetV1ProductsId(c *gin.Context, id int32) {
	c.JSON(http.StatusNotFound, NewJsonErr(ErrNotFound))
}

func (*Impl) GetV1Users(c *gin.Context) {
	users, err := db.NewQueries().GetUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	if users == nil {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (*Impl) PostV1Users(c *gin.Context) {
	// err := OapiRequestBodyValidate(c)
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, NewJsonErr(err))
	// 	return
	// }

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, floralApi.CreateUserResponse{
		Id:          createdUser.ID,
		Email:       createdUser.Email,
		FirstName:   createdUser.FirstName,
		LastName:    createdUser.LastName,
		PhoneNumber: createdUser.PhoneNumber,
	})
}

func (*Impl) GetV1UsersId(c *gin.Context, id int32) {
	c.JSON(http.StatusNotFound, NewJsonErr(ErrNotFound))
}

func (*Impl) ErrorHandler(c *gin.Context, err error, code int) {
	c.JSON(code, NewJsonErr(err))
}

//	func (*Impl) GetV1Favorite(c *gin.Context, params floralApi.GetV1FavoriteParams) {
//		c.JSON(http.StatusOK, []struct{}{})
//	}
