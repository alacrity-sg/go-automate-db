package main

import (
	dbsvc "go-automate-database/database"
	"go-automate-database/utils"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	//
	var settings *dbsvc.PGSettings
	var err error

	if os.Getenv("SQL_HOST") != "" {
		// Default to direct input mode
		log.Println("Defaulting to environment variable mode")
		settings, err = utils.ParseEnvironmentVariables()
	} else {
		// Env mode
		log.Println("Using direct input mode")
		settings, err = utils.ParseInputFlags(os.Args[1:])
	}

	if err != nil {
		log.Fatalf(err.Error())
	}
	err = settings.CreateDatabaseWithUser()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
