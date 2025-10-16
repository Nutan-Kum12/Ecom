package main

import (
	"log"
	"os"

	"github.com/Nutan-Kum12/Ecom/config"
	"github.com/Nutan-Kum12/Ecom/db"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Net:    "tcp",
		Addr:   config.Envs.DBAddress,
		// DBName: config.Envs.DBName,404 page not found
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[(len(os.Args) - 1)] // get the last argument
	switch cmd {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command: %s", cmd)
	}
}
