package db

import (
	"fmt"
	"time"

	"github.com/devShahriar/H"
	"gorm.io/gorm"
	"magic.pathao.com/carta/carta-acm/internal/config"
	"magic.pathao.com/carta/carta-acm/internal/contract"
	"magic.pathao.com/carta/carta-acm/internal/helper"
	logger "magic.pathao.com/data/de-logger"
)

type DbInstance struct {
	Logger logger.Logger
	Db     *gorm.DB
}

var _ DB = &DbInstance{}

func NewDB() DB {

	gormInstance := NewGORMInstance()
	log := logger.GetSimpleLogger("acm_db")
	db := &DbInstance{Logger: log, Db: gormInstance}

	db.RunMigration()
	return db

}

func (d *DbInstance) RunMigration() {
	d.Db.AutoMigrate(
		contract.Users{},
		contract.ApiMeta{},
		contract.Organization{},
		contract.OrganizationMember{},
		contract.Permissions{},
		contract.Roles{},
		contract.RolePermissions{},
	)
}

func (d *DbInstance) Login(req contract.LoginRequest) (contract.LoginResponse, error) {

	existingData := &contract.Users{}

	condition := contract.Users{
		Id: req.Id,
	}

	result := d.Db.Model(&contract.Users{}).Where(condition).First(&existingData)

	if result.Error == nil {
		d.Logger.Infof("Successful Login for user id:%v", req.Id)

		return contract.LoginResponse{
			AccessToken: existingData.AccessToken,
		}, nil
	}

	return contract.LoginResponse{}, result.Error
}

func (d *DbInstance) SignUp(req contract.SignUpRequest) (contract.LoginResponse, error) {

	organizationId, err := d.GetOrganizationId(req.OrgName)

	if err != nil || organizationId == nil {

		ordId, err := d.IngestOrganizationMeta(req)
		if err != nil {
			return contract.LoginResponse{}, err
		}

		if organizationId == nil {
			organizationId = &ordId
		}

		memberMeta := contract.AddMemberReq{
			OrganizationMember: contract.OrganizationMember{
				MemberEmail:    req.Email,
				OrganizationId: *organizationId,
				RoleId:         1,
				CreatedAt:      time.Now(),
			},
		}

		_, err = d.AddOrganizationMember(memberMeta)

		if err != nil {
			return contract.LoginResponse{}, err
		}

	}

	orgMember, err := d.CheckMemberExists(*organizationId, req.Email)
	if err != nil {
		return contract.LoginResponse{}, err
	}

	permissions, err := d.GetPermissionByRole(orgMember.RoleId)
	if err != nil {
		return contract.LoginResponse{}, err
	}

	jwtPayLoad := contract.JwtPayload{
		Id:          req.Id,
		Email:       req.Email,
		OrgId:       *organizationId,
		RoleId:      orgMember.RoleId,
		MemberId:    orgMember.Id,
		Permissions: permissions,
	}

	accessToken, err := helper.GenerateAccessToken(jwtPayLoad)

	userMeta := contract.Users{
		Id:          req.Id,
		UserName:    req.UserName,
		CreatedAt:   time.Now(),
		Designation: req.Designation,
		MemberId:    orgMember.Id,
		AccessToken: accessToken,
	}

	response, err := d.CreateUser(userMeta)
	if err != nil {
		return contract.LoginResponse{}, fmt.Errorf("error while creating new user")
	}

	return response, nil
}

func (d *DbInstance) CreateUser(req contract.Users) (contract.LoginResponse, error) {

	err := d.Db.Model(&contract.Users{}).Create(req).Error
	if err != nil {
		return contract.LoginResponse{}, err
	}

	return contract.LoginResponse{AccessToken: req.AccessToken}, nil
}

func (d *DbInstance) GetPermissionByRole(roleId int) ([]contract.Permissions, error) {

	permissions := []contract.Permissions{}

	err := d.Db.Table("role_permissions as r").
		Joins("join permissions as p on r.permissions_id = p.id").
		Where("r.role_id = ?", roleId).
		Select("p.id, p.name").
		Scan(&permissions).Error
	if err != nil {
		return []contract.Permissions{}, err
	}

	return permissions, nil
}

func (d *DbInstance) CheckMemberExists(orgId string, email string) (*contract.OrganizationMember, error) {

	membersMeta := &contract.OrganizationMember{}

	condition := contract.OrganizationMember{
		MemberEmail:    email,
		OrganizationId: orgId,
	}

	result := d.Db.Model(&contract.OrganizationMember{}).Where(condition).First(membersMeta)
	if result.Error != nil {
		return nil, fmt.Errorf("Member does not exist. Please contact your organization admin")
	}
	return membersMeta, nil
}

func (d *DbInstance) GetOrganizationId(Name string) (*string, error) {

	var organizationId string

	result := d.Db.Model(&contract.Organization{}).Where("name = ?", Name).Select("id").First(&organizationId)
	if result.Error != nil {
		return nil, fmt.Errorf("Organization does not exist")
	}
	return &organizationId, nil

}

