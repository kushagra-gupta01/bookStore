package config
import(
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
)
var(
	db* gorm.DB
)
func connect(){
	var err error	//gpt
	d,err = gorm.Open("mysql","akhil:Axlesharma@12/simplerest?charset=utf8&parseTime=true&loc=Local")
	if err!=nil{
		panic(err)	
	}
	db =d
}
func getDB() *gorm.DB{
	return db
}