package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var(
	db* gorm.DB
)
func Connect(){
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})  #docs
	d,err := gorm.Open(mysql.Open("akhil:Axlesharma@12/simplerest?charset=utf8&parseTime=true&loc=Local"),&gorm.Config{})
	if err!=nil{
		panic(err)	
	}
	db =d
}
func GetDB() *gorm.DB{
	return db
}