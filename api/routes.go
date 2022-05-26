package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/samwhf/backendTest/docs"
	"github.com/samwhf/backendTest/handlers/user"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetUpRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	//user routes
	v1.GET("/user/:id", user.Get)
	v1.POST("/user", user.Create)
	v1.PUT("/user/:id", user.Update)
	v1.DELETE("/user/:id", user.Delete)
	// Swagger Endpoint
	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
