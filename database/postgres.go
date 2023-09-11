package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type DBSettings[T any] struct {
}
type PGSettings struct {
	Username      string
	Password      string
	Host          string
	Port          string
	Database      string
	NewPGSettings []*NewPGSettings
}

type NewPGSettings struct {
	Username string
	Password string
	Database string
}

func (settings *PGSettings) CreateDatabaseWithUser() error {
	db := settings.ConnectPostgresDb()
	defer db.Close()

	if len(settings.NewPGSettings) == 0 {
		return errors.New("new database settings not configured. Please consult the manual")
	}

	for _, newDatabase := range settings.NewPGSettings {
		createDatabaseStmt := fmt.Sprintf("CREATE DATABASE %s", newDatabase.Database)

		_, err := db.Exec(createDatabaseStmt)
		if err != nil {
			return err
		} else {
			println(fmt.Sprintf("Created database [%s] successfully", newDatabase.Database))
		}

		createUserStmt := fmt.Sprintf("CREATE USER %s WITH ENCRYPTED PASSWORD '%s'", newDatabase.Username, newDatabase.Password)
		_, err = db.Exec(createUserStmt)
		if err != nil {
			return err
		} else {
			println(fmt.Sprintf("Created user [%s] successfully", newDatabase.Username))
		}

		grantUserStmt := fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s", newDatabase.Database, newDatabase.Username)
		_, err = db.Exec(grantUserStmt)
		if err != nil {
			return err
		} else {
			println(fmt.Sprintf("Granted user [%s] permissions to database [%s] successfully", newDatabase.Username, newDatabase.Database))
		}

	}
	return nil
}

func (settings *PGSettings) ConnectPostgresDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		settings.Host,
		settings.Port,
		settings.Username,
		settings.Password,
		settings.Database,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
