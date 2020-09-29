package model

import (
	"fmt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Dependency struct {
	Db *gorm.DB
}

type Ctx struct {
	echo.Context
	*Dependency
	HasError error
}

func (ctx *Ctx) ReturnError(err error) bool {
	if err != nil {
		ctx.JSON(200, M.SX{
			`error`: err.Error(),
		})
		ctx.HasError = err
	}
	return ctx.HasError != nil
}

func (ctx *Ctx) IsError(err error, description string) bool {
	if L.IsError(err, description) {
		ctx.JSON(200, M.SX{
			`error`: description,
		})
		ctx.HasError = err
	}
	return ctx.HasError != nil
}

func InitDependency() *Dependency {
	server := &Dependency{}
	var err error
	usr := os.Getenv(`POSTGRES_USER`)
	pwd := os.Getenv(`POSTGRES_PASSWORD`)
	db := os.Getenv(`POSTGRES_DB`)
	host := os.Getenv(`DB_HOST`)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, usr, pwd, db)
	server.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	L.PanicIf(err, `failed initialize postgres connection`)
	err = server.Db.AutoMigrate(&Product{})
	L.PanicIf(err, `failed automigration`)
	return server
}

func (s *Dependency) Wrap(handler func(*Ctx)) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx2 := Ctx{
			Context:    ctx,
			Dependency: s,
		}
		handler(&ctx2)
		return ctx2.HasError
	}
}

func (s *Dependency) Close() {
}
