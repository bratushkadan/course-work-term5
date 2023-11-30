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
	r.Use(getCors())
	r.Use(middleware.OapiRequestValidator(apiImpl.Swagger))
	api.RegisterHandlersWithOptions(r, floralApi, api.GinServerOptions{
		ErrorHandler: floralApi.ErrorHandler,
	})

	r.Run()
}

func getCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD",
	}

	return cors.New(config)
}
