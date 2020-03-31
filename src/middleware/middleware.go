package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func ComputeCostTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	log.Printf("func %s cost: %v", c.HandlerName(), cost)
}
