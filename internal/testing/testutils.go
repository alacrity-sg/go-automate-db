package testing

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateTestId() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
