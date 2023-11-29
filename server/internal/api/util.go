package api

import (
	floralApi "floral/generated/api"

	"github.com/getkin/kin-openapi/openapi3"
)

var Swagger *openapi3.T

func init() {
	swaggerSpec, err := floralApi.GetSwagger()
	if err != nil {
		panic(err)
	}
	Swagger = swaggerSpec
}

func NewJsonErr(err error) floralApi.ErrorResponse {
	// return gin.H{"error": err.Error()}
	return floralApi.ErrorResponse{Error: err.Error()}
}
