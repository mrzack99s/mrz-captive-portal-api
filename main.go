package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/jinzhu/gorm"
	"github.com/mrzack99s/mrz-captive-portal-api/configs"
	"github.com/mrzack99s/mrz-captive-portal-api/routes"
)

var err error

func main() {

	configs.ParseSystemConfig()

	configs.DB, err = gorm.Open("mysql", configs.DbURL(configs.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer configs.DB.Close()

	mode := gin.DebugMode
	if configs.SystemConfig.ZAuth.API.Production {
		mode = gin.ReleaseMode
	}

	gin.SetMode(mode)

	router := routes.SetupRouter()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.Run(configs.SystemConfig.ZAuth.API.Port)

}
