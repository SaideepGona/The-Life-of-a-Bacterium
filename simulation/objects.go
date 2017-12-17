// Combined Effort

package main

type Petri struct {
	radius      float64
	allBacteria []Bacteria
	allFoodpack []Foodpackage
	allPredator []Predator
	allDrugpack []Drugpackage
	allPredKill []PredatorKiller
}

type Bacteria struct {
	position         Coords
	currentEnergy    float64
	energyCapacity   float64
	stepSize         float64
	energyEfficiency float64
	sizeRadius       float64
	repEnergy        float64
	attackRange      float64
	ABenzyme         ABenzyme
	resistEnzyme     ResistEnzyme
	dna              DNA
	hasRep           bool
	strain           int
}

type Foodpackage struct {
	position Coords
	energy   float64
}

type Predator struct {
	position         Coords
	sizeRadius       float64
	currentEnergy    float64
	energyEfficiency float64
	energyCapacity   float64
	stepSize         float64
	repEnergy        float64
	hasRep           bool
}

type PredatorKiller struct {
	position Coords
	potency  float64
}

type Drugpackage struct {
	position Coords
	lock     int
	potency  float64
}

type Coords struct {
	coorX float64
	coorY float64
}

// ------------------ Bacteria enzymes --------------------

type ResistEnzyme struct {
	key     int
	potency float64
}

type ABenzyme struct {
	lock    int
	potency float64
}

//------------------ DNA related objects ------------------------

type Phenotype struct { // A phenotype and associated aggregate function information
	aggFunction string
	aggFuncArgs []string
	edges       []Edge
}

type Edge struct { // An edge defined by endpoints and with an edge function/arguments
	phenotype string
	gene      string
	weight    float64
}

type Gene struct {
	values []float64
} // Genome with gene names and corresponding numerical slices

type DNA struct {
	phenotypes   map[string]Phenotype // Contains all phenotypes the DNA "controls"
	edges        map[string]Edge      // Contains the edges from phenotype to gene which determine how phenotypes are expressed
	genome       map[string]Gene      // Stores all the genes and current gene values in the bacterial genome
	mutRate      float64              // Represents a probability of mutation
	mutMagnitude float64              // If a mutation occurs, is a benchmark for the magnitude of mutation
	boundsLow    float64              // Represents some bounds on the values individual gene elements can take
	boundsHigh   float64
	geneSize     int // Represents the length of each gene
	sampleSize   int
	lksize       int // Represents the number of samples chosen during a selection event per gene
}
