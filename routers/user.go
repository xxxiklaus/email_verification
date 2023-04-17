package routers

import (
	"context"
	"xxxiklaus/email-verification/controllers"
	"xxxiklaus/email-verification/middles"
	"xxxiklaus/email-verification/services"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(ctx context.Context, userService services.UserServices, router *gin.RouterGroup) {
	controllersImpl := controllers.NewUserControllersImpl(ctx, userService)
	router.POST("/signup", controllersImpl.SignUp)
	authMiddleware, err := middles.InitAuthMiddlewares(controllersImpl)
	if err != nil {
		panic(err)
	}
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/refresh_token", authMiddleware.RefreshHandler)
	router.GET("/verify_email", controllersImpl.VerifyEmail)

	{
		auth := router.Use(authMiddleware.MiddlewareFunc())
		auth.POST("/logout", authMiddleware.LogoutHandler)
		auth.GET("/me", controllersImpl.Info)
		auth.POST("/send_email", controllersImpl.SendEmail)
	}
}
