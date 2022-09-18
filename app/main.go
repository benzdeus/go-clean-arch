package main

import (
	"go-clean-arch/configs"
	"go-clean-arch/modules/servers"
	mysqlpkg "go-clean-arch/pkg/database/mysqlPkg"
	"log"
)

func main() {

	// ? Load config
	config, err := configs.LoadConfig("uat")

	if err != nil {
		log.Panicf("err: %+v", err)
	}

	db, err := mysqlpkg.NewMysqlDBConnection(&config.Mysql)
	if err != nil {

		log.Printf("error, can't connect to database, %s", err.Error())
	}

	defer db.Close()
	s := servers.NewServer(&config, db)
	s.Run()

}
