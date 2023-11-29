package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	floralApi "floral/generated/api"
	"floral/generated/database"
	"floral/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (*Impl) PostV1Categories(c *gin.Context) {
	var category floralApi.PostV1CategoriesJSONRequestBody
	err := json.NewDecoder(c.Request.Body).Decode(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewJsonErr(err))
		return
	}

	createdCategory, err := db.NewQueries().AddCategory(context.Background(), database.AddCategoryParams{
		Name:        category.Name,
		Description: pgtype.Text{String: category.Description},
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, floralApi.CategoryResponse{
		Name:        createdCategory.Name,
		Description: createdCategory.Description.String,
	})
}
func (*Impl) GetV1Categories(c *gin.Context) {
	categories, err := db.NewQueries().GetProductsCategories(context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	jsonCategories := make(floralApi.CategoriesResponse, 0, len(categories))

	for _, category := range categories {
		jsonCategories = append(jsonCategories, floralApi.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description.String,
		})
	}
	c.JSON(http.StatusOK, jsonCategories)
}
func (*Impl) GetV1CategoriesId(c *gin.Context, id int32) {
	category, err := db.NewQueries().GetProductsCategory(context.Background(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, NewJsonErr(ErrNotFound))
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, floralApi.CategoryResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description.String,
	})
}
