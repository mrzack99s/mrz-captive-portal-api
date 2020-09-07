package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrzack99s/mrz-captive-portal-api/configs"
	"github.com/mrzack99s/mrz-captive-portal-api/structs"
)

func AppendLoginSession(loginSession *structs.ZAuthLoginSession) (err error) {
	if err = configs.DB.Create(loginSession).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLoginSession(loginSession *structs.ZAuthLoginSession, ipAddress string) (err error) {
	configs.DB.Where("ip_address = ?", ipAddress).Delete(loginSession)
	return nil
}

func CountMemberByUsername(username string) int {
	var loginSession *structs.ZAuthLoginSession
	var count int
	configs.DB.Model(&loginSession).Where("username = ?", username).Count(&count)

	return count
}
