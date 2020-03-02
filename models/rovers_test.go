package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidCommandRoverA(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	rover := NewRover(1, 2, "N")

	marsPlateua.DeployRover(rover)

	command := "LMLMLMLMM"

	rover.ProcessCommand(command, marsPlateua)

	assert.Equal(t, "1,3", rover.GetPosition())
	assert.Equal(t, "1 3 N", rover.DescribeState())
}

func Test_ValidCommandRoverB(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	rover := NewRover(3, 3, "E")

	marsPlateua.DeployRover(rover)

	command := "MMRMMRMRRM"

	rover.ProcessCommand(command, marsPlateua)

	assert.Equal(t, "5,1", rover.GetPosition())
	assert.Equal(t, "5 1 E", rover.DescribeState())
}

func Test_ValidCommandStuckUpperRightCorner(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	rover := NewRover(0, 0, "N")

	marsPlateua.DeployRover(rover)

	command := "MMMMMMMMMMMMMMMMRMMMMMMMMMMMMMMMMMMMM"

	rover.ProcessCommand(command, marsPlateua)

	assert.Equal(t, "5,5", rover.GetPosition())
	assert.Equal(t, "5 5 E", rover.DescribeState())
}

func Test_ValidCommandStuckEastEdge(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	rover := NewRover(0, 0, "E")

	marsPlateua.DeployRover(rover)

	command := "MMMMMMMMMMMMMMMMMM"

	rover.ProcessCommand(command, marsPlateua)

	assert.Equal(t, "5,0", rover.GetPosition())
	assert.Equal(t, "5 0 E", rover.DescribeState())
}

func Test_ValidCommandNoCoordsChange(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	rover := NewRover(0, 0, "W")

	marsPlateua.DeployRover(rover)

	command := "MMMMMMMMMMMMMMMMMMLMMMMMMMMMMMMMMMMMMMMMMMM"

	rover.ProcessCommand(command, marsPlateua)

	assert.Equal(t, "0,0", rover.GetPosition())
	assert.Equal(t, "0 0 S", rover.DescribeState())
}

func Test_ValidPathsForTwoRovers(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	roverA := NewRover(1, 2, "N")
	roverB := NewRover(3, 3, "E")

	marsPlateua.DeployRover(roverA)
	marsPlateua.DeployRover(roverB)

	command := "LMLMLMLMM"

	roverA.ProcessCommand(command, marsPlateua)

	command = "MMRMMRMRRM"

	roverB.ProcessCommand(command, marsPlateua)

	assert.Equal(t, "1 3 N", roverA.DescribeState())
	assert.Equal(t, "5 1 E", roverB.DescribeState())
	assert.Equal(t, "1,3", roverA.GetPosition())
	assert.Equal(t, "5,1", roverB.GetPosition())
}

func Test_InvalidPathsForTwoRoversClash(t *testing.T) {
	marsPlateua := NewPlateau(5, 5)

	assert.NotNil(t, marsPlateua)

	roverA := NewRover(0, 0, "N")
	roverB := NewRover(3, 0, "N")

	marsPlateua.DeployRover(roverA)
	marsPlateua.DeployRover(roverB)

	command := "MRMLMRMLMRM"

	roverA.ProcessCommand(command, marsPlateua)

	command = "MMMMMLMMM"

	roverB.ProcessCommand(command, marsPlateua)

	assert.Equal(t, "3 3 E", roverA.DescribeState())
	assert.Equal(t, "0 2 W", roverB.DescribeState())
	assert.Equal(t, "3,3", roverA.GetPosition())
	assert.Equal(t, "0,2", roverB.GetPosition())
}
