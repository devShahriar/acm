package db

import "magic.pathao.com/carta/carta-acm/internal/contract"

type DB interface {
	Login(req contract.LoginRequest) (contract.LoginResponse, error)
	SignUp(req contract.SignUpRequest) (contract.LoginResponse, error)
	CreateUser(req contract.Users) (contract.LoginResponse, error)

	GetOrganizationId(Name string) (*string, error)
	IngestOrganizationMeta(requestPayload contract.SignUpRequest) (string, error)
	GetOrganizations() ([]contract.Organization, error)

	AddOrganizationMember(req contract.AddMemberReq) (string, error)
	CheckMemberExists(orgId string, email string) (*contract.OrganizationMember, error)
	GetMembersMeta(orgId string) (*[]contract.GetMembersResp, error)
	DeleteMember(memberId string) error

	GetRoles() ([]contract.Roles, error)
	UpdateRole(req contract.OrganizationMember) error

	GenerateApiKey(token string, req contract.ReqPayload) error
	GetApiKeys(userId string) ([]contract.ApiMeta, error)

	UpdateApiLimit(req contract.ReqPayload) error
	UpdateApiStatus(req contract.ReqPayload) error
	UpdateApiExpiry(req contract.ReqPayload) error
	UpdateByColumnName(column string, req contract.ReqPayload, data interface{}) error
	UpdateAccessToken(userId string, req contract.OrganizationMember) error

	GetUser(userId string) (*contract.Users, error)
	CheckValidUser(token string) (*contract.JwtPayload, error)

	UpdateUsage(req contract.ReqPayload) error
	GetApiMeta(apiKey, svcType string) (*contract.ApiMeta, error)
}
