package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-captive-portal-api/configs"
	"github.com/mrzack99s/mrz-captive-portal-api/runtime"
	"github.com/mrzack99s/mrz-captive-portal-api/structs"
)

func Logout(c *gin.Context) {

	var input structs.ZAuthIPAddress
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var loginSession structs.ZAuthLoginSession
	checkDelete := DeleteLoginSession(&loginSession, input.IPAddress)

	message := map[string]interface{}{
		"IPAddress": input.IPAddress,
		"ShareKey":  configs.SystemConfig.ZAuth.Operator.ShareKey,
	}

	runtime.Run("logout", message)

	if checkDelete != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":    true,
			"ipAddress": input.IPAddress,
		})
	}

}
