package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Mysql(hostname string, port int, username string, password string, dbname string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			username, password, hostname, port, dbname))
	if err != nil {
		return nil, err
	}
	DB = db
	return DB, nil
}

func InitDB() *gorm.DB {
	// driverName := "mysql"
	// host := "127.0.0.1"
	// port := "3306"
	// database := "lg-gin-vue"
	// username := "lg-gin-vue"
	// passwd := "lg-gin-vue"
	// charset := "utf-8"
	// args := fmt.Sprintf("%s:%s@tcp(%s:%s)?charset=%s&parseTime=true",
	// 	username,
	// 	passwd,
	// 	host,
	// 	port,
	// 	database,
	// 	charset,
	// )
	db, err := gorm.Open("mysql", "lg-gin-vue:lg-gin-vue@(127.0.0.1)/lg-gin-vue?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
	return db
}

func Run() {
	db := InitDB()
	defer db.Close()
}
