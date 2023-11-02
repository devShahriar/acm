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

func (a AccountManagerService) CheckPermission(permissions []contract.Permissions, perm string) bool {

	allowed := false

	if len(permissions) > 0 {
		for _, v := range permissions {
			if v.Name == perm {
				allowed = true
				a.Logger.Infof("Permission allowed for %v:", perm)
				return allowed
			}
		}
	}
	a.Logger.Infof("Permission not allowed for %v:", perm)
	return allowed
}
