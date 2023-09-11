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

func TestParseInputFlagsShort_NoNewInput(t *testing.T) {
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

func TestParseInputFlagsFullSplit(t *testing.T) {
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
