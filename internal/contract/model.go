package contract

import "time"

type Users struct {
	Id           string
	CreatedAt    time.Time
	Email        string
	PasswordHash string
	UserName     string
	ApiMeta      []ApiMeta `gorm:"foreignKey:UserId;references:Id"`
}

type ApiMeta struct {
	UserId      string
	ApiKey      string
	ServiceType string
	Limit       int64 `gorm:"type:bigint"`
	Usage       int64 `gorm:"type:bigint"`
	Expiry      time.Time
	Status      string
}
