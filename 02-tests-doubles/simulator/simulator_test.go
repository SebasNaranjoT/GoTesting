package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCatch(t *testing.T) {
	t.Run("success path", func (t *testing.T)  {
		//arrange
		sm := NewSimulatorMock()
		sm.result = true
		expected := true
		distance := 4.0
		speed := 4.0
		catchSpeed := 4.0
		//act
		result := sm.CanCatch(distance, speed, catchSpeed)

		//assert
		assert.Equal(t, expected, result)
	})

	t.Run("failure path", func (t *testing.T)  {
		//arrange
		sm := NewSimulatorMock()
		sm.result = false
		expected := false
		distance := 4.0
		speed := 4.0
		catchSpeed := 4.0
		//act
		result := sm.CanCatch(distance, speed, catchSpeed)

		//assert
		assert.Equal(t, expected, result)
	})
}