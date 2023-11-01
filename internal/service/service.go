package service

import (
	"github.com/labstack/echo"
	"magic.pathao.com/carta/carta-acm/internal/db"
	logger "magic.pathao.com/data/de-logger"
)

type AccountManagerService struct {
	Router *echo.Echo
	Logger logger.Logger
	Db     db.DB
}

func NewAccountManagerService(router *echo.Echo, logger logger.Logger, db db.DB) *AccountManagerService {
	return &AccountManagerService{Router: router, Logger: logger, Db: db}
}
