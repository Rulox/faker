package faker

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestYAMLLoader(t *testing.T) {
	var f Faker
	err := f.SetLocale("es_ES")
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.IsType(t, "", f.Address.City())
	assert.IsType(t, "", f.Address.Full())
}
