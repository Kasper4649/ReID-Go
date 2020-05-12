package main

import (
	. "ReID-Go/controller"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ReIDGroup := r.Group("/api/reid")
	{
		ReIDGroup.POST("/query", Query)
		ReIDGroup.POST("/search", Search)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "呵呵。",
		})
	})

	return r
}