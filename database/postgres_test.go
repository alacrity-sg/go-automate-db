package database

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"strings"
	"testing"
)

func TestCreateDatabaseWithUser(t *testing.T) {

	uniqueId := uuid.New().String()
	uniqueUsableId := strings.Replace(uniqueId, "-", "", -1)
	newDb := fmt.Sprintf("test_%s_db", uniqueUsableId)
	newUser := fmt.Sprintf("test_user_%s", uniqueUsableId)
	newPass := fmt.Sprintf("default_password")
	settings := &PGSettings{
		Username: "admin",
		Password: "admin",
		Host:     "localhost",
		Port:     "5432",
		Database: "postgres",
	}
	CreateDatabaseWithUser(settings, newUser, newPass, newDb)
	settings.Username = newUser
	settings.Password = newPass
	settings.Database = newDb
	db := connectPostgresDb(settings)
	defer db.Close()
	// CRUD test
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE test_%s (test_id varchar(50) primary key)", uniqueUsableId))
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = db.Exec(fmt.Sprintf("INSERT INTO test_%s VALUES($1)", uniqueUsableId), uniqueUsableId)
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = db.Exec(fmt.Sprintf("DELETE FROM test_%s WHERE test_id = '%s'", uniqueUsableId, uniqueUsableId))

	if err != nil {
		t.Fatalf(err.Error())
	}
}
