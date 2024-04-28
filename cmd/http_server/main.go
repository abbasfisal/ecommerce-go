package main

import (
	"github.com/abbasfisal/ecommerce-go/cmd/http_server/routes"
	"github.com/abbasfisal/ecommerce-go/internal/config"
	"github.com/abbasfisal/ecommerce-go/internal/database/mysql"
	"log"
)

func main() {
	cnf := config.NewConfig()

	err := mysql.InitClient(cnf)
	if err != nil {
		log.Fatal("database : ", err)
	}

	routes.NewServerApis(mysql.Get(), cnf).Run()

}
