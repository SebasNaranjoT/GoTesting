package prey

import (
	"integrationtests/pkg/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeed(t *testing.T){
	//arrange
	tuna := NewMockTuna()
	tuna.speed = 5.0
	expectedSpeed := 5.0

	//act
	result := tuna.GetSpeed()

	//assert
	assert.Equal(t, expectedSpeed, result)
	assert.True(t, tuna.spy)
}

func TestCreateTuna(t *testing.T){
	//arrange
	data := map[string]interface{}{
		"white_shark_speed": 10.1,
		"white_shark_x":     1.1,
		"white_shark_y":     2.2,
		"tuna_speed":        5.4,
	}
	storage := storage.NewStorageMock(data)
	expectedSpeed := 5.4

	//act
	result := CreateTuna(storage)

	//assert
	assert.Equal(t, expectedSpeed, result.GetSpeed())
}