// Saideep Gona
// Methods related to energy efficiancy

package main

func (dna *DNA) UpdateEE() {

	phenotype = "Energy Efficiency"
	newEERaw := dna.PhenotypeAverage(phenotype)
	newEE := Logistic(newEERaw, dna.phenotype.)
	dna.energyEfficiancy = newEE

}

func (bact *Bacteria) BurnEnergy(phenotype string) {

	if phenotype == "movement" {
		phenVal := bact.movement
		scale := 1.0
	}

	if phenotype == "antibiotic" {
		phenVal := bact.potency
		scale := 1.0
	}

	burnAmount := bact.energyEfficiancy * phenVal * scale
	bact.currentEnergy = bact.currentEnergy - burnAmount
}
