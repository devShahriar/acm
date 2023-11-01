package db

import (
	"magic.pathao.com/carta/carta-acm/internal/config"
	"magic.pathao.com/carta/carta-acm/internal/contract"
	"magic.pathao.com/carta/carta-acm/internal/helper"
)

func (d *DbInstance) CheckPermission() {}

func (d *DbInstance) IsUserValid(userId string) bool {

	users := &contract.Users{}

	err := d.Db.Model(&contract.Users{}).Where("id = ?", userId).First(users).Error
	if err != nil || users == nil {
		return false
	}
	return true
}

func (d *DbInstance) CheckValidUser(token string) (*contract.JwtPayload, error) {

	secretKey := config.GetAppConfig().JWTSecretKey
	helper.ParseJWTToken(token, []byte(secretKey))
	return nil, nil
}

func (d *DbInstance) IsApiKeyValid(apiKey string) bool {

	apiMeta := &contract.ApiMeta{}

	err := d.Db.Model(&contract.ApiMeta{}).Where("api_key = ?", apiKey).First(apiMeta).Error
	if err != nil || apiMeta == nil {
		return false
	}

	return true

}

func (d *DbInstance) IsValidToken(userId, apiKey string) bool {

	if d.IsApiKeyValid(apiKey) && d.IsUserValid(userId) {
		return true
	} else {
		return false
	}
}
