package main

import (
	_ "github.com/lib/pq"
	dbsvc "go-automate-database/database"
	"go-automate-database/utils"
	"log"
	"os"
)

func main() {
	//
	var settings *dbsvc.PGSettings
	var err error
	args := os.Args
	if len(args) == 0 {
		// environment variables mode
		if os.Getenv("sql_type") == "" {
			log.Fatalln("No environment variables related to sql found. Please provide via arguments or environment variables.")
			return
		}
		settings = utils.ParseEnvironmentVariables()
	} else {
		settings, err = utils.ParseInputFlags(os.Args[1:])
		if err != nil {
			log.Fatalf("Error parsing command line arguments. %s", err.Error())
		}
	}
	defer settings.CreateDatabaseWithUser("newUsername", "newPassword", "newDatabaseName")
}
