package main

import (
	"github.com/labstack/echo"
	"magic.pathao.com/carta/carta-acm/internal/db"
	"magic.pathao.com/carta/carta-acm/internal/service"
	logger "magic.pathao.com/data/de-logger"
)

func main() {

	logger := logger.GetSimpleLogger("acm")
	db := db.NewDB()

	router := echo.New()
	svc := service.NewAccountManagerService(router, logger, db)
	service.RegisterApi(svc)
	router.Logger.Fatal(router.Start(":8080"))

}
