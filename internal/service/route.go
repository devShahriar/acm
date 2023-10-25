package service

func RegisterApi(acm *AccountManagerService) {
	acm.Router.POST("/v1/generate/api-key", acm.Auth(acm.GenerateApiKey))

}
