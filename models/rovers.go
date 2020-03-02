package models

import (
	"fmt"
	"strings"
)

// Rover - defines rover attributes
type Rover struct {
	x    int
	y    int
	head string
}

// NewRover constructs a new rover
func NewRover(posX, posY int, head string) *Rover {
	return &Rover{x: posX, y: posY, head: strings.ToUpper(head)}
}

// ProcessCommand gets the command and process it depending on the current state of the rover
func (r *Rover) ProcessCommand(command string, marsPlateau *Plateau) {
	for _, cmd := range strings.Split(strings.ToUpper(command), "") {
		if r.isValidHeadCommand(cmd) {
			r.moveHead(cmd)
		} else if r.isValidMovementCommand(cmd) {
			r.moveRover(marsPlateau)
		}
	}
}

// GetX -
func (r *Rover) GetX() int {
	return r.x
}

// GetY -
func (r *Rover) GetY() int {
	return r.y
}

// DescribeState returns rover state in the format positionX positionY positionHead
func (r *Rover) DescribeState() string {
	return fmt.Sprintf("%d %d %s", r.x, r.y, r.head)
}

// GetPosition returns coordinates X and Y only in x,y format
func (r *Rover) GetPosition() string {
	return fmt.Sprintf("%d,%d", r.x, r.y)
}

// isValidHeadCommand validates if the input is valid or not
func (r *Rover) isValidHeadCommand(command string) bool {
	return command == "L" || command == "R"
}

// isValidMovementCommand validates if the input is valid or not
func (r *Rover) isValidMovementCommand(command string) bool {
	return command == "M"
}

// moveHead checks for all the valid movements for the head according to its current state and the new direction
func (r *Rover) moveHead(command string) {
	switch r.head {
	case "N":
		switch command {
		case "L":
			r.head = "W"
		case "R":
			r.head = "E"
		}
	case "E":
		switch command {
		case "L":
			r.head = "N"
		case "R":
			r.head = "S"
		}
	case "S":
		switch command {
		case "L":
			r.head = "E"
		case "R":
			r.head = "W"
		}
	case "W":
		switch command {
		case "L":
			r.head = "S"
		case "R":
			r.head = "N"
		}
	}
}

// moveRover changes the position of the rover according to the position of the head
// It changes the coordinates if the command produces valid coordinates within the plateau limits
func (r *Rover) moveRover(marsPlateau *Plateau) bool {
	validMovement := false

	switch r.head {
	case "N":
		newY := r.y + 1
		validMovement = marsPlateau.IsValidPairCoordinates(r.x, newY)
		if validMovement {
			r.y = newY
		}
	case "E":
		newX := r.x + 1
		validMovement = marsPlateau.IsValidPairCoordinates(newX, r.y)
		if validMovement {
			r.x = newX
		}
	case "S":
		newY := r.y - 1
		validMovement = marsPlateau.IsValidPairCoordinates(r.x, newY)
		if validMovement {
			r.y = newY
		}
	case "W":
		newX := r.x - 1
		validMovement = marsPlateau.IsValidPairCoordinates(newX, r.y)
		if validMovement {
			r.x = newX
		}
	}

	return validMovement
}
