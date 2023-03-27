package data

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){
	
	DB, _ = gorm.Open(sqlite.Open("./gorm.sqlite"), &gorm.Config{})
	DB.AutoMigrate(&Employee {})
	var antal int64
	DB.Model(Employee{}).Count(&antal)

	if antal == 0{
		DB.Create(&Employee{Age: 50, Name: "Steffe", City: "Holm"})
		DB.Create(&Employee{Age: 35, Name: "Maria", City: "Törsk"})
	}
}