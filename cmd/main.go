package main

import (
	"log"

	"github.com/Nutan-Kum12/Ecom/cmd/api"
	"github.com/Nutan-Kum12/Ecom/config"
	"github.com/Nutan-Kum12/Ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)

	}

	server := api.NewAPIserver(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
