
// Struct and methods for Bacterial DNA

/*

The DNA can be thought of as a series of slices. Each slice is a "gene" of sorts. The 
"phenotypes" we choose to define rely on values derived from sampling from DNA genes. 
A single gene can be a sampling source for one or more phenotypes. Likewise, a single 
phenotype can sample from one or more genes. When a gene "serves" more then one phenotype,
this means that these phenotypes share some dependence.

*/

package main

type Phenotype struct {							// A phenotype and associated aggregate function information
	name string
	aggFunction string
	AggFuncArgs []int
}

type Edge struct {								// An edge defined by endpoints and with an edge function/arguments
	phenotype string
	gene string
	edgeFunction string
	edgeFuncArgs []int
}

type Genome map[string][]float64				// Genome with genome names and corresponding numerical slices

type DNA struct {
	phenotypes []Phenotype						// Contains all phenotypes the DNA "controls"
	edges []Edge								// Contains the edges from phenotype to gene which determine how phenotypes are expressed
	genome Genome								// Stores all the genes and current gene values in the bacterial genome
	mutRate float64								// Represents a probability of mutation
	mutMagnitude float64						// If a mutation occurs, is a benchmark for the magnitude of mutation
	bounds [2]float64							// Represents some bounds on the values individual gene elements can take
	geneSize int								// Represents the length of each gene
	sampleCount int								// Represents the number of samples chosen during a selection event per gene 
}

// ********************************************************* DNA Methods and Related Functions *********************************************************************************************

// ----------------- Read DNA from File --------------------------

func ReadInDNA() DNA {

	/*
	Read in from a DNA file to build the archetypal DNA struct
	*/

}

// ----------------- MUTATING THE DNA --------------------------

func (dna *DNA) MutateDNA() {

	/*
	Given a dna object, mutates all the genes at once by calling a genome mutate method.
	*/

	for gene := range dna.genome {
		Mutate(dna.genome[gene], dna.mutRate, dna.mutMagnitude)
	}
}

	func Mutate (gene *[]float64, mutationRate float64, mutationMagnitude) {

		/*
		Mutates input genome via pointer
		*/

		for i := 0; i < len(gene); i ++ {				// Loop through all values for gene

			newRoll := rand.Float64()					// Roll to see if mutation occurs

			if newRoll < mutationRate {

				directionRoll := rand.Int(1)			// Roll to see if mutation is positive or negative

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

func (dna *DNA) PhenotypeSample(phenotypeName string) []float64 {

	/*
	Conducts sampling from all genes associated with a phenotype
	*/

}

func (dna *DNA) SampleGene(geneName string) []float64 {

	/*
	Given a gene name samples from the gene and returns the raw sample result
	*/

	randIndex := rand.Perm(0, dna.geneSize)

	sampleSlice := make([]float64, 0)

	for i := 0; i < dna.sampleCount; i ++ {

		sampleSlice = append(sampleSlice, dna.genome[geneName][randIndex[i]])

	}

	return sampleSlice

} 

// ----------------- EDGE FUNCTION LIBRARY ----------------------------			> 	This is for functions that act upon the output of a single-gene sample