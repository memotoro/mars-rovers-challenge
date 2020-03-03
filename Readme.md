# Mars Rover Challenge

This exercise assumes that Go v1.13 is installed in the computer.

## Run the program

Create the following folder structure in the GOPATH of your computer.

```
mkdir -p $GOPATH/src/github.com/memotoro
```

Inside this new folder, download the project and locate the terminal in the root directory.

```
cd $GOPATH/src/github.com/memotoro/mars-rovers-challenge
```

The program comes with some default values to simulate a basic test case.

To check the program parameters run the following command.

```
go run main.go --help
```

The output should be similar to the following:

```
usage: main [<flags>]

Flags:
  --help             Show context-sensitive help (also try --help-long and --help-man).
  --plateau-max-x=5  Maximum coordinate X of the plateau
  --plateau-max-y=5  Maximum coordinate Y of the plateau
  --rover-commands="1,2,N,LMLMLMLMM-3,3,E,MMRMMRMRRM"  
                     Rover and command X,Y,HEAD,COMMAND separated by hyphen -
```

To run the program with the default values, run the following command:

```
go run main.go
```

The output should be similar to the following:

```
2020/03/02 16:56:13 Plateau (0,0) (5,5)
2020/03/02 16:56:13 Rover Deployed 1 2 N
2020/03/02 16:56:13 Rover Final State 1 3 N
2020/03/02 16:56:13 Rover Deployed 3 3 E
2020/03/02 16:56:13 Rover Final State 5 1 E
```

To run with different parameters, run a command similar to the following:

```
go run main.go --plateau-max-x 10 --plateau-max-y 10 --rover-commands 0,0,N,MMMMMRMMMM-1,1,E,LMRLMRLMRLMRR
```

The output should be similar to the following:

```
2020/03/02 16:56:29 Plateau (0,0) (10,10)
2020/03/02 16:56:29 Rover Deployed 0 0 N
2020/03/02 16:56:29 Rover Final State 4 5 E
2020/03/02 16:56:29 Rover Deployed 1 1 E
2020/03/02 16:56:29 Rover Final State 1 5 S
```

## Test files

There are test files for the plateau and the rovers

Follow the instructions to run the test cases.

Locate the terminal in the root directory of the project.

Then run the following command to run all test files for both plateau and rovers.

You could extend the test suite with more test cases to verify the completeness of the code.

```
go test ./... -v
```

Output should be similar to the following:

```
?   	github.com/memotoro/mars-rovers-challenge	[no test files]
=== RUN   Test_ValidPlateau
--- PASS: Test_ValidPlateau (0.00s)
=== RUN   Test_InvalidPlateau
--- PASS: Test_InvalidPlateau (0.00s)
=== RUN   Test_ValidCoords
--- PASS: Test_ValidCoords (0.00s)
=== RUN   Test_InvalidCoords
--- PASS: Test_InvalidCoords (0.00s)
=== RUN   Test_InvalidNegativeCoords
--- PASS: Test_InvalidNegativeCoords (0.00s)
=== RUN   Test_DeployRoversToPlateau
--- PASS: Test_DeployRoversToPlateau (0.00s)
=== RUN   Test_ValidPositionForRoverInPlateau
--- PASS: Test_ValidPositionForRoverInPlateau (0.00s)
=== RUN   Test_InvalidPositionForRoverInPlateau
--- PASS: Test_InvalidPositionForRoverInPlateau (0.00s)
=== RUN   Test_InvalidPositionForNewRoversInPlateu
--- PASS: Test_InvalidPositionForNewRoversInPlateu (0.00s)
=== RUN   Test_ValidCommandRoverA
--- PASS: Test_ValidCommandRoverA (0.00s)
=== RUN   Test_ValidCommandRoverB
--- PASS: Test_ValidCommandRoverB (0.00s)
=== RUN   Test_ValidCommandStuckUpperRightCorner
--- PASS: Test_ValidCommandStuckUpperRightCorner (0.00s)
=== RUN   Test_ValidCommandStuckEastEdge
--- PASS: Test_ValidCommandStuckEastEdge (0.00s)
=== RUN   Test_ValidCommandNoCoordsChange
--- PASS: Test_ValidCommandNoCoordsChange (0.00s)
=== RUN   Test_ValidPathsForTwoRovers
--- PASS: Test_ValidPathsForTwoRovers (0.00s)
=== RUN   Test_InvalidPathsForTwoRoversClash
--- PASS: Test_InvalidPathsForTwoRoversClash (0.00s)
PASS
ok  	github.com/memotoro/mars-rovers-challenge/models	0.005s
```
