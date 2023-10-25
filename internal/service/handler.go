package service

import (
	"net/http"

	"github.com/labstack/echo"
	"magic.pathao.com/carta/carta-acm/internal/config"
	"magic.pathao.com/carta/carta-acm/internal/contract"
)

func (a AccountManagerService) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := c.Request().Header.Get("Authorization")
		if authKey == config.GetAppConfig().AppKey {
			return next(c)
		} else {
			return c.JSON(http.StatusForbidden, "You are not authorized")
		}
	}
}

func (a AccountManagerService) GenerateApiKey(c echo.Context) error {

	requestPayload := &contract.GenerateApiKeyReq{}

	ParseRequest(c.Request(), requestPayload)
	a.Logger.Infof("%s", requestPayload)

	err := a.Db.GenerateApiKey(*requestPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, requestPayload)

}
