// Saideep Gona
// Methods related to energy efficiancy
package main

import (

	"fmt"
)
func (bact *Bacteria) UpdateEE() {

	// Updates energy efficiency

	newEERaw := bact.dna.PhenotypeAverage("EE")
	fmt.Println(newEERaw)
	newEE := Logistic(newEERaw, bact.dna.phenotypes["EE"].aggFuncArgs)
	bact.energyEfficiency = newEE

}

func (p *Petri) UpdateAllEE() {

	// Updates the energy efficiency for all bacteria in petri dish

	for _, bacteria := range p.allBacteria {
		bacteria.UpdateEE()
	}
}
