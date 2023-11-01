package contract

import (
	"time"
)

type Users struct {
	Id          string
	CreatedAt   time.Time
	UserName    string
	AccessToken string
	MemberId    string
	Designation string
	ApiMeta     []ApiMeta `gorm:"foreignKey:UserId;references:Id"`
}

type ApiMeta struct {
	UserId      string
	ApiKey      string
	ApiKeyId    string
	ServiceType string
	Limit       int64 `gorm:"type:bigint"`
	Usage       int64 `gorm:"type:bigint"`
	Expiry      time.Time
	Status      string
}

type Organization struct {
	Id                 string
	Name               string
	Details            string
	Email              string
	Link               string
	CreatedAt          time.Time
	OrganizationMember []OrganizationMember `gorm:"foreignKey:OrganizationId;references:Id"`
}
type OrganizationMember struct {
	Id             string    `json:"id"`
	OrganizationId string    `json:"organization_id"`
	MemberEmail    string    `json:"email"`
	RoleId         int       `json:"role_id"`
	CreatedAt      time.Time `json:"created_at"`
	Status         string    `json:"status"`
}

type Permissions struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Roles struct {
	Id   int
	Name string
}

type RolePermissions struct {
	RoleId        int
	PermissionsId int
}
