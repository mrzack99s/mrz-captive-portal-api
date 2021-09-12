package main

import (
	"fmt"
	"os"
	"strconv"
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

	val, present := os.LookupEnv("RADIUS_URL")
	if present {
		configs.SystemConfig.ZAuth.Radius.HostURL = val
	}

	val, present = os.LookupEnv("RADIUS_SECRET")
	if present {
		configs.SystemConfig.ZAuth.Radius.Secret = val
	}

	val, present = os.LookupEnv("MySQL_HOSTNAME")
	if present {
		configs.SystemConfig.ZAuth.MySQL.HostIP = val
	}

	val, present = os.LookupEnv("MySQL_PORT")
	if present {
		intVal, _ := strconv.Atoi(val)
		configs.SystemConfig.ZAuth.MySQL.Port = intVal
	}

	val, present = os.LookupEnv("MySQL_USRNAME")
	if present {
		configs.SystemConfig.ZAuth.MySQL.Username = val
	}

	val, present = os.LookupEnv("MySQL_PASSWORD")
	if present {
		configs.SystemConfig.ZAuth.MySQL.Password = val
	}

	val, present = os.LookupEnv("MySQL_DB")
	if present {
		configs.SystemConfig.ZAuth.MySQL.DBName = val
	}

	val, present = os.LookupEnv("OPERATOR_SHAREKEY")
	if present {
		configs.SystemConfig.ZAuth.Operator.ShareKey = val
	}

	val, present = os.LookupEnv("OPERATOR_URL")
	if present {
		configs.SystemConfig.ZAuth.Operator.HostURL = val
	}

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
