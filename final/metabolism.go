package main

// Bacteria burns basal metabolic energy every turn.
func (p *Petri) CostBasicEnergy(cost float64) {
  for i := 0; i < len(p.allBacteria) ; i ++ {
    p.allBacteria[i].currentEnergy -= cost
  }
}

/*For every iteration, every bacteria costs some basic energy even it doesn't move*/
func (p *Petri) CostBasicEnergyPredator() {
    for i := 0; i < len(p.allPredator) ; i ++ {
      p.allPredator[i].currentEnergy -= 20.0
    }
}

func (p *Petri) IsAlive(index int) bool {
  if p.allBacteria[index].currentEnergy <= 0.0 {
    return false
  }
  return true
}

/*check if the predator is live. If it is, return true. If not, return false*/
func (p *Petri) IsLivePredator(i int) bool {
  if p.allPredator[i].currentEnergy <= 0.0 {
    return false
  }
  return true
}

// If bacteria are dead (out of energy), this function removes them.
func (p *Petri) ChecktoDeleteBact(){
  for index := 0; index < len(p.allBacteria); index ++ {
    if p.IsAlive(index) == false {
       p.allBacteria = append(p.allBacteria[:index], p.allBacteria[index+1:]...)
    }
  }
}

// This function removes used up food package on Petri dish.
func (p *Petri) ChecktoDeleteFood(){
  for index := 0; index < len(p.allFoodpack); index ++ {
    if p.allFoodpack[index].energy <= 0.0 {
      p.allFoodpack = append(p.allFoodpack[:index], p.allFoodpack[index+1:]...)
    }
  }
}

/*check to see if predator is live by calling isLivePredator*/
func (p *Petri) ChecktoDeletePred(){
    for index := 0; index < len(p.allPredator); index ++ {
      if p.IsLivePredator(index) == false {
         p.allPredator = append(p.allPredator[:index], p.allPredator[index+1:]...)
      }
    }
}

func (bact *Bacteria) UpdateEE() {

	// Updates energy efficiency
	newEERaw := bact.dna.PhenotypeAverage("EE")
	newEE := Logistic(newEERaw, bact.dna.phenotypes["EE"].aggFuncArgs)
	bact.energyEfficiency = newEE

}

func (p *Petri) UpdateAllEE() {
	//fmt.Println("UpdateAllEE")
	// Updates the energy efficiency for all bacteria in petri dish
	for i := 0; i < len(p.allBacteria); i++{
		p.allBacteria[i].UpdateEE()
	}
}
