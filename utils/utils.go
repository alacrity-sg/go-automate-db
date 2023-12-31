package utils

import (
	"errors"
	"flag"
	"go-automate-database/database"
	"log"
	"os"
	"strings"
)

type CommandLineInputs struct {
	database    string `default:""`
	username    string `default:""`
	password    string `default:""`
	host        string `default:""`
	port        string `default:""`
	dbType      string `default:""`
	newDatabase string `default:""`
	newUsername string `default:""`
	newPassword string `default:""`
}

func ParseEnvironmentVariables() (*database.PGSettings, error) {
	settings := &database.PGSettings{
		Username:      os.Getenv("SQL_USERNAME"),
		Password:      os.Getenv("SQL_PASSWORD"),
		Host:          os.Getenv("SQL_HOST"),
		Port:          os.Getenv("SQL_PORT"),
		Database:      os.Getenv("SQL_DATABASE"),
		NewPGSettings: make([]*database.NewPGSettings, 0),
	}

	// Check if required variables are available
	if settings.Username == "" || settings.Password == "" || settings.Host == "" {
		return nil, errors.New("required environment settings [SQL_USERNAME, SQL_PASSWORD, SQL_HOST] is empty")
	}

	// Check if required new database settings is added
	newDbSettings := &database.NewPGSettings{
		Username: os.Getenv("SQL_NEW_USERNAME"),
		Password: os.Getenv("SQL_NEW_PASSWORD"),
		Database: os.Getenv("SQL_NEW_DATABASE"),
	}
	if newDbSettings.Username == "" || newDbSettings.Password == "" || newDbSettings.Database == "" {
		return nil, errors.New("required environment settings [SQL_NEW_USERNAME, SQL_NEW_PASSWORD, SQL_NEW_DATABASE] is empty")
	}
	settings.NewPGSettings = append(settings.NewPGSettings, newDbSettings)
	return settings, nil
}

func ParseInputFlags(args []string) (*database.PGSettings, error) {
	inputs := &CommandLineInputs{}
	flag.StringVar(&inputs.dbType, "t", "", "to specify sql type")
	flag.StringVar(&inputs.dbType, "type", inputs.dbType, "to specify sql type")

	flag.StringVar(&inputs.username, "u", "", "username to connect to db as")
	flag.StringVar(&inputs.username, "username", inputs.username, "username to connect to db as")

	flag.StringVar(&inputs.password, "p", "", "password to connect to db with")
	flag.StringVar(&inputs.password, "password", inputs.password, "password to connect to db with")

	flag.StringVar(&inputs.host, "h", "", "host to connect to db with")
	flag.StringVar(&inputs.host, "host", inputs.host, "host to connect to db with")

	flag.StringVar(&inputs.port, "P", "", "host to connect to db with")
	flag.StringVar(&inputs.port, "port", inputs.port, "host to connect to db with")

	flag.StringVar(&inputs.database, "db", "", "database to connect to")
	flag.StringVar(&inputs.database, "database", inputs.database, "database to connect to")

	flag.StringVar(&inputs.newDatabase, "new_db", "", "database to connect to")
	flag.StringVar(&inputs.newUsername, "new_username", "", "database to connect to")
	flag.StringVar(&inputs.newPassword, "new_password", "", "database to connect to")

	err := flag.CommandLine.Parse(args)
	if err != nil {
		log.Fatalf("Fatal error processing arguments. %s", err.Error())
	}
	return buildDatabaseSettings(inputs)
}

func buildDatabaseSettings(inputs *CommandLineInputs) (*database.PGSettings, error) {
	// Check if struct is empty:
	if inputs.dbType == "" {
		return nil, errors.New("required input [t,type] is not provided. please review the documentation")
	}
	log.Println(inputs)
	switch strings.ToLower(inputs.dbType) {
	case "postgres":
		settings := &database.PGSettings{
			Username:      inputs.username,
			Password:      inputs.password,
			Host:          inputs.host,
			Port:          inputs.port,
			Database:      inputs.database,
			NewPGSettings: make([]*database.NewPGSettings, 0),
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
			return nil, errors.New("required argument [username] is empty.")
		}
		if settings.Password == "" {
			return nil, errors.New("required argument [password] is empty.")
		}
		if settings.Host == "" {
			return nil, errors.New("required argument [host] is empty.")
		}

		if inputs.newDatabase == "" && inputs.newUsername == "" && inputs.newPassword == "" {
			return nil, errors.New("required arguments [new_db, new_username, new_password] is empty.")
		}
		pgSettings := &database.NewPGSettings{
			Database: inputs.newDatabase,
			Username: inputs.newUsername,
			Password: inputs.newPassword,
		}
		settings.NewPGSettings = append(settings.NewPGSettings, pgSettings)
		return settings, nil

	default:
		return nil, errors.New("input sqlType is not of an accepted value")
	}
}
