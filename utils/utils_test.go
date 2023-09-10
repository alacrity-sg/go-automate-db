package utils

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseInputFlagsNoInput(t *testing.T) {
	var args []string
	_, err := ParseInputFlags(args)
	assert.Error(t, err)
}

func TestParseInputFlagsShortFullSplit(t *testing.T) {
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
	response, err := ParseInputFlags(args)
	assert.NoError(t, err)
	assert.Equal(t, response.Username, args[3])
	assert.Equal(t, response.Password, args[5])
	assert.Equal(t, response.Port, args[7])
	assert.Equal(t, response.Host, args[9])
	assert.Equal(t, response.Database, args[11])
}

func TestParseInputFlagsFullSplit(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	args := []string{
		"-type", "postgres",
		"-username", "admin",
		"-password", "admin",
		"-port", "5432",
		"-host", "localhost",
		"-database", "postgres",
	}
	response, err := ParseInputFlags(args)
	assert.NoError(t, err)
	assert.Equal(t, response.Username, args[3])
	assert.Equal(t, response.Password, args[5])
	assert.Equal(t, response.Port, args[7])
	assert.Equal(t, response.Host, args[9])
	assert.Equal(t, response.Database, args[11])
}
