package main

import (
	. "ReID-Go/src/controller"
	. "ReID-Go/src/middleware"
	"ReID-Go/src/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var config util.Conf

func main() {

	r := gin.Default()
	r.Use(ComputeCostTime)
	config.GetConf()

	ReIDGroup := r.Group("/reid")
	{
		ReIDGroup.POST("/query", Query)
		ReIDGroup.POST("/search", Search)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "呵呵。",
		})
	})

	err := r.Run(":" + config.GinServePort)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
