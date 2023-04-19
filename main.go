package main

import (
	"context"
	"log"
	"xxxiklaus/email-verification/initialize"
	"xxxiklaus/email-verification/middles"
	"xxxiklaus/email-verification/routers"
	"xxxiklaus/email-verification/services"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

var ctx context.Context

//swagger 注释
// @title Email Verification API
// @version 1.0.0
// @description Create and verify email addresses
// @host 127.0.0.1:8080
// @BasePath /api

//JWT 注释
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	ctx = context.Background()
	//config load
	initialize.InitConfig(".")
	//mongodb client
	initialize.InitClient(ctx)
	defer initialize.CloseClient(ctx)
	//di services
	services.InitUserCollection()
	userService := services.NewUserServicesImpl(services.GetUserCollection(), ctx)
	//gin
	server := gin.Default()
	server.Use(requestid.New())
	server.Use(middles.AddCors()) //跨域中间件

	api := server.Group("/api")
	routers.InitCommonRouter(api)
	routers.InitUserRouter(ctx, userService, api)
	routers.InitSwaggerRouter(server)

	//swagger url
	log.Println("swagger url:" + initialize.GetConfig().BaseUrl + "/swagger/index.html")
	log.Fatal(server.Run(":" + initialize.GetConfig().Port))
}
