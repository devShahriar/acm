package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"magic.pathao.com/carta/carta-acm/internal/config"
	"magic.pathao.com/carta/carta-acm/internal/contract"
	"magic.pathao.com/carta/carta-acm/internal/helper"
)

func (a AccountManagerService) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := c.Request().Header.Get("Authorization")
		if authKey == config.GetAppConfig().AppKey {
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, contract.Msg{
				Message: "Unauthorized",
			})
		}
	}
}

func (a AccountManagerService) Login(c echo.Context) error {

	requestPayload := &contract.LoginRequest{}

	ParseRequest(c.Request(), requestPayload)
	a.Logger.Infof("%s", requestPayload)

	response, err := a.Db.Login(*requestPayload)
	if err != nil {
		return c.JSON(http.StatusForbidden, contract.Msg{
			Message: fmt.Sprintf("Login failed. Error: %v", err.Error()),
		})
	}
	return c.JSON(http.StatusOK, response)
}

func (a AccountManagerService) SignUp(c echo.Context) error {

	requestPayload := &contract.SignUpRequest{}

	ParseRequest(c.Request(), requestPayload)
	a.Logger.Infof("%s", requestPayload)

	response, err := a.Db.SignUp(*requestPayload)
	if err != nil {
		return c.JSON(http.StatusForbidden, contract.Msg{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

func (a AccountManagerService) GetOrganizations(c echo.Context) error {

	organizations, err := a.Db.GetOrganizations()
	if err != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, organizations)

}

func (a AccountManagerService) GetRoles(c echo.Context) error {

	roles, err := a.Db.GetRoles()
	if err != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, roles)

}

func (a AccountManagerService) AddOrganizationMember(c echo.Context) error {

	requestPayload := &contract.AddMemberReq{}

	ParseRequest(c.Request(), requestPayload)
	a.Logger.Infof("%s", requestPayload)

	member, err := a.Db.CheckMemberExists(requestPayload.OrganizationId, requestPayload.MemberEmail)

	if member != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{
			Message: "Member already exist",
		})
	}
	if err != nil && member == nil {
		_, err = a.Db.AddOrganizationMember(*requestPayload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, contract.Msg{Message: err.Error()})
		}
	}
	return c.JSON(http.StatusCreated, nil)

}

func (a AccountManagerService) GetMembers(c echo.Context) error {

	orgId := c.QueryParam("organization_id")

	membersList, err := a.Db.GetMembersMeta(orgId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, membersList)
}

func (a AccountManagerService) UpdateRole(c echo.Context) error {

	requestPayload := &contract.OrganizationMember{}

	ParseRequest(c.Request(), requestPayload)
	a.Logger.Infof("%s", requestPayload)

	err := a.Db.UpdateRole(*requestPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, nil)

}

func (a AccountManagerService) DeleteMember(c echo.Context) error {

	memberId := c.QueryParam("id")

	err := a.Db.DeleteMember(memberId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, nil)

}
func (a AccountManagerService) GenerateApiKey(c echo.Context) error {

	requestPayload := &contract.ReqPayload{}

	ParseRequest(c.Request(), requestPayload)
	a.Logger.Infof("%s", requestPayload)

	if !a.Db.IsUserValid(requestPayload.UserId) {
		return c.JSON(http.StatusForbidden, contract.Msg{
			Message: fmt.Sprintf("User id:%s not authorized", requestPayload.UserId),
		})
	}

	token, err := helper.GenerateUniqueAPIKey(requestPayload.UserId, requestPayload.Email, requestPayload.ServiceType)
	if err != nil {
		a.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, contract.Msg{
			Message: err.Error(),
		})
	}

	err = a.Db.GenerateApiKey(token, *requestPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, requestPayload)

}

func (a AccountManagerService) GetApiKeys(c echo.Context) error {

	userId := c.QueryParam("user-id")

	if !a.Db.IsUserValid(userId) {
		return c.JSON(http.StatusForbidden, contract.Msg{
			Message: fmt.Sprintf("User id:%s not authorized", userId),
		})
	}

	response, err := a.Db.GetApiKeys(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)

}

func (a AccountManagerService) UpdateApiMeta(c echo.Context) error {

	requestPayload := &contract.ReqPayload{}
	ParseRequest(c.Request(), requestPayload)

	if !a.Db.IsValidToken(requestPayload.UserId, requestPayload.ApiKey) {
		return c.JSON(http.StatusForbidden, contract.Msg{
			Message: fmt.Sprintf("User id:%s not authorized", requestPayload.UserId),
		})
	}
	fields := c.QueryParam("fields")
	columnArray := strings.Split(fields, ",")

	for _, column := range columnArray {

		switch column {
		case contract.COLUMN_LIMIT:
			err := a.Db.UpdateApiLimit(*requestPayload)
			if err != nil {
				return err
			}
		case contract.COLUMN_EXPIRY:
			err := a.Db.UpdateApiExpiry(*requestPayload)
			if err != nil {
				return err
			}

		case contract.COLUMN_STATUS:
			err := a.Db.UpdateApiStatus(*requestPayload)
			if err != nil {
				return err
			}
		}
	}
	a.Logger.Infof("%+v", *requestPayload)
	return c.JSON(http.StatusOK, contract.Msg{
		Message: "Updated Successfully",
	})

}

func (a AccountManagerService) AddPermission(c echo.Context) error {

	requestPayload := &contract.ReqPayload{}
	ParseRequest(c.Request(), requestPayload)

	if !a.Db.IsValidToken(requestPayload.UserId, requestPayload.ApiKey) {
		return c.JSON(http.StatusForbidden, contract.Msg{
			Message: fmt.Sprintf("User id:%s not authorized", requestPayload.UserId),
		})
	}

	err := a.Db.GenerateApiKey(requestPayload.ApiKey, *requestPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, requestPayload)

}

func (a AccountManagerService) UpdateUsage(c echo.Context) error {

	requestPayload := &contract.ReqPayload{}
	ParseRequest(c.Request(), requestPayload)

	err := a.Db.UpdateUsage(*requestPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, contract.Msg{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, requestPayload)
}

func (a AccountManagerService) GetApiMeta(c echo.Context) error {
	apiKey := c.Request().Header.Get("api_key")
	svcType := c.Request().Header.Get("service_type")

	apiMeta, err := a.Db.GetApiMeta(apiKey, svcType)
	if err != nil {
		return c.JSON(http.StatusForbidden,
			contract.Msg{
				Message: err.Error(),
			})
	}
	return c.JSON(http.StatusOK, apiMeta)
}
