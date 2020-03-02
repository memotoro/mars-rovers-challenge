package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidPlateau(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	assert.Equal(t, "(0,0) (5,5)", marsPlateua.GetDimensions())
}

func Test_InvalidPlateau(t *testing.T) {
	marsPlateua := NewPlateau(-5, -5)

	assert.Nil(t, marsPlateua)
}

func Test_ValidCoords(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	valid := marsPlateua.IsValidPairCoordinates(2, 3)

	assert.Equal(t, true, valid)
}

func Test_InvalidCoords(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	valid := marsPlateua.IsValidPairCoordinates(6, 5)

	assert.Equal(t, false, valid)
}

func Test_InvalidNegativeCoords(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	valid := marsPlateua.IsValidPairCoordinates(-1, 5)

	assert.Equal(t, false, valid)
}

func Test_DeployRoversToPlateau(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	roverA := NewRover(0, 0, "N")
	added := marsPlateua.DeployRover(roverA)

	assert.Equal(t, 1, marsPlateua.NumberOfRovers())
	assert.Equal(t, true, added)

	roverB := NewRover(1, 1, "S")
	added = marsPlateua.DeployRover(roverB)
	assert.Equal(t, true, added)

	assert.Equal(t, 2, marsPlateua.NumberOfRovers())
}

func Test_ValidPositionForRoverInPlateau(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	roverA := NewRover(0, 0, "N")
	roverB := NewRover(1, 1, "S")

	added := marsPlateua.DeployRover(roverA)
	assert.Equal(t, true, added)

	added = marsPlateua.DeployRover(roverB)
	assert.Equal(t, true, added)

	validPosition := marsPlateua.IsPositionAvailable(1, 0)
	assert.Equal(t, true, validPosition)

	validPosition = marsPlateua.IsPositionAvailable(0, 1)
	assert.Equal(t, true, validPosition)
}

func Test_InvalidPositionForRoverInPlateau(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	roverA := NewRover(0, 0, "N")
	roverB := NewRover(1, 1, "S")

	added := marsPlateua.DeployRover(roverA)
	assert.Equal(t, true, added)

	added = marsPlateua.DeployRover(roverB)
	assert.Equal(t, true, added)

	validPosition := marsPlateua.IsPositionAvailable(1, 1)
	assert.Equal(t, false, validPosition)

	validPosition = marsPlateua.IsPositionAvailable(0, 0)
	assert.Equal(t, false, validPosition)
}

func Test_InvalidPositionForNewRoversInPlateu(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	roverA := NewRover(0, 0, "N")

	added := marsPlateua.DeployRover(roverA)
	assert.Equal(t, true, added)

	roverB := NewRover(0, 0, "S")

	added = marsPlateua.DeployRover(roverB)
	assert.Equal(t, false, added)

	assert.Equal(t, 1, marsPlateua.NumberOfRovers())
}
