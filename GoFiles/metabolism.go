 // Saideep Gona

 
 func (p *Petri) BasalMetabolizeAndDeath() {

	// For all bacteria burns basal metabolism and checks if bacteria are dead (out of energy). Removes them if so.

	liveList := make([]int, 0)

	for index, bacteria := range p.allBacteria {
		bacteria.BasalBurn(p.basalMetabolism)
		if bacteria.currentEnergy > 0 {
			liveList.append(bacteria)
		}
	}

	p.allBacteria = liveList

 }

 func (b *Bacteria) BasalBurn(burn float64) {

	// Burns input amount of energy from current energy reserves

	b.currentEnergy = b.currentEnergy-burn

 }
