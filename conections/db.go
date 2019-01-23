package conections

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	DB *gorm.DB
)

func ConnectDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:123123@tcp(www.ydl.com:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
}
