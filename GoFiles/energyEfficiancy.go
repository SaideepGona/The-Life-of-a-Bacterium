// Saideep Gona
// Methods related to energy efficiancy

package main

func (dna *DNA) UpdateEE() {

	// Updates energy efficiency

	newEERaw := dna.PhenotypeAverage("EE")
	newEE := Logistic(newEERaw, dna.phenotype.aggFuncArgs)
	dna.energyEfficiancy = newEE

}

func (p *Petri) UpdateAllEE() {

	// Updates the energy efficiency for all bacteria in petri dish

	for index, bacteria := range p.allBacteria {
		bacteria.dna.UpdateEE()
	}
}

