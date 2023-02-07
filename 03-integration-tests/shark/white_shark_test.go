package shark

import (
	"errors"
	"integrationtests/pkg/storage"
	"integrationtests/prey"
	"integrationtests/simulator"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWhiteShark_Hunt(t *testing.T){
	//arrange
	data := map[string]interface{}{
		"white_shark_speed": 10.1,
		"white_shark_x":     1.1,
		"white_shark_y":     2.2,
		"tuna_speed":        5.4,
	}
	storage := storage.NewStorageMock(data)
	prey := prey.CreateTuna(storage)
	simulator := simulator.NewCatchSimulator(10.1)
	whiteShark := CreateWhiteShark(simulator, storage)

	//act
	result := whiteShark.Hunt(prey)

	//arrange
	assert.NoError(t, result)
	assert.True(t, storage.Spy)
}

func TestWhiteShark_Hunt_Fail(t *testing.T){
	//arrange
	data := map[string]interface{}{
		"white_shark_speed": 5.4,
		"white_shark_x":     1.1,
		"white_shark_y":     2.2,
		"tuna_speed":        10.4,
	}
	storage := storage.NewStorageMock(data)
	prey := prey.CreateTuna(storage)
	simulator := simulator.NewCatchSimulator(10.1)
	whiteShark := CreateWhiteShark(simulator, storage)
	expectedResult := errors.New("could not hunt the prey")

	//act
	result := whiteShark.Hunt(prey)

	//arrange
	assert.Error(t, result)
	assert.Equal(t, expectedResult, result)
	assert.True(t, storage.Spy)
}

func BenchmarkHunt(b *testing.B){
	data := map[string]interface{}{
		"white_shark_speed": 10.1,
		"white_shark_x":     1.1,
		"white_shark_y":     2.2,
		"tuna_speed":        5.4,
	}
	storage := storage.NewStorageMock(data)
	prey := prey.CreateTuna(storage)
	simulator := simulator.NewCatchSimulator(10.1)
	whiteShark := CreateWhiteShark(simulator, storage)
	for i := 0; i < b.N; i++{
		whiteShark.Hunt(prey)
	}
}