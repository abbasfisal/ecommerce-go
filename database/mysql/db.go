package mysql

import (
	"fmt"
	"github.com/abbasfisal/ecommerce-go/config"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitClient(config *config.Config) error {

	var err error

	dsn := config.Db.UserName + ":" + config.Db.Password + "@tcp" + "(" + config.Db.Host + ":" + config.Db.Port + ")/" + config.Db.Name + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := dbClient.DB()
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Get() *gorm.DB {
	return dbClient
}

func Close() {
	connection, _ := dbClient.DB()
	err := connection.Close()
	if err != nil {
		return
	}
}
