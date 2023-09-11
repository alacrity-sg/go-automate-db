package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestParseInputFlagsNoInput(t *testing.T) {
	var args []string
	_, err := ParseInputFlags(args)
	assert.Error(t, err)
}

func TestParseInputFlagsShortFullSplit(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	uniqueUsableId := strings.Replace(uuid.New().String(), "-", "", -1)
	args := []string{
		"-t", "postgres",
		"-u", "admin",
		"-p", "admin",
		"-P", "5432",
		"-h", "localhost",
		"-db", "postgres",
		"-new_db", fmt.Sprintf("test_%s", uniqueUsableId),
		"-new_username", fmt.Sprintf("test_user_%s", uniqueUsableId),
		"-new_password", uniqueUsableId,
	}
	flag.NewFlagSet("test", flag.ContinueOnError)
	response, err := ParseInputFlags(args)
	assert.NoError(t, err)
	assert.Equal(t, response.Username, args[3])
	assert.Equal(t, response.Password, args[5])
	assert.Equal(t, response.Port, args[7])
	assert.Equal(t, response.Host, args[9])
	assert.Equal(t, response.Database, args[11])
}

func TestParseInputFlagsShortNoNewInput(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	args := []string{
		"-t", "postgres",
		"-u", "admin",
		"-p", "admin",
		"-P", "5432",
		"-h", "localhost",
		"-db", "postgres",
	}

	flag.NewFlagSet("test", flag.ContinueOnError)
	_, err := ParseInputFlags(args)
	log.Println(err.Error())
	assert.Error(t, err)
}

func TestParseInputFlagsFull(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	uniqueUsableId := strings.Replace(uuid.New().String(), "-", "", -1)
	args := []string{
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
	response, err := ParseInputFlags(args)
	assert.NoError(t, err)
	assert.Equal(t, response.Username, args[3])
	assert.Equal(t, response.Password, args[5])
	assert.Equal(t, response.Port, args[7])
	assert.Equal(t, response.Host, args[9])
	assert.Equal(t, response.Database, args[11])
}

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
