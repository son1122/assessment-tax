package main

import (
	"database/sql"
	"fmt"
	"github.com/son1122/assessment-tax/constant"
	"github.com/son1122/assessment-tax/db"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq" // Import your SQL driver
)

func executeSQLFile(db *sql.DB, filepath string) error {

	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	sqlCommands := string(fileContent)
	_, err = db.Exec(sqlCommands)
	if err != nil {
		return err
	}

	return nil
}

func main2() {
	constant.InitConfig()
	cfg := constant.Get()

	db.InitDB(cfg.DatabaseURL)
	if err := executeSQLFile(db.DB, "./master_deduct.sql"); err != nil {
		log.Fatalf("Failed to execute master_deduct.sql: %v", err)
	}
	if err := executeSQLFile(db.DB, "./master_tax_level.sql"); err != nil {
		log.Fatalf("Failed to execute master_tax_level.sql: %v", err)
	}

	fmt.Println("Database initialized successfully!")
}
