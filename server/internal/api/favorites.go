package api

import (
	"context"
	floralApi "floral/generated/api"
	"floral/generated/database"
	"floral/internal/auth"
	"floral/internal/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Impl) GetV1Favorites(c *gin.Context, params floralApi.GetV1FavoritesParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	rows, err := db.NewQueries().GetFavorites(context.Background(), userClaim.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	response := make(floralApi.Favorites, 0, len(rows))

	for _, r := range rows {
		response = append(response, floralApi.Favorite{
			AddedFavorite: r.AddedFavorite.Time.UnixMilli(),
			CategoryId:    r.CategoryID,
			CategoryName:  r.CategoryName,
			Description:   r.Description.String,
			ImageUrl:      r.ImageUrl,
			Price:         r.Price,
			ProductId:     r.ID,
			Name:          r.Name,
			StoreId:       r.StoreID,
			StoreName:     r.StoreName,
		})
	}

	c.JSON(http.StatusOK, response)
}
func (*Impl) PostV1Favorites(c *gin.Context, params floralApi.PostV1FavoritesParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	id, err := db.NewQueries().AddFavorite(context.Background(), database.AddFavoriteParams{
		UserID:    userClaim.Id,
		ProductID: params.ProductId,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	c.JSON(http.StatusOK, floralApi.FavoriteProductResponse{Id: id})
}
func (*Impl) DeleteV1Favorites(c *gin.Context, params floralApi.DeleteV1FavoritesParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	id, err := db.NewQueries().DeleteFavorite(context.Background(), database.DeleteFavoriteParams{
		UserID:    userClaim.Id,
		ProductID: params.ProductId,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	c.JSON(http.StatusOK, floralApi.FavoriteProductResponse{Id: id})
}
