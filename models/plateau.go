package models

import (
	"fmt"
)

// Plateau - defines plateau attributes
type Plateau struct {
	minX   int
	minY   int
	maxX   int
	maxY   int
	rovers []*Rover
}

// NewPlateau constructs a new plateau
func NewPlateau(maxX, maxY int) *Plateau {
	// Creates a plateau only with positive coordinates for the upper-right corner
	if maxX <= 0 || maxY <= 0 {
		return nil
	}

	return &Plateau{minX: 0, minY: 0, maxX: maxX, maxY: maxY}
}

// DeployRover adds a new rover to the plateau
func (p *Plateau) DeployRover(rover *Rover) bool {
	if p.IsPositionAvailable(rover.GetX(), rover.GetY()) {
		p.rovers = append(p.rovers, rover)
		return true
	}

	return false
}

// GetDimensions gets dimensions
func (p *Plateau) GetDimensions() string {
	return fmt.Sprintf("(%d,%d) (%d,%d)", p.minX, p.minY, p.maxX, p.maxY)
}

// NumberOfRovers returns the total number of rovers deployed
func (p *Plateau) NumberOfRovers() int {
	return len(p.rovers)
}

// IsPositionAvailable checks if the position is not occupied by another rover in the plateau
func (p *Plateau) IsPositionAvailable(x, y int) bool {
	for index := range p.rovers {
		rover := p.rovers[index]
		// If any other rover has the same position
		if rover.GetPosition() == fmt.Sprintf("%d,%d", x, y) {
			return false
		}
	}

	return true
}

// IsValidPairCoordinates validates if a set of coordinates x,y are valid according to the plateau dimensions
func (p *Plateau) IsValidPairCoordinates(x, y int) bool {
	return p.IsPositionAvailable(x, y) && x >= p.minX && x <= p.maxX && y >= p.minY && y <= p.maxY
}
