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
	args := os.Args
	if len(args) == 0 {
		// environment variables mode
		if os.Getenv("sql_type") == "" {
			log.Fatalln("No environment variables related to sql found. Please provide via arguments or environment variables.")
			return
		}
		settings = utils.ParseEnvironmentVariables()
	} else {
		settings = utils.ParseInputFlags()

		newUsername := os.Getenv("new_username")
		newPassword := os.Getenv("new_password")
		newDatabaseName := os.Getenv("new_database")

		defer dbsvc.CreateDatabaseWithUser(settings, newUsername, newPassword, newDatabaseName)
	}
}
