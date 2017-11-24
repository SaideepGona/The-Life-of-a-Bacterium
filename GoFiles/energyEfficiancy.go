// Saideep Gona
// Methods related to energy efficiancy
package main

func (bact *Bacteria) UpdateEE() {

	// Updates energy efficiency

	newEERaw := bact.dna.PhenotypeAverage("EE")
	newEE := Logistic(newEERaw, bact.dna.phenotype.aggFuncArgs)
	bact.dna.energyEfficiancy = newEE

}

func (p *Petri) UpdateAllEE() {

	// Updates the energy efficiency for all bacteria in petri dish

	for index, bacteria := range p.allBacteria {
		bacteria.UpdateEE()
	}
}

