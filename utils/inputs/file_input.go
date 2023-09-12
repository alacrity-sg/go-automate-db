package inputs

import (
	"go-automate-database/database"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type FileInputV1SQL struct {
	Type     string `yaml:"type"`
	Mode     string `yaml:"mode"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type FileInputV1Databases struct {
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type FileInputV1 struct {
	Version   int                    `yaml:"version"`
	SQL       FileInputV1SQL         `yaml:"sql"`
	Databases []FileInputV1Databases `yaml:"databases"`
}

func (input *FileInputV1) GenerateDatabaseSettings() *database.PGSettings {
	settings := &database.PGSettings{
		Username:      input.SQL.Username,
		Password:      input.SQL.Password,
		Host:          input.SQL.Host,
		Port:          input.SQL.Port,
		Database:      input.SQL.Database,
		NewPGSettings: []*database.NewPGSettings{},
	}
	for _, db := range input.Databases {
		settings.NewPGSettings = append(settings.NewPGSettings, &database.NewPGSettings{
			Username: db.Username,
			Password: db.Password,
			Database: db.Name,
		})
	}

	return settings
}

func ParseInputFile(fileName string) (*database.PGSettings, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	input := &FileInputV1{}
	log.Println(string(data))
	err = yaml.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}

	return input.GenerateDatabaseSettings(), nil
}
