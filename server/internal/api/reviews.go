package api

import (
	floralApi "floral/generated/api"

	"github.com/gin-gonic/gin"
)

func (*Impl) GetV1Reviews(c *gin.Context, params floralApi.GetV1ReviewsParams) {}
func (*Impl) PostV1Reviews(c *gin.Context, params floralApi.PostV1ReviewsParams) {
	var requestBody floralApi.PostV1ReviewsJSONRequestBody
	_ = requestBody
}
func (*Impl) PatchV1Reviews(c *gin.Context, params floralApi.PatchV1ReviewsParams) {
	var requestBody floralApi.PatchV1ReviewsJSONRequestBody
	_ = requestBody
}
func (*Impl) DeleteV1Reviews(c *gin.Context, params floralApi.DeleteV1ReviewsParams) {}

func (*Impl) ErrorHandler(c *gin.Context, err error, code int) {
	c.JSON(code, NewJsonErr(err))
}
