package inputs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	helper "go-automate-database/internal/testing"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestParseInputFile(t *testing.T) {
	uniqueId := helper.GenerateTestId()
	config := &FileInputV1{
		Version: 1,
		SQL: FileInputV1SQL{
			Username: "admin",
			Password: "admin",
			Host:     "localhost",
			Port:     "5432",
			Database: "postgres",
			Type:     "postgres",
			Mode:     "operations",
		},
		Databases: []FileInputV1Databases{
			{
				Name:     fmt.Sprintf("test_%s", uniqueId),
				Username: fmt.Sprintf("test_%s", uniqueId),
				Password: fmt.Sprintf(uniqueId),
			},
		},
	}
	data, err := yaml.Marshal(config)
	assert.NoError(t, err)

	file, err := os.CreateTemp("", fmt.Sprintf("sample-%s.yaml", uniqueId))
	assert.NoError(t, err)
	write, err := file.Write(data)
	assert.NoError(t, err)
	fmt.Printf("wrote %d bytes\n", write)
	response, err := ParseInputFile(file.Name())
	assert.NoError(t, err)

	assert.Equal(t, response.Host, config.SQL.Host)
	assert.Equal(t, response.Port, config.SQL.Port)
	assert.Equal(t, response.Username, config.SQL.Username)
	assert.Equal(t, response.Password, config.SQL.Password)
	assert.Equal(t, response.Database, config.SQL.Database)
	assert.Equal(t, response.NewPGSettings[0].Username, config.Databases[0].Username)
	assert.Equal(t, response.NewPGSettings[0].Password, config.Databases[0].Password)
	assert.Equal(t, response.NewPGSettings[0].Database, config.Databases[0].Name)
}
