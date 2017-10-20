// Struct and methods for Bacterial DNA

/*

The DNA can be thought of as a series of slices. Each slice is a "gene" of sorts. The 
"phenotypes" we choose to define rely on values derived from sampling from DNA genes. 
A single gene can be a sampling source for one or more phenotypes. Likewise, a single 
phenotype can sample from one or more genes. When a gene "serves" more then one phenotype,
this means that these phenotypes share some dependence.

*/

/* Bacterial DNA Struct
	code: Maps
*/

type PhenotypeToGene []string					// Represents a slice of phenotype names which are grouped because they 
												// sample from a single gene of interest
type Gene []float64								// A gene is just a slice of floats

type DNA struct {
	code map[PhenotypeToGene]Gene
	mutRate float64								// Represents a probability of mutation
	mutMagnitude float64						// If a mutation occurs, is a benchmark for the magnitude of mutation
	bounds [2]float64							// Represents some bounds on the values individual gene elements can take
	geneSize int								// Represents the length of each gene
	sampleCount int								// Represents the number of samples chosen during a selection event per gene 
}

type BluePrint struct {

}

// In order to construct a brand new genome, we need to have a BluePrint representing all the bacterial phenotypes as well as 
// how these phenotypes are related to one another. 