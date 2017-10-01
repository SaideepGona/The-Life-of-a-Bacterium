// Saideep Gona, Programming for Scientist's HW 2
// Due September 20, 2017

package main

import (
	"fmt"
	"math"
	"math/rand"
	//"os"
	//"strconv"
)

func RandWalk(width, height, stepSize float64, numberSteps int) float64 {

	// Workhorse function. Simulates all steps, prints out appropriate statements and calculates the distance

	// Inputs
	//	width: Max width of the field
	//	height: Max height of the field
	//	stepSize: Total distance covered for each step
	//	numberSteps: Number of steps to be taken
	// Outputs
	//	distance: Distance from starting point to end point after walk

	x := width/2.0
	y := height/2.0
	fmt.Println(x, y)	// Start

	for step := 0; step < numberSteps; step++ {

		x, y = RandomStep(x , y, stepSize, width, height)
		fmt.Println(x, y)	// Position after each step
	}

	distance := Distance(width/2.0, height/2.0, x, y)
	fmt.Println("Distance =", distance) //traversed distance

	return distance

}

func RandomStep(oldX, oldY, stepSize, w, h float64) (float64, float64) {

	// Determines the location of the next step in the walk

	// Inputs
	//	oldX: x coordinate of position before step
	//	oldY: y coordinate of position before step
	//	stepSize: Total distance covered per step
	//	w: width of field
	//	h: height of field
	// Outputs
	//	newX: x position after step
	//	newY: y position after step


	newX := oldX
	newY := oldY

	for (newX == oldX && newY == oldY) || !InField(newX, newY, w, h) {	// Makes sure that new position is different from start or that new position is in-bounds
																		// in order to exit while loop	
		deltaX, deltaY :=  ThetaGen()
		newX = oldX + stepSize * deltaX
		newY = oldY  + stepSize * deltaY

	}

	return newX, newY
}

func InField(x, y, w, h float64) bool {

	// Checks if given input position is in the given bounds.

	// Inputs
	//	x: x input position
	//	y: y input position
	//	w: width bound
	//	h: height bound
	// Output
	//	A boolean. True if position is legal. False if illegal.

	if x < 0 || x > w || y < 0 || y > h {
		return false
	} 

	return true

}

func ThetaGen() (float64, float64) {

	// Generates a random theta [0,2pi] and returns x and y components of it

	theta := rand.Float64()
	theta = theta * math.Pi * 2.0

	return math.Sin(theta), math.Cos(theta)

}

func Distance(sX, sY, eX, eY float64) float64 {

	// Calculates the distance between two points 

	// Inputs
	//	{sX, sY}: start point
	//	{eX, eY}: end point
	// Output
	//	Returns the distance between points

	return math.Sqrt(math.Pow((eX-sX),2) + math.Pow((eY-sY), 2))

}
/*
func main () {

	// Main handles incorrect command line input, seeding rand, and running RandWalk

	if len(os.Args) > 6 {  // Too many args, prints and exits
		fmt.Println("Error: Too many command line arguments")
		os.Exit(1)
	}

	if len(os.Args) < 6 {	// Too few args, prints and exits
		fmt.Println("Error: Too few command line arguments")
		os.Exit(1)
	}

	// If strconv parsing functions return errors, print message and exits program

	osWidth, err1 := strconv.ParseFloat(os.Args[1], 64)	
	if err1 != nil || osWidth <= 0 {
		fmt.Println("Error: Check the first argument. Should be a positive real number pertaining to width")
		os.Exit(1)
	}

	osHeight, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil || osHeight <= 0{
		fmt.Println("Error: Check the second argument. Should be a positive real number pertaining to height")
		os.Exit(1)
	}

	osStepSize, err3 := strconv.ParseFloat(os.Args[3], 64)
	if err3 != nil {
		fmt.Println("Error: Check the third argument. Should be a positive real number pertaining to the step size")
		os.Exit(1)
	}

	osNumberSteps, err4 := strconv.Atoi(os.Args[4])
	if err4 != nil || osNumberSteps < 1 {
		fmt.Println("Error: Check the fourth argument. Should be a positive, non-zero integer pertaining to the number of steps")
		os.Exit(1)
	}

	osSeed, err5 := strconv.ParseInt(os.Args[5], 10, 64)
	if err5 != nil {
		fmt.Println("Error: Check the fifth argument. Should be an integer for seeding")
		os.Exit(1)
	}

	rand.Seed(osSeed)

	RandWalk(osWidth, osHeight, osStepSize, osNumberSteps)

}
*/