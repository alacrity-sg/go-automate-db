package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	dbsvc "go-automate-database/database"
	"os"
	"strings"
	"testing"
)

func assertCrud(t *testing.T, uniqueId string, database string, username string, password string) {
	// CRUD test

	settings := &dbsvc.PGSettings{
		Database: database,
		Username: username,
		Password: password,
		Host:     "localhost",
		Port:     "5432",
	}
	db := settings.ConnectPostgresDb()
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE test_%s (test_id varchar(50) primary key)", uniqueId))
	assert.NoError(t, err)

	_, err = db.Exec(fmt.Sprintf("INSERT INTO test_%s VALUES($1)", uniqueId), uniqueId)
	assert.NoError(t, err)

	_, err = db.Exec(fmt.Sprintf("DELETE FROM test_%s WHERE test_id = '%s'", uniqueId, uniqueId))
	assert.NoError(t, err)
	db.Close()
}

func TestMainPostgresEnv(t *testing.T) {
	uniqueUsableId := strings.Replace(uuid.New().String(), "-", "", -1)
	var inputMap map[string]string

	inputMap = map[string]string{
		"SQL_USERNAME":     "admin",
		"SQL_PASSWORD":     "admin",
		"SQL_DATABASE":     "postgres",
		"SQL_HOST":         "localhost",
		"SQL_PORT":         "5432",
		"SQL_NEW_USERNAME": fmt.Sprintf("test_%s", uniqueUsableId),
		"SQL_NEW_PASSWORD": uniqueUsableId,
		"SQL_NEW_DATABASE": fmt.Sprintf("test_%s", uniqueUsableId),
	}
	for k, v := range inputMap {
		t.Setenv(k, v)
	}
	main()
	assertCrud(t, uniqueUsableId, inputMap["SQL_NEW_DATABASE"], inputMap["SQL_NEW_USERNAME"], inputMap["SQL_NEW_PASSWORD"])
}

func TestMainPostgresDirectInput(t *testing.T) {
	uniqueUsableId := strings.Replace(uuid.New().String(), "-", "", -1)
	args := []string{os.Args[0],
		"-type", "postgres",
		"-username", "admin",
		"-password", "admin",
		"-port", "5432",
		"-host", "localhost",
		"-database", "postgres",
		"-new_db", fmt.Sprintf("test_%s", uniqueUsableId),
		"-new_username", fmt.Sprintf("test_user_%s", uniqueUsableId),
		"-new_password", uniqueUsableId,
	}
	os.Args = args
	main()
	assertCrud(t, uniqueUsableId, args[14], args[16], args[18])
}
