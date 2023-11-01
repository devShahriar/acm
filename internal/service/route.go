package service

/*
addMembers
getMembers
deleteMembers

updateRole
generateApiKey
getApiKeys
updateApiKeyMeta
addApiPermission
*/

func RegisterApi(acm *AccountManagerService) {

	acm.Router.POST("/v1/login", acm.Auth(acm.Login))
	acm.Router.POST("/v1/signup", acm.Auth(acm.SignUp))

	acm.Router.POST("/v1/generate/api-key", acm.Auth(acm.GenerateApiKey))
	acm.Router.GET("/v1/api-keys", acm.Auth(acm.GetApiKeys))
	acm.Router.POST("/v1/update-api-meta", acm.Auth(acm.UpdateApiMeta))
	acm.Router.POST("/v1/add-permission", acm.Auth(acm.AddPermission))

	acm.Router.GET("/v1/get-organizations", acm.Auth(acm.GetOrganizations))
	acm.Router.GET("/v1/roles", acm.Auth(acm.GetRoles))
	acm.Router.PATCH("/v1/update-role", acm.Auth(acm.UpdateRole))

	acm.Router.GET("/v1/members", acm.Auth(acm.GetMembers))
	acm.Router.POST("/v1/members/add", acm.Auth(acm.AddOrganizationMember))
	acm.Router.DELETE("/v1/delete/member", acm.Auth(acm.DeleteMember))

	acm.Router.POST("/v1/update-usage", acm.Auth(acm.UpdateUsage))
	acm.Router.GET("/v1/get-api-meta", acm.Auth(acm.GetApiMeta))
}
