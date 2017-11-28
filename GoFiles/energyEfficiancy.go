// Saideep Gona
// Methods related to energy efficiancy
package main

func (bact *Bacteria) UpdateEE() {

	// Updates energy efficiency

	newEERaw := bact.dna.PhenotypeAverage("EE")
	fmt.Println(newEERaw)
	newEE := Logistic(newEERaw, bact.dna.phenotypes["EE"].aggFuncArgs)
	bact.energyEfficiency = newEE

}

func (p *Petri) UpdateAllEE() {

	// Updates the energy efficiency for all bacteria in petri dish

	for i :=0; i < len(p.allBacteria); i++ {
		fmt.Println(bacteria)
		p.allBacteria[i].UpdateEE()
	}
}

