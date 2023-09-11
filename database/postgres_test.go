package database

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func assertCrud(t *testing.T, uniqueId string, database string, username string, password string) {
	// CRUD test

	settings := &PGSettings{
		Database: database,
		Username: username,
		Password: password,
		Host:     "localhost",
		Port:     "5432",
	}
	db := settings.connectPostgresDb()
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE test_%s (test_id varchar(50) primary key)", uniqueId))
	assert.NoError(t, err)

	_, err = db.Exec(fmt.Sprintf("INSERT INTO test_%s VALUES($1)", uniqueId), uniqueId)
	assert.NoError(t, err)

	_, err = db.Exec(fmt.Sprintf("DELETE FROM test_%s WHERE test_id = '%s'", uniqueId, uniqueId))
	assert.NoError(t, err)
	db.Close()
}

func TestCreateDatabaseWithUser_Single(t *testing.T) {
	uniqueId := uuid.New().String()
	uniqueUsableId := strings.Replace(uniqueId, "-", "", -1)

	settings := &PGSettings{
		Username:      "admin",
		Password:      "admin",
		Host:          "localhost",
		Port:          "5432",
		Database:      "postgres",
		NewPGSettings: make([]*NewPGSettings, 0),
	}
	settings.NewPGSettings = append(settings.NewPGSettings, &NewPGSettings{
		Database: fmt.Sprintf("test_%s_db", uniqueUsableId),
		Username: fmt.Sprintf("test_user_%s", uniqueUsableId),
		Password: uniqueUsableId,
	})
	settings.CreateDatabaseWithUser()
	assertCrud(t, uniqueUsableId,
		settings.NewPGSettings[0].Database,
		settings.NewPGSettings[0].Username,
		settings.NewPGSettings[0].Password,
	)
}

func TestCreateDatabaseWithUser_Multiple(t *testing.T) {
	uniqueId := uuid.New().String()
	uniqueUsableId := strings.Replace(uniqueId, "-", "", -1)

	settings := &PGSettings{
		Username:      "admin",
		Password:      "admin",
		Host:          "localhost",
		Port:          "5432",
		Database:      "postgres",
		NewPGSettings: make([]*NewPGSettings, 0),
	}

	for i := 0; i < 2; i++ {
		settings.NewPGSettings = append(settings.NewPGSettings, &NewPGSettings{
			Database: fmt.Sprintf("test_%d_%s_db", i, uniqueUsableId),
			Username: fmt.Sprintf("test_user_%d_%s", i, uniqueUsableId),
			Password: uniqueUsableId,
		})
	}

	settings.CreateDatabaseWithUser()

	for i := 0; i < 2; i++ {
		log.Println(settings.NewPGSettings[i])
		assertCrud(t, uniqueUsableId,
			settings.NewPGSettings[i].Database,
			settings.NewPGSettings[i].Username,
			settings.NewPGSettings[i].Password,
		)
	}
}

func TestCreateDatabaseWithUser_EmptyNew(t *testing.T) {

	settings := &PGSettings{
		Username:      "admin",
		Password:      "admin",
		Host:          "localhost",
		Port:          "5432",
		Database:      "postgres",
		NewPGSettings: make([]*NewPGSettings, 0),
	}

	err := settings.CreateDatabaseWithUser()
	assert.Error(t, err)
}

