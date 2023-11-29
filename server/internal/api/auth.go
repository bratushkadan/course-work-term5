package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	floralApi "floral/generated/api"
	"floral/internal/auth"
	"floral/internal/db"
	authn "floral/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (*Impl) GetV1AuthTokenStore(c *gin.Context, params floralApi.GetV1AuthTokenStoreParams) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	creds, err := db.NewQueries().GetStoreCreds(ctx, params.Email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusBadRequest, NewJsonErr(ErrBadCredentials))
			return
		}
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
func (*Impl) PostV1AuthTokenStore(c *gin.Context, params floralApi.PostV1AuthTokenStoreParams) {
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
func (impl *Impl) GetV1AuthTokenUser(c *gin.Context, params floralApi.GetV1AuthTokenUserParams) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	creds, err := db.NewQueries().GetUserCreds(ctx, params.Email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusBadRequest, NewJsonErr(ErrBadCredentials))
			return
		}
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
