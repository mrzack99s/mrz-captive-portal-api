package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-captive-portal-api/controllers"
)

var rootPath = "/zauth/v2beta/authenticate"

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORS)

	r.POST(rootPath+"/exit", controllers.Logout)
	r.POST(rootPath+"/request", controllers.GetAuthenticate)
	r.POST(rootPath+"/requestEAP", controllers.GetEAPAuthenticate)

	return r
}

func CORS(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
