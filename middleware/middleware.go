package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func ComputeCostTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	log.Printf("func %s cost: %v", c.HandlerName(), cost)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			ctx.Next()
		}
	}
}
