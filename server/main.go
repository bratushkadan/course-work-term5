package main

import (
	api "floral/generated/api"

	apiImpl "floral/internal/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func main() {
	var floralApi = &apiImpl.Impl{}

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.OapiRequestValidator(apiImpl.Swagger))
	// api.RegisterHandlers(r)
	api.RegisterHandlersWithOptions(r, floralApi, api.GinServerOptions{
		ErrorHandler: floralApi.ErrorHandler,
	})

	r.Run()
}
