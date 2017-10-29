package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

func main () {

	rand.Seed(time.Now().Unix())			// Random Seed
	width := 100.0							// Width of environment	
	height := 100.0							// Height of environment
	timesIter := 10000						// Number of iterations
	SingleCellSim(width, height, timesIter)

}

func SingleCellSim (w,h float64, time int) {

	// Simulates a single cell. Holds the DNA and position. Also creates a graph
	// Inputs:
	//	w: Width of container
	//	h: Height of container
	// Outputs:
	//	Creates some kind of visualization

	var cellDNA = []float64{0.0, 0.0, 0.0, 0.0}
	
	x := 50.0
	y := 50.0
	blue := MakeColor(0,0,255)

	movementGraph := CreateNewCanvas(time, 100)			// Plots the movement distance at each time point
	movementGraph.SetStrokeColor(blue)
	movementGraph.MoveTo(0,0)

	pathChart := CreateNewCanvas(5*int(w)+1, 5*int(h)+1)					// Plots path of cell over time in environment
	pathChart.SetStrokeColor(blue)
	pathChart.MoveTo(5.0*x,5.0*(h-y))

	stepSize := 0.0

	for round:=0; round < time; round++ {				// Iterates through time steps

		
		stepSize = CalcMovement(cellDNA)				// Modifies step size based on the DNA values
		movementGraph.MoveTo(float64(round), 100)
		movementGraph.LineTo(float64(round), 10.0*(10.0-stepSize))

		x, y = RandomStep(x , y, stepSize, w, h)		// Finds the next movement location
		pathChart.LineTo(5.0*x, 5.0*y)

		fmt.Println("Position:",x,y)
		fmt.Println("DNA:", cellDNA)
		MutateDNA(cellDNA)
	}

	movementGraph.Stroke()
	movementGraph.SaveToPNG("MovementGraph.PNG")

	pathChart.Stroke()
	pathChart.SaveToPNG("PathChart.PNG")

}

func MutateDNA(cellDNA []float64) {

	// Returns a "mutated" version of the current DNA 
	// Inputs: 
	//	cellDNA: Current DNA slice
	// Outputs:
	//	No direct output, modifies original slice in a random way	

	for i := range(cellDNA) {
		randVal := rand.Intn(100)

		if randVal < 25 {
			cellDNA[i] = cellDNA[i] - float64(rand.Intn(3) + 1)
			if cellDNA[i] < -20 {
				cellDNA[i] = cellDNA[i] + 1
			}
		} else if randVal > 75 {
			cellDNA[i] = cellDNA[i] + float64(rand.Intn(3) + 1)
			if cellDNA[i] > 20 {
				cellDNA[i] = cellDNA[i] - 1
			}
		}
	}

}

func CalcMovement(cellDNA []float64) float64 {

	// Given the current DNA slice, calculates what the movement value should be
	// Inputs: 
	//	cellDNA: Current DNA slice
	// Outputs:
	//	logistTransform: Final movement value (after logistic transform)	

	sum := 0.0

	for _, val := range cellDNA {
		sum += val
	}

	average := sum/float64(len(cellDNA))
	fmt.Println("average:", average)
	maxBound := 10.0
	steepness := 0.5
	midpoint := 0.0
	logistTransform := maxBound / (1+math.Exp(-steepness * (average-midpoint)))
	fmt.Println("final movement:", logistTransform)

	return logistTransform

}
