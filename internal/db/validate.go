package db

import (
	"fmt"

	"magic.pathao.com/carta/carta-acm/internal/config"
	"magic.pathao.com/carta/carta-acm/internal/contract"
	"magic.pathao.com/carta/carta-acm/internal/helper"
)

func (d *DbInstance) GetUser(userId string) (*contract.Users, error) {

	users := &contract.Users{}

	err := d.Db.Model(&contract.Users{}).Where("id = ?", userId).First(users).Error
	if err != nil || users == nil {
		return nil, err
	}
	return users, nil
}

func (d *DbInstance) CheckValidUser(token string) (*contract.JwtPayload, error) {

	secretKey := config.GetAppConfig().JWTSecretKey
	jwtPayload, err := helper.ParseJWTToken(token, []byte(secretKey))

	if err != nil || jwtPayload == nil {
		return nil, fmt.Errorf("invalid user access token")
	}

	user, err := d.GetUser(jwtPayload.Id)

	if err != nil {
		return nil, fmt.Errorf("user doesn't exist in user table. id:", jwtPayload.Id)
	}

	if user.AccessToken != token {
		return nil, fmt.Errorf("invalid user token user_id:%v token:%v", jwtPayload.Id, token)
	}

	return jwtPayload, nil
}
