package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Timur_Salimov_97_TS@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
