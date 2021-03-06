package main

import (
	. "ReID-Go/middleware"
	"ReID-Go/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title ReID API
// @version 1.0
// @description This is a ReID's server API.
// @termsOfService https://:).moe

// @contact.name Kasper
// @contact.url https://kasper.moe
// @contact.email me@Kasper.moe

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api


func main() {
	util.InitConfig()

	r := gin.Default()
	r.Use(ComputeCostTime, CORSMiddleware())
	r = CollectRoute(r)

	panic(r.Run(":" + viper.GetString("server.port")))
}