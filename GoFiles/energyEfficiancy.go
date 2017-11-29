// Saideep Gona
// Methods related to energy efficiancy
package main

import (

	"fmt"
)
func (bact *Bacteria) UpdateEE() {

	// Updates energy efficiency
	newEERaw := bact.dna.PhenotypeAverage("EE")
	newEE := Logistic(newEERaw, bact.dna.phenotypes["EE"].aggFuncArgs)
	bact.energyEfficiency = newEE

}

func (p *Petri) UpdateAllEE() {
	fmt.Println("UpdateAllEE")
	// Updates the energy efficiency for all bacteria in petri dish
	for i := 0; i < len(p.allBacteria); i++{
		p.allBacteria[i].UpdateEE()
	}
}
