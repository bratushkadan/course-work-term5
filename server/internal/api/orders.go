package api

import (
	floralApi "floral/generated/api"

	"github.com/gin-gonic/gin"
)

func (*Impl) GetV1Orders(c *gin.Context, params floralApi.GetV1OrdersParams)               {}
func (*Impl) GetV1OrdersId(c *gin.Context, id int32, params floralApi.GetV1OrdersIdParams) {}
func (*Impl) PostV1Orders(c *gin.Context, params floralApi.PostV1OrdersParams) {
}
func (*Impl) PatchV1Orders(c *gin.Context) {
	var requestBody floralApi.PatchV1OrdersJSONRequestBody
	_ = requestBody
}