func (d *DbInstance) IngestOrganizationMeta(requestPayload contract.SignUpRequest) (string, error) {

	orgData := contract.Organization{
		Id:        helper.GenerateUUID(),
		Name:      requestPayload.OrgName,
		Details:   requestPayload.OrgDetails,
		Link:      requestPayload.OrgLink,
		Email:     requestPayload.OrgEmail,
		CreatedAt: time.Now(),
	}
	result := d.Db.Model(&contract.Organization{}).Create(&orgData)
	if result.Error != nil {
		return "", fmt.Errorf("Couldnot insert into organization table:%v", result.Error)
	}
	return orgData.Id, nil
}

func (d *DbInstance) AddOrganizationMember(req contract.AddMemberReq) (string, error) {

	memberId := helper.GenerateUUID() // generate id
	orgData := contract.OrganizationMember{
		Id:             memberId,
		MemberEmail:    req.MemberEmail,
		OrganizationId: req.OrganizationId,
		RoleId:         req.RoleId,
		CreatedAt:      time.Now(),
	}

	orgData.Status = H.If(orgData.RoleId == 1, "ACTIVE", "PENDING")

	result := d.Db.Model(&contract.OrganizationMember{}).Create(&orgData)
	if result.Error != nil {
		return "", fmt.Errorf("Failed while inserting into organization_member table data:%v error:%v", orgData, result.Error)
	}
	return orgData.Id, nil
}

func (d *DbInstance) GetOrganizations() ([]contract.Organization, error) {

	organizations := []contract.Organization{}

	result := d.Db.Model(&contract.Organization{}).Select("id, name").Find(&organizations)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to fetch organizations list. Error:%v", result.Error)
	}
	return organizations, nil
}

func (d *DbInstance) GetRoles() ([]contract.Roles, error) {

	roles := []contract.Roles{}

	result := d.Db.Model(&contract.Roles{}).Select("id, name").Find(&roles)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to fetch role list. Error:%v", result.Error)
	}
	return roles, nil
}

func (d *DbInstance) GetMembersMeta(orgId string) (*[]contract.GetMembersResp, error) {

	membersMeta := []contract.GetMembersResp{}

	result := d.Db.Table("users as u").
		Select("u.id, u.user_name, u.designation,  om.id, om.member_email, om.role_id, om.status").
		Joins("JOIN organization_members om ON u.member_id = om.id ").
		Joins("JOIN organizations o ON om.organization_id = o.id ").
		Where("om.organization_id = ?", orgId).
		Find(&membersMeta)
	if result.Error != nil {
		return nil, fmt.Errorf("No members found for this organization. Error:%v", result.Error)
	}
	return &membersMeta, nil
}

func (d *DbInstance) UpdateRole(req contract.OrganizationMember) error {

	result := d.Db.Model(&contract.OrganizationMember{}).
		Where("id = ?", req.Id).
		Updates(contract.OrganizationMember{RoleId: req.RoleId})

	if result.Error != nil {
		return fmt.Errorf("This member does not exist. Error:%v", result.Error)
	}
	return nil
}

func (d *DbInstance) UpdateAccessToken(userId string, req contract.OrganizationMember) error {
	// GET New permission for this new roleid

	permissions, err := d.GetPermissionByRole(req.RoleId)
	if err != nil {
		fmt.Errorf("error while fetching permission for memberId:%v roleId:%v", req.Id, req.RoleId)
	}

	user, err := d.GetUser(userId)
	if err != nil {
		return fmt.Errorf("error while fetching user meta for userId:%v memberId:%v roleId:%v", userId, req.Id, req.RoleId)
	}

	jwtPayload, err := helper.ParseJWTToken(user.AccessToken, []byte(config.GetAppConfig().JWTSecretKey))
	if err != nil {
		return fmt.Errorf("error while parsing user access_token for userId:%v memberId:%v roleId:%v", userId, req.Id, req.RoleId)
	}

	d.Logger.Infof("Parsed JWT token : %+v", jwtPayload)

	// newPayLoad := contract.JwtPayload{
	// 	Id:          jwtPayload.Id,
	// 	MemberId:    jwtPayload.MemberId,
	// 	OrgId:       jwtPayload.OrgId,
	// 	RoleId:      jwtPayload.RoleId,
	// 	Permissions: permissions,
	// 	Email:       jwtPayload.Email,
	// }
	jwtPayload.Permissions = permissions
	d.Logger.Infof("New Parsed JWT token : %+v", jwtPayload)

	newAccessToken, err := helper.GenerateAccessToken(*jwtPayload)
	if err != nil {
		return fmt.Errorf("error while updating user access_token for userId:%v memberId:%v roleId:%v", userId, req.Id, req.RoleId)
	}

	updateErr := d.Db.Model(&contract.Users{}).Where("id = ?", userId).Update("access_token", newAccessToken).Error
	if updateErr != nil {
		return fmt.Errorf("error while updating user access_token for userId:%v memberId:%v roleId:%v", userId, req.Id, req.RoleId)
	}

	return nil
}

