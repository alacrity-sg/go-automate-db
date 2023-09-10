package main

import (
	"flag"
	_ "github.com/lib/pq"
	"go-automate-database/utils"
	"log"
	"os"
)

func main() {
	//
	//var settings *dbsvc.PGSettings
	args := os.Args
	if len(args) == 0 {
		// environment variables mode
		if os.Getenv("sql_type") == "" {
			log.Fatalln("No environment variables related to sql found. Please provide via arguments or environment variables.")
			return
		}
		utils.ParseEnvironmentVariables()
	} else {
		var name string
		flag.StringVar(&name, "name", "", "asdasd")
		a := os.Args[1:]
		if args != nil {
			log.Println(args)
			a = args
		}
		err := flag.CommandLine.Parse(a)
		if err != nil {
			println(err.Error())
		}
		log.Println(name)
		//newUsername := os.Getenv("new_username")
		//newPassword := os.Getenv("new_password")
		//newDatabaseName := os.Getenv("new_database")

		//defer settings.CreateDatabaseWithUser(newUsername, newPassword, newDatabaseName)
	}
}
