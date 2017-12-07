
// Struct and methods for Bacterial DNA - Saideep Gona

/*

The DNA can be thought of as a series of slices. Each slice is a "gene" of sorts. The
"phenotypes" we choose to define rely on values derived from sampling from DNA genes.
A single gene can be a sampling source for one or more phenotypes. Likewise, a single
phenotype can sample from one or more genes. When a gene "serves" more then one phenotype,
this means that these phenotypes share some dependence.

*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math/rand"
	"math"
)

// ********************************************************* DNA Methods and Related Functions *********************************************************************************************

// ----------------- Read DNA from File and Build Genome Template --------------------------

func MakeNewDNA() DNA {

	// Creates brand new DNA object using the readin file format
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error with accessing current working directory")
	}
	filePath := wd + "/DNA_Blueprint.txt"
	dnaFile := ReadDNAFile(filePath)
	newDNA := BuildDNA(dnaFile)
	return newDNA
}

func ReadDNAFile(filename string) []string {

	// TAKEN DIRECTLY FROM HW5 WRITEUP
	// Opens a text file and creates a line-by-line slice of the contents

	in, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: couldn’t read in the DNA template")
		os.Exit(1)
	}
	// create the variable to hold the lines
	var lines []string = make([]string, 0)
	// for every line in the file
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
	// append it to the lines slice
		lines = append(lines, scanner.Text())
	}

	// check that all went ok
	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(1)
	}
	// close the file and return the lines
	in.Close()

	return lines
}

func BuildDNA(fileLines []string) DNA {

	// Converts the line by line file into a valid DNA object

	var fullDNA DNA
	fullDNA.phenotypes = make(map[string]Phenotype)
	fullDNA.genome = make(map[string]Gene)
	fullDNA.edges = make(map[string]Edge)

	// Reads file contencts and converts them to a properly portioned DNA object
	currentState := "configuration"
	for i := 1; i < len(fileLines); i++ {

		if fileLines[i] == "" {
			continue
		}
		if fileLines[i] == "Genes"{											// Check if input type is changing, and move to next line if so
			currentState = "genes"
			continue
		}
		if fileLines[i] == "Edges"{
			currentState = "edges"
			continue
		}
		if fileLines[i] == "Phenotypes"{
			currentState = "phenotypes"
			continue
		}

		if currentState == "configuration" {								// Check current read-in state while traversing lines
			currentLine := strings.Split(fileLines[i], "=")
			fullDNA.MakeConfig(currentLine)
		}
		if currentState == "phenotypes" {
			currentLine := strings.Split(fileLines[i], ",")
			fullDNA.MakePhen(currentLine)
		}
		if currentState == "genes" {
			currentLine := strings.Split(fileLines[i], ",")
			fullDNA.MakeGene(currentLine)
		}
		if currentState == "edges" {
			currentLine := strings.Split(fileLines[i], ",")
			fullDNA.MakeEdge(currentLine)
		}
	}
	return fullDNA
}

func (dna *DNA) MakeConfig(currentLine []string) {

	configVal, err := strconv.ParseFloat(currentLine[1], 64)
	if err != nil {
		fmt.Println("Error: Config value cannot be read")
		os.Exit(1)
	}

	if currentLine[0] == "Mutation Rate" {
		dna.mutRate = configVal
	}
	if currentLine[0] == "Mutation Magnitude" {
		dna.mutMagnitude = configVal
	}
	if currentLine[0] == "Low Bound" {
		dna.boundsLow = configVal
	}
	if currentLine[0] == "High Bound" {
		dna.boundsHigh = configVal
	}
	if currentLine[0] == "Gene Size" {
		dna.geneSize = int(configVal)
	}
	if currentLine[0] == "Sample Size" {
		dna.sampleSize = int(configVal)
	}
	if currentLine[0] == "LK Size" {
		dna.lksize = int(configVal)
	}

}

// MakeXXX DNA methods construct parts of the DNA (phenotypes, genes, edges)

func (dna *DNA) MakePhen(currentLine []string) {

	var phenotype Phenotype
	phenName := currentLine[0]
	phenotype.aggFunction = currentLine[1]
	phenotype.aggFuncArgs = currentLine[2:]
	dna.phenotypes[phenName] = phenotype

}

func (dna *DNA) MakeGene(currentLine []string) {

	if currentLine[1] == "normal" {
		geneValues := make([]float64, dna.geneSize)
		var gene Gene
		gene.values = geneValues
		geneName := currentLine[0]
		dna.genome[geneName] = gene
	} else if currentLine[1] == "lk" {
		geneValues := make([]float64, dna.lksize)
		var gene Gene
		gene.values = geneValues
		geneName := currentLine[0]
		dna.genome[geneName] = gene
	}
}

func (dna *DNA) MakeEdge(currentLine []string) {

	if len(currentLine) != 4 {
		fmt.Println("Wrong number of arguments for edge")
		os.Exit(1)
	}
	edgeWeight, err := strconv.ParseFloat(currentLine[3], 64)
	if err != nil {
		fmt.Println("Error: Edge weight not convertable to float")
		os.Exit(1)
	}
	var edge Edge
	edgeName := currentLine[0]
	edge.phenotype = currentLine[1]
	edge.gene = currentLine[2]
	edge.weight = edgeWeight

	dna.edges[edgeName] = edge

}

// ----------------- MUTATING THE DNA --------------------------

func (p *Petri) MutateAll() {
	for i := 0; i < len(p.allBacteria); i ++ {
		p.allBacteria[i].MutateDNA()
	//	fmt.Println("bactCheck",i)
	//	fmt.Println(p.allBacteria[0])
	}
}

func (bact *Bacteria) MutateDNA() {

	/*
	Given a dna object, mutates all the genes at once by calling a genome mutate method.
	*/
	for gene,_ := range bact.dna.genome {
		if gene != "Lock" && gene != "Key" {
			new := bact.dna.genome[gene]
			//fmt.Println(new)
			new.Mutate(bact.dna.mutRate, bact.dna.mutMagnitude,bact.dna.boundsLow,bact.dna.boundsHigh)
			bact.dna.genome[gene] = new
		} else {
			new := bact.dna.genome[gene]
			new.MutateLK(bact.dna.mutRate)
			bact.dna.genome[gene] = new
		}
	}
}