func (d *DbInstance) DeleteMember(memberId string) error {

	result := d.Db.Model(&contract.OrganizationMember{}).Where("id = ?", memberId).Delete(&contract.OrganizationMember{})
	if result.Error != nil || result.RowsAffected == 0 {
		return fmt.Errorf("This member does not exist. Error:%v", result.Error)
	}
	return nil
}

func (d *DbInstance) GenerateApiKey(token string, req contract.ReqPayload) error {

	data := contract.ApiMeta{
		UserId:      req.UserId,
		ApiKey:      token,
		Limit:       req.Limit,
		Usage:       0,
		Expiry:      req.Expiry,
		ServiceType: req.ServiceType,
		Status:      "ACTIVE",
	}

	condition := contract.ApiMeta{
		UserId:      req.UserId,
		ApiKey:      req.ApiKey,
		ServiceType: req.ServiceType,
	}

	existingApiMeta := contract.ApiMeta{}

	result := d.Db.Model(&contract.ApiMeta{}).Where(condition).Attrs(data).FirstOrCreate(&existingApiMeta)

	if result.Error == nil && result.RowsAffected > 0 {
		d.Logger.Infof("New api meta added for service type:%v, api key:%v and user id:%v",
			data.ServiceType,
			data.ApiKey,
			data.UserId)
	} else {
		d.Logger.Infof("Error: Existing api meta found for service type:%v, api key:%v and user id:%v",
			existingApiMeta.ServiceType,
			existingApiMeta.ApiKey,
			existingApiMeta.UserId)
		return fmt.Errorf("Error: Existing api meta found for service type:%v and user id:%v",
			existingApiMeta.ServiceType,
			existingApiMeta.UserId)
	}

	return result.Error
}

func (d *DbInstance) GetApiKeys(userId string) ([]contract.ApiMeta, error) {

	apiKeys := []contract.ApiMeta{}

	err := d.Db.Model(&contract.ApiMeta{}).Where("user_id = ?", userId).Find(&apiKeys).Error
	if err != nil {
		d.Logger.Error(err.Error())
		return []contract.ApiMeta{}, err
	}

	return apiKeys, nil
}

func (d *DbInstance) UpdateByColumnName(column string, req contract.ReqPayload, data interface{}) error {

	err := d.Db.Model(&contract.ApiMeta{}).
		Where("user_id = ? AND api_key = ? AND service_type = ?", req.UserId, req.ApiKey, req.ServiceType).
		Update(column, data).Error

	if err != nil {
		d.Logger.Error(err.Error())
		return err
	}

	return nil
}

func (d *DbInstance) UpdateApiLimit(req contract.ReqPayload) error {

	//UpdateApiLimit in db
	err := d.UpdateByColumnName(contract.COLUMN_LIMIT, req, req.Limit)
	if err != nil {
		return err
	}

	//Call Auth Service to update limit in redis
	return err
}

func (d *DbInstance) UpdateApiExpiry(req contract.ReqPayload) error {

	//UpdateApiExpiry in db
	err := d.UpdateByColumnName(contract.COLUMN_EXPIRY, req, req.Expiry)
	if err != nil {
		return err
	}
	//Call Auth Service to update expiry in redis
	return err
}

func (d *DbInstance) UpdateApiStatus(req contract.ReqPayload) error {

	//UpdateApiStatus in db
	err := d.UpdateByColumnName(contract.COLUMN_STATUS, req, req.Status)
	if err != nil {
		return err
	}

	//Need to call auth for adding key if status [active] or delete key if status [deactive]
	return err
}

func (d *DbInstance) UpdateUsage(req contract.ReqPayload) error {

	condition := contract.ApiMeta{
		ApiKey:      req.ApiKey,
		ServiceType: req.ServiceType,
		Usage:       req.PreviousUsage,
	}

	result := d.Db.Model(&contract.ApiMeta{}).Where(&condition).UpdateColumn("usage", req.CurrentUsage)
	if result.Error == nil && result.RowsAffected > 0 {
		d.Logger.Infof("Api usage updated for service type:%v", req.ServiceType)
	} else {
		d.Logger.Infof("Error:%v", result.Error)
		return fmt.Errorf("Error: Api usage could not be updated for service type:%v",
			req.ServiceType)
	}
	return result.Error
}

func (d *DbInstance) GetApiMeta(apiKey, svcType string) (*contract.ApiMeta, error) {

	apiMeta := &contract.ApiMeta{}

	condition := contract.ApiMeta{
		ApiKey:      apiKey,
		ServiceType: svcType,
	}

	result := d.Db.Model(&contract.ApiMeta{}).Where(condition).First(&apiMeta)

	if result.Error != nil {
		d.Logger.Infof("Api key not found. Error:%+v", result.Error)
		return nil, fmt.Errorf("Api key not found")
	}

	if !apiMeta.Expiry.After(time.Now()) {
		d.Logger.Infof("Token expired. Api Key:%s, Service Type:%s", apiMeta.ApiKey, apiMeta.ServiceType)
		return nil, fmt.Errorf("Your token has been expired at:%s", apiMeta.Expiry)
	}
	return apiMeta, nil
}

func (d *DbInstance) PopulatePermissionTable() {

}
