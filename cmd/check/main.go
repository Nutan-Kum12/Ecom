package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Nutan-Kum12/Ecom/config"
	"github.com/Nutan-Kum12/Ecom/db"
	mysqlCfg "github.com/go-sql-driver/mysql"
)

func main() {
	database, err := db.NewMySQLStorage(mysqlCfg.Config{
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
	defer database.Close()

	// Check if users table exists and show its structure
	fmt.Println("Checking if users table exists...")

	rows, err := database.Query("SHOW TABLES LIKE 'users'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("✅ Users table exists!")

		// Show table structure
		fmt.Println("\nTable structure:")
		structRows, err := database.Query("DESCRIBE users")
		if err != nil {
			log.Fatal(err)
		}
		defer structRows.Close()

		fmt.Printf("%-20s %-20s %-10s %-10s %-20s %-10s\n", "Field", "Type", "Null", "Key", "Default", "Extra")
		fmt.Println(strings.Repeat("-", 100))

		for structRows.Next() {
			var field, fieldType, null, key, extra string
			var defaultVal sql.NullString
			err := structRows.Scan(&field, &fieldType, &null, &key, &defaultVal, &extra)
			if err != nil {
				log.Fatal(err)
			}

			defaultStr := "NULL"
			if defaultVal.Valid {
				defaultStr = defaultVal.String
			}

			fmt.Printf("%-20s %-20s %-10s %-10s %-20s %-10s\n", field, fieldType, null, key, defaultStr, extra)
		}
	} else {
		fmt.Println("❌ Users table does not exist!")
	}

	// Show all tables
	fmt.Println("\nAll tables in database:")
	allTablesRows, err := database.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer allTablesRows.Close()

	for allTablesRows.Next() {
		var tableName string
		err := allTablesRows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("- %s\n", tableName)
	}
}