func (gene *Gene) MutateLK(mutationRate float64) {

	/*
	Mutates either the lock gene or the key gene. Consists of an equal probability
	cycle -5 to 5 where adding to 5 yields -5 and subtracting from -5 yields 5
	*/

	mutateRoll := rand.Float64()
	if mutateRoll < mutationRate/5.0 {

		directionRoll := rand.Intn(2)
		if directionRoll == 0 {
			if gene.values[0] == 5.0 {
				gene.values[0] = 4.0
			} else if gene.values[0] == -5 {
				gene.values[0] =5.0
			} else{
				gene.values[0] -= 1.0
			}
		}
		if directionRoll == 1 {
			if gene.values[0] == 5.0 {
				gene.values[0] = -5.0
			} else if gene.values[0] == -5 {
				gene.values[0] =-4.0
			} else{
				gene.values[0] -= 1.0
			}
		}
	}
}

func (gene *Gene) Mutate(mutationRate, mutationMagnitude, low, high float64) {

	/*
	Mutates input gene via pointer
	*/


	for i := 0; i < len(gene.values); i ++ {				// Loop through all values for gene
		newRoll := rand.Float64()					// Roll to see if mutation occurs
		//fmt.Println(newRoll)
		if newRoll < mutationRate {
			directionRoll := rand.Intn(2)			// Roll to see if mutation is positive or negative
			if directionRoll == 0 {
				gene.values[i] += mutationMagnitude
				if gene.values[i] > high {
					gene.values[i] -= mutationMagnitude/2.0
				}
			} else {
				gene.values[i] -= mutationMagnitude
				if gene.values[i] < low {
					gene.values[i] += mutationMagnitude/2.0
				}
			}
		}
	}
}

// ----------------- END MUTATE DNA ------------------------------

// ----------------- SAMPLING METHODS ----------------------------

func (p *Petri) AllPhenotypeExpectation(phenotypeName string) float64 {
	sum := 0.0
	for i := 0; i < len(p.allBacteria); i ++ {
		currentExp := p.allBacteria[i].dna.PhenotypeExpectation(phenotypeName)
		sum += currentExp
	}

	aveExp := sum/float64(len(p.allBacteria))
	LogExp := Logistic(aveExp, p.allBacteria[0].dna.phenotypes["EE"].aggFuncArgs)

	return LogExp
}

func (dna *DNA) PhenotypeExpectation(phenotypeName string) float64 {

	weightedExp := 0.0

	for _, edge := range dna.edges {
		if edge.phenotype == phenotypeName{
			geneMean := Mean(dna.genome[edge.gene].values)
			weight := edge.weight
			weightedExp += weight*geneMean
		}
	}
	return weightedExp
}

func (dna *DNA) PhenotypeAverage(phenotypeName string) float64 {

	/*
	Conducts sampling from all genes associated with a phenotype
	*/

	weightedSum := 0.0
	for _, edge := range dna.edges {
		if edge.phenotype == phenotypeName{
			newSample := dna.SampleGene(edge.gene)
			sampleMean := Mean(newSample)
			weight := edge.weight
			weightedSum += weight*sampleMean
		}
	}

	//fmt.Println("weighted", weightedSum)
	return weightedSum
}


func (dna *DNA) SampleGene(geneName string) []float64 {

	/*
	Given a gene name samples from the gene and returns the raw sample result
	*/

	randIndex := rand.Perm(dna.geneSize)
	sampleSlice := make([]float64, 0)

	for i := 0; i < dna.sampleSize; i ++ {

		sampleSlice = append(sampleSlice, dna.genome[geneName].values[randIndex[i]])

	}
	return sampleSlice
}

func Mean(list []float64) float64 {

	// Calculates mean from a slice of floats

	var sum float64

	for i := 0; i < len(list); i++{
		sum += list[i]
	}
	return sum/float64(len(list))
}

func Logistic(inputVal float64, arguments []string) float64{

	// Passes an input into a logistic function as well as arguments for the function and returns the output.

	floatArgs := make([]float64,0)

	if len(arguments) != 3 {
		fmt.Println("Wrong number of arguments to logistic function")
	}

	for _,arg := range arguments {
		argVal, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println("Error: Arg value for logistic cannot be read")
			os.Exit(1)
		}
		floatArgs = append(floatArgs, argVal)
	}

	max := floatArgs[0]
	steepness := floatArgs[1]
	midpoint := floatArgs[2]

	output := max/(1.0 + math.Exp(((-1)*steepness)*(inputVal-midpoint)))
	return output

}

// ----------------- END SAMPLING METHODS ------------------------

// ----------------- DNA PLOTTING METHODS ------------------------
/*
func AnimatePhenotypes(phenMap map[string][]float64) {

	// Animates the average values of all phenotypes over time

	animationImages := make([]image.Image,0)

	phenotypeList := make([]string, 0)
	phenotypeProgressions := make([][]float64, 0)

	for phen, list := range phenMap {									// Converts phenotype map into a slice of phenotypes and corresponding data progression
		phenotypeList = append(phenotypeList, phen)
		phenotypeProgressions := append(phenotypeProgressions, list)
	}

	numSteps := len(phenotypeProgressions[0])



}
*/
