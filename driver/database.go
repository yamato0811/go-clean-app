package driver

import (
	"clean-app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProvidedDatabaseConnection() *gorm.DB {
	fmt.Println(config.GetConfig())
	dbConfig := config.GetConfig().Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/gdb?charset=utf8&parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}