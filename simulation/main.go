// Code by Kwanho Kim

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // rand.Seed(time.Now().UTC().UnixNano())

	// Takes in 8 inputs as parameters
	if len(os.Args) != 9 {
		fmt.Println("Error: number of command line parameter!")
		os.Exit(1)
	}

	radius, err1 := strconv.ParseFloat(os.Args[1], 64) // default 200
	if err1 != nil {                                   // there was a problem
		fmt.Println("Error: radius is not a number!")
		os.Exit(1)
	} else if radius < 0 {
		fmt.Println("Error: raidus must be positive!")
		os.Exit(1)
	}

	numInitialBact, err2 := strconv.Atoi(os.Args[2]) // default 3, max 5
	if err2 != nil {                                 // there was a problem
		fmt.Println("Error: numInitialBact is not an integer")
		os.Exit(1)
	} else if numInitialBact < 0 {
		fmt.Println("Error: numInitialBact must be positive!")
		os.Exit(1)
	}

	numIteration, err3 := strconv.Atoi(os.Args[3]) // default 200
	if err3 != nil {                               // there was a problem
		fmt.Println("Error: numIteration is not an integer")
		os.Exit(1)
	} else if numIteration < 0 {
		fmt.Println("Error: numIteration must be positive!")
		os.Exit(1)
	}

	drugIntro, err4 := strconv.Atoi(os.Args[4]) // default 30
	if err4 != nil {                            // there was a problem
		fmt.Println("Error: drugIntro is not an integer")
		os.Exit(1)
	} else if drugIntro < 0 {
		fmt.Println("Error: drugIntro must be positive!")
		os.Exit(1)
	}

	predatorIntro, err5 := strconv.Atoi(os.Args[5]) // default 50
	if err5 != nil {                                // there was a problem
		fmt.Println("Error: predatorIntro is not an integer")
		os.Exit(1)
	} else if predatorIntro < 0 {
		fmt.Println("Error: predatorIntro must be positive!")
		os.Exit(1)
	}

	numPred, err6 := strconv.Atoi(os.Args[6]) // default 2
	if err6 != nil {                          // there was a problem
		fmt.Println("Error: numPred is not an integer")
		os.Exit(1)
	} else if numPred < 0 {
		fmt.Println("Error: numPred must be positive!")
		os.Exit(1)
	}

	energyContent, err7 := strconv.ParseFloat(os.Args[7], 64) // default 13
	if err7 != nil {                                          // there was a problem
		fmt.Println("Error: energyContent is not a number")
		os.Exit(1)
	} else if energyContent < 0 {
		fmt.Println("Error: energyContent must be positive!")
		os.Exit(1)
	}

	basalMetabolicRate, err8 := strconv.ParseFloat(os.Args[8], 64) // default .5
	if err8 != nil {                                               // there was a problem
		fmt.Println("Error: basalMetabolicRate is not a number")
		os.Exit(1)
	} else if basalMetabolicRate < 0 {
		fmt.Println("Error: basalMetabolicRate must be positive!")
		os.Exit(1)
	}

	var p Petri
	p.radius = radius
	/*
	   //----------------------- Default settings ----------------------
	     radius := 200
	     numInitialBact := 5
	     numIteration := 200
	     drugIntro := 30
	     predatorIntro := 60
	     numPred := 2
	     energyContent := 13.0
	     basalMetabolicRate := 0.5
	   //---------------------------------------------------------------
	*/
	p.InitializeBact(numInitialBact)
	p.InitializeFoodpackage(5000, energyContent)
	// p.InitializePredator(numInitialPred)
	// p.InitializeDrugpackage(50, 5, 50)
	gifImages := p.AnimationPetri(numInitialBact, numIteration, drugIntro, predatorIntro, numPred, basalMetabolicRate)
	Process(gifImages, "data/simulation")
}
