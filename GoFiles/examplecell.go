package main

import (

	"fmt"
	"math/rand"
	"math"
)

func main () {

	rand.Seed(10)
	width := 100.0
	height := 100.0
	timesIter := 100
	SingleCellSim(width, height, timesIter)

}

func SingleCellSim (w,h float64, time int) {

	var cellDNA = []float64{1.0,5.0,10.0,2.0}
	
	x := 50.0
	y := 50.0

	for round:=0; round < time; round++ {

		stepSize := CalcMovement(cellDNA)
		x, y = RandomStep(x , y, stepSize, w, h)
		fmt.Println("Position:",x,y)
		fmt.Println("DNA:", cellDNA)
		MutateDNA(cellDNA)
	}

}

func MutateDNA(cellDNA []float64) {

	for i := range(cellDNA) {
		randVal := rand.Intn(100)
		if randVal < 25 {
			cellDNA[i] = cellDNA[i] - 1
		} else if randVal > 75 {
			cellDNA[i] = cellDNA[i] + 1
		}
	}

}

func CalcMovement(cellDNA []float64) float64 {

	sum := 0.0

	for _, val := range cellDNA {
		sum += val
	}

	average := sum/float64(len(cellDNA))
	fmt.Println("average:", average)
	maxBound := 10.0
	steepness := 1.0
	midpoint := 5.0
	logistTransform := maxBound / (1+math.Exp(-steepness * (average-midpoint)))
	fmt.Println("final movement:", logistTransform)

	return logistTransform

}
/*
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

*/