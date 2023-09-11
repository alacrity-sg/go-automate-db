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
	args := os.Args
	if len(args) == 0 {
		// environment variables mode
		if os.Getenv("sql_type") == "" {
			log.Fatalln("No environment variables related to sql found. Please provide via arguments or environment variables.")
			return
		}
		settings, err = utils.ParseEnvironmentVariables()
	} else {
		settings, err = utils.ParseInputFlags(os.Args[1:])

	}

	if err != nil {
		log.Fatalf(err.Error())
	}
	defer settings.CreateDatabaseWithUser()
}
