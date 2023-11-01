package test

import (
	"fmt"
	"testing"

	"magic.pathao.com/carta/carta-acm/internal/contract"
	"magic.pathao.com/carta/carta-acm/internal/helper"
)

func TestGenerateApiKey(t *testing.T) {
	apiKey, err := helper.GenerateUniqueAPIKey(
		"122344",
		"ramisa.mom@pathoa.com",
		"search",
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(apiKey)
}

func TestGenerateAccessToken(t *testing.T) {

	permissions := []contract.Permissions{
		{
			Id:   1,
			Name: "getApiKeys",
		},
		{
			Id:   2,
			Name: "generateApiKeys",
		},
		{
			Id:   3,
			Name: "addMembers",
		},
	}

	payLoad := contract.JwtPayload{
		Id:          "12321",
		Email:       "example@gamil.com",
		RoleId:      1,
		OrgId:       "jihlojij",
		MemberId:    "123",
		Permissions: permissions,
	}

	token, err := helper.GenerateAccessToken(payLoad)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("token:", token)
}
