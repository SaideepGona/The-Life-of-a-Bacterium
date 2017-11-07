
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
)

type Phenotype struct {							// A phenotype and associated aggregate function information
	aggFunction string
	aggFuncArgs []string
	edges []*Edge
}

type Edge struct {								// An edge defined by endpoints and with an edge function/arguments
	phenotype string
	gene string
	weight float64
}

type Gene []float64								// Genome with gene names and corresponding numerical slices

type DNA struct {
	phenotypes map[string]Phenotype				// Contains all phenotypes the DNA "controls"
	edges map[string]Edge				// Contains the edges from phenotype to gene which determine how phenotypes are expressed
	genome map[string]Gene						// Stores all the genes and current gene values in the bacterial genome
	mutRate float64								// Represents a probability of mutation
	mutMagnitude float64						// If a mutation occurs, is a benchmark for the magnitude of mutation
	boundsLow float64							// Represents some bounds on the values individual gene elements can take
	boundsHigh float64
	geneSize int								// Represents the length of each gene
	sampleSize int								// Represents the number of samples chosen during a selection event per gene 
}

// ********************************************************* DNA Methods and Related Functions *********************************************************************************************

// ----------------- Read DNA from File and Build Genome Template --------------------------

func MakeNewDNA() DNA {

	// Creates brand new DNA object using the readin file format

	filePath := os.Getwd() + "/../OtherFiles/DNA_blueprint.txt"
	dnaFile := ReadDNAFile(filePath)
	return BuildDNA(dnaFile)

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
			currentLine := strings.Split(fileLines[i], ":")
			fullDNA.makeConfig(currentLine)
		}
		if currentState == "phenotypes" {
			currentLine := strings.Split(fileLines[i], ",")
			fullDNA.makePhen(currentLine)
		}
		if currentState == "genes" {
			fullDNA.makeGene(fileLines[i])
		}
		if currentState == "edges" {
			currentLine := strings.Split(fileLines[i], ",")
			fullDNA.makeEdge(currentLine)
		}
	}
	return fullDNA
}

func (dna *DNA) makeConfig(currentLine []string) {

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
	if currentLine[0] == "BoundLow" {
		dna.boundsLow = configVal
	}
	if currentLine[0] == "BoundHigh" {
		dna.boundsHigh = configVal
	}
	if currentLine[0] == "Gene Size" {
		dna.geneSize = int(configVal)
	}
	if currentLine[0] == "Sample Size" {
		dna.sampleSize = int(configVal)
	}

}

func (dna *DNA) makePhen(currentLine []string) {									// makeXXX DNA methods construct parts of the DNA (phenotypes, genes, edges)

	var phenotype Phenotype
	phenName := currentLine[0]
	phenotype.aggFunction = currentLine[1]
	phenotype.aggFuncArgs = currentLine[2:]
	dna.phenotypes[phenName] = phenotype

}

func (dna *DNA) makeGene(currentLine string) {

	gene := make(Gene, dna.geneSize)
	geneName := currentLine
	dna.genome[geneName] = gene
}

func (dna *DNA) makeEdge(currentLine []string) {

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

func (dna *DNA) MutateDNA() {

	/*
	Given a dna object, mutates all the genes at once by calling a genome mutate method.
	*/

	for gene := range dna.genome {
		dna.genome[gene].Mutate(dna.mutRate, dna.mutMagnitude)
	}
}

func (gene Gene) Mutate (mutationRate, mutationMagnitude float64) {


	/*
	Mutates input genome via pointer
	*/

	for i := 0; i < len(gene); i ++ {				// Loop through all values for gene

		newRoll := rand.Float64()					// Roll to see if mutation occurs

		if newRoll < mutationRate {

			directionRoll := rand.Intn(1)			// Roll to see if mutation is positive or negative

			if directionRoll == 0 {
				gene[i] += mutationMagnitude
			} else {
				gene[i] -= mutationMagnitude
			}

		}
	}
}

// ----------------- END MUTATE DNA ------------------------------

// ----------------- SAMPLING METHODS ----------------------------

func (dna *DNA) PhenotypeAverage(phenotypeName string) float64 {

	/*
	Conducts sampling from all genes associated with a phenotype
	*/

	weightedSum := 0.0
	edges := dna.phenotypes[phenotypeName].edges

	for i := 0; i < len(edges); i++ {

		newSample := dna.SampleGene(edges[i].gene)
		sampleMean := Mean(newSample)
		weight := edges.weight
		weightedSum += weight*sampleMean

	}

	return fullSample

}

func (dna *DNA) SampleGene(geneName string) []float64 {

	/*
	Given a gene name samples from the gene and returns the raw sample result
	*/

	randIndex := rand.Perm(dna.geneSize)

	sampleSlice := make([]float64, 0)

	for i := 0; i < dna.sampleSize; i ++ {

		sampleSlice = append(sampleSlice, dna.genome[geneName][randIndex[i]])

	}

	return sampleSlice

} 

// ----------------- PHENOTYPE FUNCTION LIBRARY ----------------------------			> 	This is for functions that act upon the output of a single-gene sample

func Mean(list []float64) float64 {

	var sum float64

	for i := 0; i < len(list); i++{
		sum += list[i]
	}
	return sum/float64(len(list))

}




