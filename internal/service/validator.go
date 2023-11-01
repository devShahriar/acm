package service

import (
	"errors"

	"magic.pathao.com/carta/carta-acm/internal/contract"
)

func (a AccountManagerService) ValidateUpdateRequest(req contract.ReqPayload) error {
	if req.UserId == "" || req.ApiKey == "" || req.ServiceType == "" {
		a.Logger.Error("User Id or Api Key is empty")
		return errors.New("User Id or Api Key is empty")
	}
	return nil
}
