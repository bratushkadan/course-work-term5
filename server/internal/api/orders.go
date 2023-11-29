package api

import (
	"context"
	"encoding/json"
	"errors"
	floralApi "floral/generated/api"
	"floral/generated/database"
	"floral/internal/auth"
	"floral/internal/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var (
	ErrCartEmpty error = errors.New("cart is empty")
)

func (*Impl) GetV1Orders(c *gin.Context, params floralApi.GetV1OrdersParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	rows, err := db.NewQueries().GetOrders(context.Background(), userClaim.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	response := make(floralApi.OrdersResponse, 0, len(rows))

	for _, r := range rows {
		response = append(response, floralApi.Order{
			Id:             r.ID,
			UserId:         r.UserID,
			Status:         floralApi.OrderStatus(r.Status.FloralOrderStatus),
			StatusModified: r.StatusModified.Time.UnixMilli(),
			Created:        r.Created.Time.UnixMilli(),
		})
	}

	c.JSON(http.StatusOK, response)
}
func (*Impl) GetV1OrdersId(c *gin.Context, id int32, params floralApi.GetV1OrdersIdParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	orderInfo, err := db.NewQueries().GetOrder(context.Background(), database.GetOrderParams{
		ID:     id,
		UserID: userClaim.Id,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, NewJsonErr(ErrNotFound))
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	rows, err := db.NewQueries().GetOrderPositions(context.Background(), database.GetOrderPositionsParams{
		OrderID: id,
		UserID:  userClaim.Id,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	orderPositions := make([]floralApi.OrderPosition, 0, len(rows))

	for _, r := range rows {
		orderPositions = append(orderPositions, floralApi.OrderPosition{
			ProductId:    r.ID,
			Name:         r.Name,
			Description:  r.Description.String,
			Quantity:     r.Quantity,
			Price:        r.Price,
			ImageUrl:     r.ImageUrl,
			CategoryId:   r.CategoryID.Int32,
			CategoryName: r.CategoryName,
			StoreId:      r.StoreID,
			StoreName:    r.StoreName,
		})
	}

	response := floralApi.OrderResponse{
		Id:             orderInfo.ID,
		UserId:         orderInfo.UserID,
		Status:         floralApi.OrderStatus(orderInfo.Status.FloralOrderStatus),
		StatusModified: orderInfo.StatusModified.Time.UnixMilli(),
		Created:        orderInfo.Created.Time.UnixMilli(),
		Positions:      orderPositions,
	}

	response.Created = orderInfo.Created.Time.UnixMilli()

	c.JSON(http.StatusOK, response)
}
func (*Impl) PostV1Orders(c *gin.Context, params floralApi.PostV1OrdersParams) {
	userClaim, err := auth.Token.VerifyToken(params.XAuthToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, NewJsonErr(ErrUnauthorized))
		return
	}

	rows, err := db.NewQueries().GetCart(context.Background(), userClaim.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(rows) == 0 {
		c.JSON(http.StatusBadRequest, NewJsonErr(ErrCartEmpty))
		return
	}

	addedOrder, err := db.NewQueries().AddOrder(context.Background(), userClaim.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	orderPositions := make([]database.AddOrderPositionsParams, 0, len(rows))
	for _, r := range rows {
		orderPositions = append(orderPositions, database.AddOrderPositionsParams{
			OrderID:   addedOrder.ID,
			ProductID: r.ProductID,
			Quantity:  r.Quantity,
		})
	}

	_, err = db.NewQueries().AddOrderPositions(context.Background(), orderPositions)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	_, err = db.NewQueries().ClearCart(context.Background(), userClaim.Id)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, floralApi.CreatedOrderResponse{
		Id:             addedOrder.ID,
		UserId:         addedOrder.UserID,
		Status:         floralApi.OrderStatus(addedOrder.Status.FloralOrderStatus),
		StatusModified: addedOrder.Created.Time.UnixMilli(),
		Created:        addedOrder.Created.Time.UnixMilli(),
	})
}
func (*Impl) PatchV1Orders(c *gin.Context) {
	// Не должно использоваться пользователем, в системе нет роли для admin/store *owner*, поэтому
	// ручку можно дернуть без токена
	var requestBody floralApi.PatchV1OrdersJSONRequestBody
	_ = requestBody
	err := json.NewDecoder(c.Request.Body).Decode(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewJsonErr(err))
		return
	}

	order, err := db.NewQueries().UpdateOrderStatus(context.Background(), database.UpdateOrderStatusParams{
		ID: requestBody.Id,
		Status: database.NullFloralOrderStatus{
			FloralOrderStatus: database.FloralOrderStatus(requestBody.Status),
			Valid:             true,
		},
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewJsonErr(err))
		return
	}

	c.JSON(http.StatusOK, floralApi.Order{
		Id:             order.ID,
		UserId:         order.UserID,
		Created:        order.Created.Time.UnixMilli(),
		Status:         floralApi.OrderStatus(order.Status.FloralOrderStatus),
		StatusModified: order.Created.Time.UnixMilli(),
	})
}
