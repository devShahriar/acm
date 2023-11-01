package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"magic.pathao.com/carta/carta-acm/internal/config"
)

func NewGORMInstance() *gorm.DB {

	conf := config.GetAppConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.DbConf.Host,
		conf.DbConf.User,
		conf.DbConf.Password,
		conf.DbConf.DbName,
		conf.DbConf.Port,
	)

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			//If a query takes more time than SlowQueryThreshold then gorm will alert in log pointing the slow query
			SlowThreshold:             time.Duration(conf.DbConf.SlowQueryThreshold) * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,                                                      // Log level
			IgnoreRecordNotFoundError: false,                                                            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                                                             // Enable colorized logs
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		panic("failed to connect to the database")
	}

	return db

}
