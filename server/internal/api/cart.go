package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	floralApi "floral/generated/api"
	"floral/generated/database"
	"floral/internal/auth"
	"floral/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (*Impl) GetV1Cart(c *gin.Context, params floralApi.GetV1CartParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	userId := userClaim.Id
	rows, err := db.NewQueries().GetCart(context.Background(), userId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	response := make(floralApi.CartResponse, 0, len(rows))

	for _, r := range rows {
		response = append(response, floralApi.CartProduct{
			ProductId:    r.ProductID,
			Name:         r.Name,
			Description:  r.Description.String,
			ImageUrl:     r.ImageUrl,
			Price:        r.Price,
			Quantity:     r.Quantity,
			CategoryId:   r.CategoryID.Int32,
			CategoryName: r.CategoryName,
			StoreId:      r.StoreID,
			StoreName:    r.StoreName,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (*Impl) PostV1Cart(c *gin.Context, params floralApi.PostV1CartParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	var requestBody floralApi.PostV1CartJSONBody
	err = json.NewDecoder(c.Request.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, NewJsonErr(err))
		return
	}

	userId := userClaim.Id

	var response floralApi.CartPositionResponse

	if requestBody.Quantity != 0 {
		var resultRow database.UpsertCartPositionRow
		resultRow, err = db.NewQueries().UpsertCartPosition(context.Background(), database.UpsertCartPositionParams{
			UserID:    userId,
			ProductID: requestBody.ProductId,
			Quantity:  requestBody.Quantity,
		})

		response.ProductId = resultRow.ProductID
		response.Quantity = resultRow.Quantity
	} else {
		var resultRow database.RemoveCartPositionRow
		resultRow, err = db.NewQueries().RemoveCartPosition(context.Background(), database.RemoveCartPositionParams{
			UserID:    userId,
			ProductID: requestBody.ProductId,
		})

		response.ProductId = resultRow.ProductID
		response.Quantity = 0
	}

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, NewJsonErr(errors.New("no cart position to remove")))
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
func (impl *Impl) PutV1Cart(c *gin.Context, params floralApi.PutV1CartParams) {
	impl.PostV1Cart(c, floralApi.PostV1CartParams(params))
}
func (impl *Impl) PatchV1Cart(c *gin.Context, params floralApi.PatchV1CartParams) {
	impl.PostV1Cart(c, floralApi.PostV1CartParams(params))
}
func (*Impl) DeleteV1Cart(c *gin.Context, params floralApi.DeleteV1CartParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	rows, err := db.NewQueries().ClearCart(context.Background(), userClaim.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	response := make(floralApi.CartPositionsResponse, 0, len(rows))

	for _, r := range rows {
		response = append(response, floralApi.CartPosition{
			ProductId: r.ProductID,
			Quantity:  r.Quantity,
		})
	}

	c.JSON(http.StatusOK, response)
}
