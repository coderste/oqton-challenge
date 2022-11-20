package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// You are facing north, you can either turn left (L) or right (R) 90 degrees and walk forward (F) the given
// number of blocks, ending at a new intersection.
//
// given:
// R1, R1, R3, R1, R1, L2, R5, L2, R5, R1, R4, L2, R3, L3, R4, L5, R4, R4, R1, L5, L4, R5, R3, L1,
// R4, R3, L2, L1, R3, L4, R3, L2, R5, R190, R3, R5, L5, L1, R54, L3, L4, L1, R4, R1, R3, L1, L1,
// R2, L2, R2, R5, L3, R4, R76, L3, R4, R191, R5, R5, L5, L40, L5, L3, R1, R3, R2, L2, L2, L4, L5,
// L4, R5, R4, R4, R2, R3, R4, L3, L2, R5, R3, L2, L1, R2, L3, R2, L1, L1, R1, L3, R5, L5, L1, L2,
// R5, R3, L3, R3, R5, R2, R5, R5, L5, L5, R25, L3, L5, L2, L1, R2, R2, L2, R2, L3, L2, R3, L5, R4,
// L4, L5, R3, L4, R1, R3, R2, R4, L2, L3, R2, L5, R5, R4, L2, R4, L1, L3, L1, L3, R1, R2, R1, L5,
// R5, R3, L3, L3, L2, R4, R2, L5, L1, L1, L5, L4, L1, L1, R1
//
// What is the shortest path to the finish?

// Direction represents the north or east direction
type Direction struct {
	North int
	East  int
}

// directions represents the north and east directions
// if we are going south or west we use a negative value to just move
// backwards in that direction
var directions = []Direction{
	{North: 1, East: 0},
	{North: 0, East: 1},
	{North: -1, East: 0},
	{North: 0, East: -1},
}

// Position represents how many steps we have to take
type Position struct {
	NorthSteps int
	EastSteps  int
}

// State represents the current state of the robot
type State struct {
	CurrentPosition Position
	FacingDirection int
}

// Turn will rotate us in the given direction
func (state *State) Turn(rotate int) *State {
	// Get the new direction
	state.FacingDirection = (state.FacingDirection + rotate + len(directions)) % len(directions)
	return state
}

// Walk will move the state forward the given number of steps
func (state *State) Walk(steps int) *State {
	// get current amount of north steps we have taken and add the new steps
	state.CurrentPosition.NorthSteps += directions[state.FacingDirection].North * steps

	// get the current amount of east steps we have taken and add the new steps
	state.CurrentPosition.EastSteps += directions[state.FacingDirection].East * steps

	// Return the updated state
	return state
}

func main() {
	input, _ := ioutil.ReadFile("./input.txt")

	state := State{}
	Distance(&state, string(input))

	var shortestDistance = Abs(state.CurrentPosition.NorthSteps) + Abs(state.CurrentPosition.EastSteps)

	// fmt.Printf("%+v\n", state)
	fmt.Printf("Shortest distance to location is %v\n", shortestDistance)
}

// distance will calculate the shortest distance to the finish
func Distance(state *State, instructions string) *State {
	// Split the instructions into a slice
	for _, instruction := range strings.Split(strings.TrimSpace(instructions), ", ") {
		// which direction are we turning in
		turn := string(instruction[0])

		// convert the amount of steps from string to int
		steps, _ := strconv.Atoi(instruction[1:])

		if turn == "R" {
			state = state.Turn(1)
		} else if turn == "L" {
			state = state.Turn(-1)
		}

		state.Walk(steps)
	}

	return state
}

// Abs returns the absolute value of x
func Abs(x int) int {
	// If x is negative then we need to make it positive
	if x < 0 {
		return -x
	}

	return x
}
