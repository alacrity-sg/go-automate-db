package database

import (
	"database/sql"
	"fmt"
)

type DBSettings[T any] struct {
}
type PGSettings struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (settings *PGSettings) CreateDatabaseWithUser(newUser string, newPass string, newDb string) {
	db := settings.connectPostgresDb()
	defer db.Close()

	createDatabaseStmt := fmt.Sprintf("CREATE DATABASE %s", newDb)

	_, err := db.Exec(createDatabaseStmt)
	if err != nil {
		println(err.Error())
	} else {
		println(fmt.Sprintf("Created database [%s] successfully", newDb))
	}

	createUserStmt := fmt.Sprintf("CREATE USER %s WITH ENCRYPTED PASSWORD '%s'", newUser, newPass)
	_, err = db.Exec(createUserStmt)
	if err != nil {
		println(err.Error())
	} else {
		println(fmt.Sprintf("Created user [%s] successfully", newUser))
	}

	grantUserStmt := fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s", newDb, newUser)
	_, err = db.Exec(grantUserStmt)
	if err != nil {
		println(err.Error())
	} else {
		println(fmt.Sprintf("Granted user [%s] permissions to database [%s] successfully", newUser, newDb))
	}

}

func (settings *PGSettings) connectPostgresDb() *sql.DB {
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
