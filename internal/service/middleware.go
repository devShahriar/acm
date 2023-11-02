package service

import (
	"net/http"

	"github.com/devShahriar/H"
	"github.com/labstack/echo"
	"magic.pathao.com/carta/carta-acm/internal/config"
	"magic.pathao.com/carta/carta-acm/internal/contract"
)

type MiddlewareNamedHandler struct {
	Perm         string
	Func         echo.HandlerFunc
	ValidateUser bool
}

func Middleware(perm string, handler echo.HandlerFunc, validateUser bool) MiddlewareNamedHandler {
	return MiddlewareNamedHandler{
		Perm:         perm,
		Func:         handler,
		ValidateUser: validateUser,
	}
}

func (a AccountManagerService) Auth(next MiddlewareNamedHandler) echo.HandlerFunc {

	return func(c echo.Context) error {
		//Authorization
		appKey := c.Request().Header.Get("Appkey")
		validUser := false

		verifiedAppKey := H.If(appKey == config.GetAppConfig().AppKey, true, false)

		if next.ValidateUser {

			token := ParseToken(c.Request(), "Authorization")

			jwtPayload, err := a.Db.CheckValidUser(token)
			if err != nil || jwtPayload == nil {
				a.Logger.Errorf("CheckValidateUser Error:%v", err)
				validUser = false
			}

			if jwtPayload != nil {

				a.Logger.Infof("Setting user_id:%v in context", jwtPayload.Id)
				c.Set("user_id", jwtPayload.Id)

				allowed := a.CheckPermission(jwtPayload.Permissions, next.Perm)
				if allowed {
					a.Logger.Infof("user allowed for func:%v", next.Perm)
					validUser = true
				}
			}
		} else {
			a.Logger.Infof("Skipping user validation for func: %v", next.Perm)
			validUser = true
		}

		if validUser && verifiedAppKey {
			a.Logger.Infof("validUser: %v verifiedAppKey: %v func: %v", validUser, verifiedAppKey, next.Perm)
			return next.Func(c)
		}

		return c.JSON(http.StatusForbidden, contract.Msg{Message: "access denied"})
	}
}
