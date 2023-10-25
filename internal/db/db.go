package db

import (
	"gorm.io/gorm"
	"magic.pathao.com/carta/carta-acm/internal/contract"
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
	)
}

func (d *DbInstance) GenerateApiKey(req contract.GenerateApiKeyReq) error {
	//Need to check if the user_id is valid

	//Generate api key
	api_key := "kokoajfiejiafe"

	data := contract.ApiMeta{
		UserId:      req.UserId,
		ApiKey:      api_key,
		Limit:       req.Limit,
		Usage:       0,
		Expiry:      req.Expiry,
		ServiceType: req.ServiceType,
		Status:      "ACTIVE",
	}
	err := d.Db.Create(data).Error

	if err != nil {
		d.Logger.Error(err.Error())
	}

	return err
}
