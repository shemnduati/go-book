package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/dialect/mysql"
)

var (
	db*gorm.DB
)


func connect(){
	d, err := gorm.Open("mysql", "store?charset=utf8$parseTime=Trues&Location=Local")
	 if err != nil{
		 panic(err)
	 }
	 db = d
}

func GetDB() *gorm.DB(
	return db
)
