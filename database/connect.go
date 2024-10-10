package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectMySQL() *gorm.DB {
	db, err := gorm.Open("mysql", "root:6437419@tcp(localhost:3306)"+"/software_data?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
