package contract

import "time"

type GenerateApiKeyReq struct {
	UserId      string    `json:"user_id"`
	ServiceType string    `json:"service_type"`
	Limit       int64     `json:"limit"`
	Expiry      time.Time `json:"expiry"`
}
