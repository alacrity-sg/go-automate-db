package inputs

import (
	"errors"
	"go-automate-database/database"
	"os"
)

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
