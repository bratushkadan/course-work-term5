package api

import (
	"context"
	"errors"
	floralApi "floral/generated/api"
	"floral/generated/database"
	"floral/internal/auth"
	"floral/internal/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrUserHasNotPurchasedProduct error = errors.New("user has not purchased product, can't add review")
)

func (*Impl) GetV1Reviews(c *gin.Context, params floralApi.GetV1ReviewsParams) {
	if params.ProductId == nil && params.StoreId == nil && params.UserId == nil {
		c.JSON(http.StatusBadRequest, NewJsonErr(errors.New(
			`at least one of the parameters "product_id", "store_id", "user_id" has to be specified`,
		)))
		return
	}

	dbQuery := database.GetProductReviewsParams{}
	if params.ProductId != nil {
		dbQuery.IsProductID = false
		dbQuery.ProductID = *params.ProductId
	}
	if params.StoreId != nil {
		dbQuery.IsStoreID = false
		dbQuery.StoreID = *params.StoreId
	}
	if params.UserId != nil {
		dbQuery.IsUserID = false
		dbQuery.UserID = pgtype.Int4{Int32: *params.UserId, Valid: true}
	}

	rows, err := db.NewQueries().GetProductReviews(context.Background(), dbQuery)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	response := make(floralApi.ReviewsResponse, 0, len(rows))
	for _, r := range rows {
		if username, ok := r.UserName.(string); ok {
			response = append(response, floralApi.Review{
				Created:    r.Created.Time.UnixMilli(),
				Id:         r.ID,
				Modified:   r.Modified.Time.UnixMilli(),
				ProductId:  r.ProductID,
				Rating:     r.Rating,
				ReviewText: r.ReviewText.String,
				UserId:     r.UserID.Int32,
				UserName:   username,
			})
		} else {
			fmt.Println(errors.New("failed to convert r.UserName to string, should not happen"))
			c.JSON(http.StatusInternalServerError, ErrInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, response)
}
func (*Impl) PostV1Reviews(c *gin.Context, params floralApi.PostV1ReviewsParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	var requestBody floralApi.PostV1ReviewsJSONRequestBody
	_ = requestBody

	exists, err := db.NewQueries().GetUserPurchasedProduct(
		context.Background(),
		database.GetUserPurchasedProductParams{
			UserID:    userClaim.Id,
			ProductID: requestBody.ProductId,
		},
	)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, NewJsonErr(ErrInternalServerError))
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, NewJsonErr(ErrUserHasNotPurchasedProduct))
		return
	}

	reviewId, err := db.NewQueries().AddProductReview(
		context.Background(),
		database.AddProductReviewParams{
			UserID:     pgtype.Int4{Int32: userClaim.Id, Valid: true},
			ProductID:  requestBody.ProductId,
			Rating:     float64(requestBody.Rating),
			ReviewText: pgtype.Text{String: requestBody.ReviewText, Valid: true},
		},
	)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	productReview, err := db.NewQueries().GetProductReview(
		context.Background(),
		reviewId,
	)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	if username, ok := productReview.UserName.(string); ok {
		c.JSON(http.StatusOK, floralApi.Review{
			Created:    productReview.Created.Time.UnixMilli(),
			Id:         productReview.ID,
			Modified:   productReview.Modified.Time.UnixMilli(),
			ProductId:  productReview.ProductID,
			Rating:     productReview.Rating,
			ReviewText: productReview.ReviewText.String,
			UserId:     productReview.UserID.Int32,
			UserName:   username,
		})
	} else {
		fmt.Println(errors.New("failed to convert r.UserName to string, should not happen"))
		c.JSON(http.StatusInternalServerError, ErrInternalServerError)
		return
	}
}
func (*Impl) PatchV1Reviews(c *gin.Context, params floralApi.PatchV1ReviewsParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	var requestBody floralApi.PatchV1ReviewsJSONRequestBody
	_ = requestBody

	reviewId, err := db.NewQueries().UpdateProductReview(
		context.Background(),
		database.UpdateProductReviewParams{
			UserID:     pgtype.Int4{Int32: userClaim.Id, Valid: true},
			ProductID:  requestBody.ProductId,
			Rating:     float64(requestBody.Rating),
			ReviewText: pgtype.Text{String: requestBody.ReviewText, Valid: true},
		},
	)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	productReview, err := db.NewQueries().GetProductReview(
		context.Background(),
		reviewId,
	)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	if username, ok := productReview.UserName.(string); ok {
		c.JSON(http.StatusOK, floralApi.Review{
			Created:    productReview.Created.Time.UnixMilli(),
			Id:         productReview.ID,
			Modified:   productReview.Modified.Time.UnixMilli(),
			ProductId:  productReview.ProductID,
			Rating:     productReview.Rating,
			ReviewText: productReview.ReviewText.String,
			UserId:     productReview.UserID.Int32,
			UserName:   username,
		})
	} else {
		fmt.Println(errors.New("failed to convert r.UserName to string, should not happen"))
		c.JSON(http.StatusInternalServerError, ErrInternalServerError)
		return
	}
}
func (*Impl) DeleteV1Reviews(c *gin.Context, params floralApi.DeleteV1ReviewsParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	reviewId, err := db.NewQueries().DeleteProductReview(
		context.Background(),
		database.DeleteProductReviewParams{
			UserID:    pgtype.Int4{Int32: userClaim.Id, Valid: true},
			ProductID: params.ProductId,
		},
	)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(ErrInternalServerError))
		return
	}

	c.JSON(http.StatusOK, floralApi.ReviewDeletedResponse{Id: reviewId})
}

func (*Impl) ErrorHandler(c *gin.Context, err error, code int) {
	c.JSON(code, NewJsonErr(err))
}
