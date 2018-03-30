package controllers

import (
	"todo/app/entities"

	"github.com/jinzhu/gorm"
	gormdb "github.com/revel/modules/orm/gorm/app"
	"github.com/revel/revel"
)

type ApplicationController struct {
	*revel.Controller
	DB *gorm.DB
}

func (c *ApplicationController) setDB() revel.Result {
	c.DB = gormdb.DB
	c.DB.LogMode(true)
	return nil
}

func initializeDB() {
	gormdb.InitDB()
	gormdb.DB.AutoMigrate(&entities.Todo{})

	revel.InterceptMethod((*ApplicationController).setDB, revel.BEFORE)
}

func init() {
	revel.OnAppStart(initializeDB)
}
