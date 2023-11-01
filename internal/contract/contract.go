package contract

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	COLUMN_LIMIT  = "limit"
	COLUMN_EXPIRY = "expiry"
	COLUMN_STATUS = "status"
)

type LoginRequest struct {
	Id        string    `json:"user_id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type SignUpRequest struct {
	LoginRequest
	OrgName     string `json:"org_name"`
	OrgDetails  string `json:"org_details"`
	OrgLink     string `json:"org_link"`
	OrgEmail    string `json:"org_email"`
	Designation string `json:"designation"`
}

type AddMemberReq struct {
	OrganizationMember
}

type ReqPayload struct {
	UserId        string    `json:"user_id"`
	ApiKey        string    `json:"api_key"`
	ServiceType   string    `json:"service_type"`
	Email         string    `json:"email"`
	Limit         int64     `json:"limit"`
	Expiry        time.Time `json:"expiry"`
	Status        string    `json:"status"`
	CurrentUsage  int64     `json:"current_usage"`
	PreviousUsage int64     `json:"previous_usage"`
}

type Msg struct {
	Message string `json:"message"`
}

type JwtPayload struct {
	Id          string
	Email       string
	RoleId      int
	OrgId       string
	MemberId    string
	Permissions []Permissions
	jwt.RegisteredClaims
}

type GetMembersResp struct {
	OrganizationMember
	Users
}
