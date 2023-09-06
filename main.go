package main

import (
	_ "github.com/lib/pq"
	dbsvc "go-automate-database/database"
	"os"
)

func singleUseDatabaseCase() {
	settings := &dbsvc.PGSettings{
		Username: os.Getenv("sql_username"),
		Password: os.Getenv("sql_password"),
		Host:     os.Getenv("sql_host"),
		Port:     os.Getenv("sql_port"),
		Database: "",
	}
	// TODO: Implement 3 input mode: yaml, cli with flags, environment variable
	if os.Getenv("sql_database") != "" {
		settings.Database = os.Getenv("sql_database")
	} else {
		// Operations mode
		settings.Database = "postgres"
	}

	newUsername := os.Getenv("new_username")
	newPassword := os.Getenv("new_password")
	newDatabaseName := os.Getenv("new_database")

	defer dbsvc.CreateDatabaseWithUser(settings, newUsername, newPassword, newDatabaseName)

}

func main() {
	//
	args := os.Args
	if len(args) == 0 {
		// file mode
		println("Not implemented yet")
	} else {
		singleUseDatabaseCase()
	}
}
