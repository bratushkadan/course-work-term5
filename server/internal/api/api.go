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
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
func (api *Impl) GetV1AuthTokenUser(c *gin.Context, params floralApi.GetV1AuthTokenUserParams) {
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

// FIXME: add logic
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
	}
	product.Category.Id = p.CategoryID
	product.Category.Name = p.CategoryName
	product.Category.Description = p.CategoryDescription.String

	c.JSON(http.StatusOK, product)
}

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

// func (*Impl) GetV1Cart(c *gin.Context, params floralApi.GetV1CartParams) {

// 	c.JSON(http.StatusOK, []struct{}{})
// }

func (*Impl) ErrorHandler(c *gin.Context, err error, code int) {
	c.JSON(code, NewJsonErr(err))
}

//	func (*Impl) GetV1Favorite(c *gin.Context, params floralApi.GetV1FavoriteParams) {
//		c.JSON(http.StatusOK, []struct{}{})
//	}
