package controllers

import (
	except "github.com/mrzack99s/mrz-captive-portal-api/exceptions"

	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrzack99s/mrz-captive-portal-api/configs"
	"github.com/mrzack99s/mrz-captive-portal-api/runtime"
	"github.com/mrzack99s/mrz-captive-portal-api/structs"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/vendors/wispr"
)

//GetAuthenticate ...
func GetAuthenticate(c *gin.Context) {

	var input structs.ZAuthCheck
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hostURL := os.Getenv("HOST_URL")
	if hostURL == "" {
		hostURL = "10.254.0.12:1812"
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "rtnsecret"
	}

	var result structs.AuthenticateResult
	except.Block{
		Try: func() {

			packet := radius.New(radius.CodeAccessRequest, []byte(secret))
			rfc2865.UserName_SetString(packet, input.Username)
			rfc2865.UserPassword_SetString(packet, input.Password)
			response, err := radius.Exchange(context.Background(), packet, hostURL)

			if err != nil {
				log.Println(err)
			} else {

				var replyPass bool
				replyPass = false

				if response.Code == radius.CodeAccessAccept {

					count := CountMemberByUsername(input.Username)
					if input.Username == "temporary" {
						if count <= 200 {
							replyPass = true
							result.Other = ""
						} else {
							replyPass = false
							result.Other = "Login session more than 200 devices"
						}
					} else {
						if count <= 3 {
							replyPass = true
							result.Other = ""
						} else {
							replyPass = false
							result.Other = "Login session more than 3 devices"
						}
					}

				} else {
					replyPass = false
					result.Other = "Authentication failed"
				}

				if replyPass {
					wisprDlSpeed := wispr.WISPrBandwidthMaxDown_Get(response)
					wisprUpSpeed := wispr.WISPrBandwidthMaxUp_Get(response)
					result.IPAddress = input.IPAddress
					result.Status = true
					result.Username = input.Username
					result.WISPrBandwidthMaxDown = uint32(wisprDlSpeed)
					result.WISPrBandwidthMaxUp = uint32(wisprUpSpeed)

					loginSession := structs.ZAuthLoginSession{IPAddress: input.IPAddress, Username: input.Username}
					checkAppend := AppendLoginSession(&loginSession)

					if checkAppend != nil {
						c.AbortWithStatus(http.StatusInternalServerError)
					}

					message := map[string]interface{}{
						"IPAddress": input.IPAddress,
						"DlSpeed":   wisprDlSpeed,
						"UpSpeed":   wisprUpSpeed,
						"ShareKey":  configs.SystemConfig.ZAuth.Operator.ShareKey,
					}

					runtime.Run("allowNet", message)

				} else {
					result.IPAddress = input.IPAddress
					result.Status = false
					result.Username = input.Username
					result.WISPrBandwidthMaxDown = 1
					result.WISPrBandwidthMaxUp = 1
				}
			}
		},
		Catch: func(e except.Exception) {
			fmt.Printf("Caught %v\n", e)
			result.IPAddress = input.IPAddress
			result.Status = false
			result.Username = input.Username
			result.WISPrBandwidthMaxDown = 1
			result.WISPrBandwidthMaxUp = 1
		},
	}.Do()

	c.JSON(http.StatusOK, gin.H{
		"status":    result.Status,
		"ipAddress": result.IPAddress,
		"username":  result.Username,
		"dlSpeed":   result.WISPrBandwidthMaxDown,
		"upSpeed":   result.WISPrBandwidthMaxUp,
		"other":     result.Other,
	})
	return

}

func GetEAPAuthenticate(c *gin.Context) {

	var input structs.ZAuthEAP
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	type Result struct {
		Attribute string
		Value     uint32
	}

	var queryResult []Result
	var result structs.AuthenticateResult

	var replyPass bool
	replyPass = false
	except.Block{
		Try: func() {
			configs.DB.Table("radgroupreply").Select("radgroupreply.attribute, radgroupreply.value").Joins("JOIN groups ON radgroupreply.groupname = groups.groupname").Joins("JOIN radusergroup ON radusergroup.groupname = groups.groupname").Where("radusergroup.username = ?", input.Username).Find(&queryResult)

			count := CountMemberByUsername(input.Username)
			if input.Username == "temporary" {
				if count <= 200 {
					replyPass = true
					result.Other = ""
				} else {
					replyPass = false
					result.Other = "Login session more than 200 devices"
				}
			} else {
				if count <= 3 {
					replyPass = true
					result.Other = ""
				} else {
					replyPass = false
					result.Other = "Login session more than 3 devices"
				}
			}

		},
		Catch: func(e except.Exception) {
			replyPass = false
			result.Other = "Authorize failed"
		},
	}.Do()

	if replyPass {
		wisprDlSpeed := queryResult[0].Value
		wisprUpSpeed := queryResult[1].Value
		result.IPAddress = input.IPAddress
		result.Status = true
		result.Username = input.Username
		result.WISPrBandwidthMaxDown = uint32(wisprDlSpeed)
		result.WISPrBandwidthMaxUp = uint32(wisprUpSpeed)

		loginSession := structs.ZAuthLoginSession{IPAddress: input.IPAddress, Username: input.Username}
		checkAppend := AppendLoginSession(&loginSession)

		if checkAppend != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		message := map[string]interface{}{
			"IPAddress": input.IPAddress,
			"DlSpeed":   wisprDlSpeed,
			"UpSpeed":   wisprUpSpeed,
			"ShareKey":  configs.SystemConfig.ZAuth.Operator.ShareKey,
		}

		runtime.Run("allowNet", message)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    result.Status,
		"ipAddress": result.IPAddress,
		"username":  result.Username,
		"dlSpeed":   result.WISPrBandwidthMaxDown,
		"upSpeed":   result.WISPrBandwidthMaxUp,
		"other":     result.Other,
	})
	return
}
