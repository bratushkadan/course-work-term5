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

func (*Impl) GetV1Products(c *gin.Context, params floralApi.GetV1ProductsParams) {
	if params.SortBy == nil {
		*params.SortBy = "id"
	}
	if params.SortOrder == nil {
		*params.SortOrder = "asc"
	}

	var pParams database.GetProductsParams
	_ = json.Unmarshal([]byte(fmt.Sprintf(`{"%s_%s": true}`, *params.SortBy, *params.SortOrder)), &pParams)

	if params.Filter != nil {
		if params.Filter.FilterLikeName != nil {
			pParams.LkName = true
			pParams.Name = *params.Filter.FilterLikeName
		}
		if params.Filter.FilterStoreId != nil {
			pParams.IsStoreID = true
			pParams.StoreID = *params.Filter.FilterStoreId
		}
		if params.Filter.FilterCategoryId != nil {
			pParams.IsCategoryID = true
			pParams.CategoryID = pgtype.Int4{Int32: *params.Filter.FilterCategoryId, Valid: true}
		}
		if params.Filter.FilterMinHeight != nil {
			pParams.IsMinHeight = true
			pParams.MinHeight = pgtype.Int4{Int32: *params.Filter.FilterMinHeight, Valid: true}
		}
		if params.Filter.FilterMaxHeight != nil {
			pParams.IsMaxHeight = true
			pParams.MaxHeight = pgtype.Int4{Int32: *params.Filter.FilterMaxHeight, Valid: true}
		}
		if params.Filter.FilterMinPrice != nil {
			pParams.IsMinPrice = true
			pParams.MinPrice = *params.Filter.FilterMinPrice
		}
		if params.Filter.FilterMaxPrice != nil {
			pParams.IsMaxPrice = true
			pParams.MaxPrice = *params.Filter.FilterMaxPrice
		}
	}

	rows, err := db.NewQueries().GetProducts(context.Background(), pParams)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	products := make(floralApi.ProductsResponse, 0, len(rows))
	for _, p := range rows {
		product := floralApi.Product{
			Created:     p.Created.Time.UnixMilli(),
			Description: p.Description.String,
			Id:          p.ID,
			ImageUrl:    p.ImageUrl,
			MinHeight:   p.MinHeight.Int32,
			MaxHeight:   p.MaxHeight.Int32,
			Name:        p.Name,
			Price:       p.Price,
			StoreId:     p.StoreID,
			StoreName:   p.StoreName,
		}
		product.Category.Id = p.CategoryID
		product.Category.Name = p.CategoryName
		product.Category.Description = p.CategoryDescription.String
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func (*Impl) PostV1Products(c *gin.Context) {
	var product floralApi.PostV1ProductsJSONRequestBody
	err := json.NewDecoder(c.Request.Body).Decode(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewJsonErr(err))
		return
	}

	productParams := database.AddProductParams{
		StoreID:     product.StoreId,
		Name:        product.Name,
		Description: pgtype.Text{String: product.Description, Valid: true},
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
		MinHeight:   pgtype.Int4{Int32: product.MinHeight, Valid: true},
		MaxHeight:   pgtype.Int4{Int32: product.MaxHeight, Valid: true},
		CategoryID:  pgtype.Int4{Int32: product.CategoryId, Valid: true},
	}

	pr, err := db.NewQueries().AddProduct(context.Background(), productParams)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	createdProduct, err := db.NewQueries().GetProduct(context.Background(), pr.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	productResponse := floralApi.ProductResponse{
		Created:     createdProduct.Created.Time.UnixMilli(),
		Description: createdProduct.Description.String,
		Id:          createdProduct.ID,
		ImageUrl:    createdProduct.ImageUrl,
		MinHeight:   createdProduct.MinHeight.Int32,
		MaxHeight:   createdProduct.MaxHeight.Int32,
		Name:        createdProduct.Name,
		Price:       createdProduct.Price,
		StoreId:     createdProduct.StoreID,
		StoreName:   createdProduct.StoreName,
	}
	productResponse.Category.Id = createdProduct.CategoryID
	productResponse.Category.Name = createdProduct.CategoryName
	productResponse.Category.Description = createdProduct.CategoryDescription.String

	c.JSON(http.StatusCreated, productResponse)
}

func (*Impl) GetV1ProductsId(c *gin.Context, id int32) {
	p, err := db.NewQueries().GetProduct(context.Background(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, NewJsonErr(ErrNotFound))
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	product := floralApi.ProductResponse{
		Created:     p.Created.Time.UnixMilli(),
		Description: p.Description.String,
		Id:          p.ID,
		ImageUrl:    p.ImageUrl,
		MinHeight:   p.MinHeight.Int32,
		MaxHeight:   p.MaxHeight.Int32,
		Name:        p.Name,
		Price:       p.Price,
		StoreId:     p.StoreID,
		StoreName:   p.StoreName,
	}
	product.Category.Id = p.CategoryID
	product.Category.Name = p.CategoryName
	product.Category.Description = p.CategoryDescription.String

	c.JSON(http.StatusOK, product)
}
