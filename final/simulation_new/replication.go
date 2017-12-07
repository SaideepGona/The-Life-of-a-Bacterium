// Kwanho Kim
// 11.05.2017

package main

import (
  "math"
  "math/rand"
  // "fmt"
)

func (b *Bacteria) CanReplicate(p *Petri) bool {
  // This function tests whether a bacterium has enough energy to replicate.
  // If it has energyCap greater than repEnergy, the function returns true.
  if b.currentEnergy > b.repEnergy && b.IsThereSpace(p) {
    return true
  }
  return false
}

func (pred *Predator) CanReplicate(p *Petri) bool {
  // This function tests whether a bacterium has enough energy to replicate.
  // If it has energyCap greater than repEnergy, the function returns true.
  if pred.currentEnergy > pred.repEnergy && pred.IsThereSpace(p) {
    return true
  }
  return false
}

func (b *Bacteria) IsThereSpace(p *Petri) bool {
  if len(p.allBacteria) == 1 {
    return true
  } else {
    count := 0
    for _, bac := range p.allBacteria {
      if bac.position != b.position && b.DistToTarget(bac) < 4*b.sizeRadius {
        count++
      }
      if count > 3 {
        return false
      }
    }
  }
  return true
}

func (pred *Predator) IsThereSpace(p *Petri) bool {
  if len(p.allPredator) == 1 {
    return true
  } else {
    for _, pr := range p.allPredator {
      for pr.position != pred.position {
        if pred.DistToTarget(pr) < 4*pred.sizeRadius {
          return false
        }
        break
      }
    }
  }
  return true
}

func (p *Petri) Replication() {
  numBact := len(p.allBacteria)
  numPred := len(p.allPredator)
  //fmt.Println(numBact)
  // if CanReplicate returns true, bacterium can undergo replication
  for i:=0; i<numBact; i++ {
    b := p.allBacteria[i]
    if b.CanReplicate(p) {
      p.CreateDaughterBac(b)
      //fmt.Println(b.currentEnergy)
      b.currentEnergy = b.currentEnergy/2  // Energy burning for replication process
      b.hasRep = true // mark this bacterium so it doesn't execute any other activities
      //fmt.Println("done rep")
      //fmt.Println(b.currentEnergy)
    }
  }
  for i:=0; i<numPred; i++ {
    pred := p.allPredator[i]
    if pred.CanReplicate(p) {
      p.CreateDaughterPred(pred)
      //fmt.Println(b.currentEnergy)
      pred.currentEnergy = pred.currentEnergy*0.5  // Energy burning for replication process
      pred.hasRep = true // mark this bacterium so it doesn't execute any other activities
      //fmt.Println("done rep")
      //fmt.Println(b.currentEnergy)
    }
  }
}

func (p *Petri) CreateDaughterBac(b Bacteria) {
  //fmt.Println(b.position)
  // Daughter cell is created at an arbitrary location next to parent
  x, y := b.PickLocation(p)
  // create a daughter bacterium at a location x, y
  p.allBacteria = append(p.allBacteria, InitializeBacterium(x, y, b))
}

func (p *Petri) CreateDaughterPred(pred Predator) {
  //fmt.Println(b.position)
  // Daughter cell is created at an arbitrary location next to parent
  x, y := pred.PickLocation(p)
  // create a daughter bacterium at a location x, y
  p.allPredator = append(p.allPredator, InitializePred(x, y, pred))
}

func (b Bacteria) PickLocation(p *Petri) (float64, float64) {
  x, y := b.RepLocation()
  for x > 2 * p.radius || y > 2 * p.radius {
    x, y = b.RepLocation()
  }
  //fmt.Println("x, y = ", x, y)
  return x, y
}

func (pred Predator) PickLocation(p *Petri) (float64, float64) {
  x, y := pred.RepLocation()
  for x > 2 * p.radius || y > 2 * p.radius {
    x, y = pred.RepLocation()
  }
  //fmt.Println("x, y = ", x, y)
  return x, y
}

func (b Bacteria) RepLocation() (float64, float64) {
  distToDaughter := b.sizeRadius*2
  //fmt.Println("dist:", distToDaughter)
  theta := RandomTheta()
  //fmt.Println("theta = ", theta)
  dx := distToDaughter*math.Cos(theta)
  dy := distToDaughter*math.Sin(theta)
  //fmt.Println(dx, dy)
  x := b.position.coorX + dx
  y := b.position.coorY + dy
  return x, y
}

func (pred Predator) RepLocation() (float64, float64) {
  distToDaughter := pred.sizeRadius*2
  //fmt.Println("dist:", distToDaughter)
  theta := RandomTheta()
  //fmt.Println("theta = ", theta)
  dx := distToDaughter*math.Cos(theta)
  dy := distToDaughter*math.Sin(theta)
  //fmt.Println(dx, dy)
  x := pred.position.coorX + dx
  y := pred.position.coorY + dy
  return x, y
}

func RandomTheta() float64 {
  return rand.Float64() * 2 * math.Pi
}

func InitializeBacterium(x, y float64, b Bacteria) Bacteria {
  var newBact Bacteria
  var newdna DNA
  newdna = b.dna
  newBact.sizeRadius = b.sizeRadius
  newBact.attackRange = b.attackRange
  newBact.ABenzyme = b.ABenzyme
  newBact.resistEnzyme = b.resistEnzyme
  newBact.currentEnergy = b.currentEnergy/2
  newBact.energyCapacity = b.energyCapacity
  newBact.repEnergy = b.repEnergy
  newBact.position.coorX = x
  newBact.position.coorY = y
  newBact.stepSize = b.stepSize
  newBact.strain = b.strain
  newBact.dna = newdna
  newBact.hasRep = true
  return newBact
}

func InitializePred(x, y float64, pred Predator) Predator {
  var newPred Predator
  newPred.position.coorX = x
  newPred.position.coorY = y
  newPred.sizeRadius = pred.sizeRadius
  newPred.currentEnergy = pred.currentEnergy * 0.5
  newPred.energyEfficiency = pred.energyEfficiency
  newPred.energyCapacity = pred.energyCapacity
  newPred.stepSize = pred.stepSize
  newPred.hasRep = true
  return newPred
}
