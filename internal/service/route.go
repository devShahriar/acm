package service

import "magic.pathao.com/carta/carta-acm/internal/contract"

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

	acm.Router.POST("/v1/login", acm.Auth(Middleware(contract.PermDefault, acm.Login, false)))
	acm.Router.POST("/v1/signup", acm.Auth(Middleware(contract.PermDefault, acm.SignUp, false)))

	acm.Router.POST("/v1/generate/api-key", acm.Auth(Middleware(contract.PermGenerateApiKey, acm.GenerateApiKey, true)))
	acm.Router.GET("/v1/api-keys", acm.Auth(Middleware(contract.PermGetApiKeys, acm.GetApiKeys, true)))

	acm.Router.POST("/v1/update-api-meta", acm.Auth(Middleware(contract.PermUpdateApiKeyMeta, acm.UpdateApiMeta, true)))
	acm.Router.POST("/v1/add-permission", acm.Auth(Middleware(contract.PermAddApiPermission, acm.AddPermission, true)))

	acm.Router.GET("/v1/get-organizations", acm.Auth(Middleware(contract.PermDefault, acm.GetOrganizations, false)))
	acm.Router.GET("/v1/roles", acm.Auth(Middleware(contract.PermDefault, acm.GetRoles, false)))

	acm.Router.PATCH("/v1/update-role", acm.Auth(Middleware(contract.PermUpdateRole, acm.UpdateRole, true)))
	acm.Router.GET("/v1/members", acm.Auth(Middleware(contract.PermGetMembers, acm.GetMembers, true)))
	acm.Router.POST("/v1/members/add", acm.Auth(Middleware(contract.PermAddMembers, acm.AddOrganizationMember, true)))
	acm.Router.DELETE("/v1/delete/member", acm.Auth(Middleware(contract.PermDeleteMembers, acm.DeleteMember, true)))

	acm.Router.POST("/v1/update-usage", acm.Auth(Middleware(contract.PermDefault, acm.UpdateUsage, false)))
	acm.Router.GET("/v1/get-api-meta", acm.Auth(Middleware(contract.PermDefault, acm.GetApiMeta, false)))
}
