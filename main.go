package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/alecthomas/kingpin"

	"github.com/memotoro/mars-rovers-challenge/models"
)

var (
	plateauMaxX   = kingpin.Flag("plateau-max-x", "Maximum coordinate X of the plateau").Default("5").Int()
	plateauMaxY   = kingpin.Flag("plateau-max-y", "Maximum coordinate Y of the plateau").Default("5").Int()
	roverCommands = kingpin.Flag("rover-commands", "Rover and command X,Y,HEAD,COMMAND separated by hyphen -").Default("1,2,N,LMLMLMLMM-3,3,E,MMRMMRMRRM").String()
)

func main() {
	kingpin.Parse()

	marsPlateau := models.NewPlateau(*plateauMaxX, *plateauMaxY)
	if marsPlateau == nil {
		log.Fatal("Invalid plateau coordinates")
	}
	log.Printf("Plateau %v", marsPlateau.GetDimensions())

	commands := strings.Split(*roverCommands, "-")

	instructions := make([]*models.Instruction, 0)

	for _, cmd := range commands {
		data := strings.Split(cmd, ",")
		posX, _ := strconv.ParseInt(data[0], 10, 64)
		posY, _ := strconv.ParseInt(data[1], 10, 64)
		head := data[2]
		movements := data[3]

		rover := models.NewRover(int(posX), int(posY), head)

		instruction := models.Instruction{Rover: rover, Command: movements}
		instructions = append(instructions, &instruction)
	}

	for index := range instructions {
		instruction := instructions[index]

		deployed := marsPlateau.DeployRover(instruction.Rover)
		if !deployed {
			continue
		}

		log.Printf("Rover Deployed %v", instruction.Rover.DescribeState())

		instruction.Rover.ProcessCommand(instruction.Command, marsPlateau)

		log.Printf("Rover Final State %v", instruction.Rover.DescribeState())
	}
}
