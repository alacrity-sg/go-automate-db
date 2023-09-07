package utils

import (
	"flag"
	"go-automate-database/database"
	"log"
	"os"
	"strings"
)

func ParseEnvironmentVariables() *database.PGSettings {
	settings := &database.PGSettings{
		Username: os.Getenv("sql_username"),
		Password: os.Getenv("sql_password"),
		Host:     os.Getenv("sql_host"),
		Port:     os.Getenv("sql_port"),
		Database: "",
	}

	if os.Getenv("sql_database") != "" {
		settings.Database = os.Getenv("sql_database")
	} else {
		// Operations mode
		settings.Database = "postgres"
	}
	return settings

}

func ParseInputFlags() *database.PGSettings {
	var sqlType string
	flag.StringVar(&sqlType, "t", "postgres", "to specify sql type, defaults to postgres")
	flag.StringVar(&sqlType, "type", sqlType, "to specify sql type")

	var sqlUsername string
	flag.StringVar(&sqlUsername, "u", "", "username to connect to db as")
	flag.StringVar(&sqlUsername, "username", sqlUsername, "username to connect to db as")

	var sqlPassword string
	flag.StringVar(&sqlPassword, "P", "", "password to connect to db with")
	flag.StringVar(&sqlPassword, "password", sqlPassword, "password to connect to db with")

	var sqlHost string
	flag.StringVar(&sqlHost, "h", "", "host to connect to db with")
	flag.StringVar(&sqlHost, "host", sqlHost, "host to connect to db with")

	var sqlPort string
	flag.StringVar(&sqlHost, "p", "", "host to connect to db with")
	flag.StringVar(&sqlHost, "port", sqlPort, "host to connect to db with")

	var sqlDatabase string
	flag.StringVar(&sqlDatabase, "db", "", "database to connect to")
	flag.StringVar(&sqlDatabase, "database", sqlDatabase, "database to connect to")

	return buildDatabaseSettings(
		sqlType, sqlUsername, sqlPassword, sqlHost, sqlPort, sqlDatabase)
}

func buildDatabaseSettings(sqlType string, sqlUsername string, sqlPassword string, sqlHost string, sqlPort string, sqlDatabase string) *database.PGSettings {
	switch strings.ToLower(sqlType) {
	case "postgres":
		settings := &database.PGSettings{
			Username: sqlUsername,
			Password: sqlPassword,
			Host:     sqlHost,
			Port:     sqlPort,
			Database: sqlDatabase,
		}
		if settings.Database == "" {
			log.Println("Database was not provided. Defaulting to postgres for operations mode")
			settings.Database = "postgres"
		}
		if settings.Port == "" {
			log.Println("Port was not provided. Defaulting to 5432 for postgres")
			settings.Port = "5432"
		}
		if settings.Username == "" {
			log.Fatalln("Required argument [username] was not provided. Exiting")
		}
		if settings.Password == "" {
			log.Fatalln("Required argument [password] was not provided. Exiting")
		}
		if settings.Host == "" {
			log.Fatalln("Required argument [host] was not provided. Exiting")
		}
		return settings
	default:
		log.Fatalln("Input sqlType is not of an accepted value.")
	}
	return nil
}
