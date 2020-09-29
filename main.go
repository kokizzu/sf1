package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"os"
	"sf1/controller/cProductManager"
	"sf1/model"

	_ "sf1/docs"
)

// @title SF1
// @version 1.0
// @description

// @contact.name Kiswono Prayogo

// @host 127.0.0.1:9090
// @BasePath /
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// init dependencies
	dep := model.InitDependency()

	// documentation: http://127.0.0.1:9090/swagger/index.html
	e.GET(`/swagger/*`, echoSwagger.WrapHandler)

	// APIs
	e.File("/", "swagger.json")
	e.GET(`/product_manager/product/list/:limit/:offset`, dep.Wrap(cProductManager.Product_List))
	e.GET(`/product_manager/product/by-id/:productId`, dep.Wrap(cProductManager.Product_ById))
	e.POST(`/product_manager/product/modify`, dep.Wrap(cProductManager.Product_Upsert)) // includes delete, insert, update

	serviceAddr := os.Getenv(`SERVICE_ADDR`)
	e.Logger.Fatal(e.Start(serviceAddr))
}
