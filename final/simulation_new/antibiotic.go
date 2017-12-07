// Kwanho Kim
// 10.21.2017

package main

import (
 "math"
// "fmt"
)

// REMEMBER TO BURN ENERGY

// Attack is a Petri dish method that mimics fighting among bacteria.
// To simulate protein-protein interaction, lock and key model is used.
// The lock and key values are intrinsic to each bacterium and determined by its DNA.
func (p *Petri) Attack() {
  for i:=0; i<len(p.allBacteria); i++ {
    if p.allBacteria[i].hasRep == true {
      continue
    } else { // if the bacterium hasn't replicated in this round
      b := p.allBacteria[i]
      pointerToB := &b

      var targets []Bacteria

      // For each bacterium in a Petri dish, it scans other bacteria in its
      // attack range and appends them to a list called targets.
      for j:=0; j<len(p.allBacteria); j++ {
        target := p.allBacteria[j]
        if i != j && pointerToB.IsTarget(target) && target.strain != b.strain {
          targets = append(targets, target)
        }
      }

      // When targets list is not empty, this bacterium will attack all target Bacteria
      // in its range. Regardless of how many it is attacking, cost of energy is fixed.
      if len(targets) != 0 {
        b.currentEnergy -= 10 // SET AMOUNT OF ENERGY TO BE CONSUMED HERE
      }

      //fmt.Println("initial targets")
      //fmt.Println(targets)
      targets = UpdateTarget(b, targets)
      //fmt.Println("after update")
      //fmt.Println(targets)

      for k:=0; k<len(p.allBacteria); k++ {
        for x:=0; x<len(targets); x++ {
          if p.allBacteria[k].position == targets[x].position {
            p.allBacteria[k] = targets[x]
          }
        }
      }
    } // end else
  } // end for
}


// UpdateTarget function takes in an attacking bacterium (b) and its target list,
// and compares each target with b to check whether lock and key matches.
func UpdateTarget(b Bacteria, targets []Bacteria) []Bacteria {
  for index := 0; index < len(targets); index++ {
    if b.ABenzyme.lock != targets[index].resistEnzyme.key { // lock and key doesn't match
      targets[index] = InflictDamageE(targets[index], b.ABenzyme.potency) // antibiotic enzyme can incur damage to its full potency
    } else if b.ABenzyme.lock == targets[index].resistEnzyme.key && b.ABenzyme.potency > targets[index].resistEnzyme.potency { // if lock and key matches
        attackDamage := b.ABenzyme.potency - targets[index].resistEnzyme.potency // attack damage compansates resistEnzyme's potency
        targets[index] = InflictDamageE(targets[index], attackDamage)
    }
  }
  return targets
}


// Simple function that compares the distance to its target bacteria and its
// attack range to decide whether a target bacteria is in range or not.
func (b *Bacteria) IsTarget(t Bacteria) bool {
  dist := b.DistToTarget(t) - t.sizeRadius
  if dist <= b.attackRange {
    return true
  }
  return false
}

// Calculates the distance between the two bacteria
func (b *Bacteria) DistToTarget(target Bacteria) float64 {
  deltaX := b.position.coorX - target.position.coorX
  deltaY := b.position.coorY - target.position.coorY
  dist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
  return dist
}

func (pred *Predator) DistToTarget(target Predator) float64 {
  deltaX := pred.position.coorX - target.position.coorX
  deltaY := pred.position.coorY - target.position.coorY
  dist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
  return dist
}

// Damage is taken in ther target bacteria as loss of energy
func InflictDamageE(t Bacteria, damage float64) Bacteria {
  inflictedDamage := damage*10
  //fmt.Println("damage")
  //fmt.Println(inflictedDamage)
  t.currentEnergy -= inflictedDamage
//  t.sizeRadius -= 0.05
  // if currentEnergy <= 0, its dead
  return t
}

func InflictDamageES(t Bacteria, damage float64) Bacteria {
  inflictedDamage := damage*10
  //fmt.Println("damage")
  //fmt.Println(inflictedDamage)
  t.currentEnergy -= inflictedDamage
  t.sizeRadius -= 0.1
  return t
}


// The following functions pull value from DNA and updates potency of two enzymes
func (bact *Bacteria) UpdateAE() {

	// Updates energy efficiency
	newAERaw := bact.dna.PhenotypeAverage("AE")
	newAE := Logistic(newAERaw, bact.dna.phenotypes["AE"].aggFuncArgs)
	bact.ABenzyme.potency = newAE

}

func (p *Petri) UpdateAllAE() {
	//fmt.Println("UpdateAllAE")
	// Updates the energy efficiency for all bacteria in petri dish
	for i := 0; i < len(p.allBacteria); i++{
		p.allBacteria[i].UpdateAE()
	}
}

func (bact *Bacteria) UpdateRE() {
	// Updates energy efficiency
	newRERaw := bact.dna.PhenotypeAverage("RE")
	newRE := Logistic(newRERaw, bact.dna.phenotypes["RE"].aggFuncArgs)
	bact.resistEnzyme.potency = newRE

}

func (p *Petri) UpdateAllRE() {
	//fmt.Println("UpdateAllEE")
	// Updates the energy efficiency for all bacteria in petri dish
	for i := 0; i < len(p.allBacteria); i++{
		p.allBacteria[i].UpdateRE()
	}
}
