package prey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeed(t *testing.T){
	//arrange
	ps := NewPreyStub()
	expectedResult := 5.0
	//act
	result := ps.GetSpeed()
	//assert
	assert.Equal(t, expectedResult, result)
}