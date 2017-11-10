// Saideep Gona
// Methods related to energy efficiancy

func (dna *DNA) UpdateEE() {

	phenotype = "Energy Efficiency"
	newEERaw := dna.PhenotypeAverage(phenotype)
	newEE := Logistic(newEERaw, dna.phenotype.)
	dna.energyEfficiancy = newEE

}

func (bact *Bacteria) BurnEnergy()
