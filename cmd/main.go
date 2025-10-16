package main

import (
	"database/sql"
	"log"

	"github.com/Nutan-Kum12/Ecom/cmd/api"
	"github.com/Nutan-Kum12/Ecom/config"
	"github.com/Nutan-Kum12/Ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Net:    "tcp",
		Addr:   config.Envs.DBAddress,
		// DBName: config.Envs.DBName,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server := api.NewAPIserver(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
func initStorage(db *sql.DB) {
	// First, create database if it doesn't exist
	createDB := "CREATE DATABASE IF NOT EXISTS ecom"
	_, err := db.Exec(createDB)
	if err != nil {
		log.Printf("Warning: Could not create database: %v", err)
	}

	// Use the database
	useDB := "USE ecom"
	_, err = db.Exec(useDB)
	if err != nil {
		log.Fatal("Could not use database:", err)
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected successfully")
}
