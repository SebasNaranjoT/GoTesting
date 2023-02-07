package hunt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	//arrange
	prey := Prey{
		speed: 3,
	}
	shark := Shark{
		tired: false,
		hungry:  true,
		speed: 4,
	}

	//act
	result := shark.Hunt(&prey)
	
	//assert
	assert.False(t, shark.hungry)
	assert.True(t, shark.tired)
	assert.NoError(t, result)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//arrange
	prey := Prey{}
	shark := Shark{
		tired: true,
	}
	expectedErr := errors.New("cannot hunt, i am really tired")

	//act
	result := shark.Hunt(&prey)

	//assert
	assert.Error(t, result)
	assert.Equal(t, result, expectedErr)
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
		//arrange
		prey := Prey{
			speed: 3,
		}
		shark := Shark{		
			hungry:  false,
		}
		expectedErr := errors.New("cannot hunt, i am not hungry")
	
		//act
		result := shark.Hunt(&prey)

		//assert
		assert.Error(t, result)
		assert.Equal(t, expectedErr, result)
}

func TestSharkCannotReachThePrey(t *testing.T) {
		//arrange
		prey := Prey{
			speed: 4,
		}
		shark := Shark{
			tired: false,
			hungry:  true,
			speed: 3,
		}
		expectedErr := errors.New("could not catch it")
	
		//act
		result := shark.Hunt(&prey)
		
		//assert
		assert.Error(t, result)
		assert.Equal(t, expectedErr, result)
}

func TestSharkHuntNilPrey(t *testing.T) {
}
