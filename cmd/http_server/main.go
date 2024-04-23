package main

import (
	"github.com/abbasfisal/ecommerce-go/config"
	"github.com/abbasfisal/ecommerce-go/database/mysql"
	"log"
)

func main() {
	cnf := config.NewConfig()

	err := mysql.InitClient(cnf)
	if err != nil {
		log.Fatal("database : ", err)
	}

}
