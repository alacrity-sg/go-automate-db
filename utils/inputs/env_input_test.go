package inputs

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseEnvironmentVariablesFull(t *testing.T) {
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
	response, err := ParseEnvironmentVariables()
	assert.NoError(t, err)

	assert.Equal(t, response.Host, inputMap["SQL_HOST"])
	assert.Equal(t, response.Port, inputMap["SQL_PORT"])
	assert.Equal(t, response.Username, inputMap["SQL_USERNAME"])
	assert.Equal(t, response.Password, inputMap["SQL_PASSWORD"])
	assert.Equal(t, response.Database, inputMap["SQL_DATABASE"])
	assert.Equal(t, response.NewPGSettings[0].Username, inputMap["SQL_NEW_USERNAME"])
	assert.Equal(t, response.NewPGSettings[0].Password, inputMap["SQL_NEW_PASSWORD"])
	assert.Equal(t, response.NewPGSettings[0].Database, inputMap["SQL_NEW_DATABASE"])
}

func TestParseEnvironmentVariablesEmptyNew(t *testing.T) {
	var inputMap map[string]string

	inputMap = map[string]string{
		"SQL_USERNAME": "admin",
		"SQL_PASSWORD": "admin",
		"SQL_DATABASE": "postgres",
		"SQL_HOST":     "localhost",
		"SQL_PORT":     "5432",
	}

	for k, v := range inputMap {
		t.Setenv(k, v)
	}
	_, err := ParseEnvironmentVariables()
	assert.Error(t, err)
}

func TestParseEnvironmentVariablesEmptyRequired(t *testing.T) {
	uniqueUsableId := strings.Replace(uuid.New().String(), "-", "", -1)

	var inputMap map[string]string

	inputMap = map[string]string{
		"SQL_NEW_USERNAME": fmt.Sprintf("test_%s", uniqueUsableId),
		"SQL_NEW_PASSWORD": uniqueUsableId,
		"SQL_NEW_DATABASE": fmt.Sprintf("test_%s", uniqueUsableId),
	}

	for k, v := range inputMap {
		t.Setenv(k, v)
	}
	_, err := ParseEnvironmentVariables()
	assert.Error(t, err)
}

func TestParseEnvironmentVariablesEmptyAll(t *testing.T) {
	var inputMap map[string]string

	inputMap = map[string]string{}

	for k, v := range inputMap {
		t.Setenv(k, v)
	}
	_, err := ParseEnvironmentVariables()
	assert.Error(t, err)
}
