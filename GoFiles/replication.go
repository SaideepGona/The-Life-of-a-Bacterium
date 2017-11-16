// Kwanho Kim
// 11.05.2017

package main

import (
  "math"
  "math/rand"
  "time"
)

func (b *Bacteria) CanReplicate(p Petri) bool {
  // This function tests whether a bacterium has enough energy to replicate.
  // If it has energyCap greater than certain point, the function returns true.
  if b.currentEnergy > b.repEnergy && b.IsThereSpace(p) {
    return true
  }
  return false
}

func (b *Bacteria) IsThereSpace(p Petri) bool {
  for _, bac := range p.allBacteria {
    if b.DistToTarget(bac) < 2*b.sizeRadius {
      return true
    }
  }
  return false
}

func (b *Bacteria) Replication(p Petri) {
  // if CanReplicate returns true, bacterium can undergo replication
  if b.CanReplicate(p) {
    b.CreateDaughterBac(p)
  }
  
  // b.BurnEnergy(Replication, 1)
}

func (b *Bacteria) CreateDaughterBac(p Petri) {
  // Daughter cell is created at an arbitrary location next to parent
  distToDaughter := b.sizeRadius*2
  theta := RandomTheta()
  x := b.position.coorX + distToDaughter*math.Cos(theta)
  y := b.position.coorY + distToDaughter*math.Sin(theta)
  // create a daughter bacterium at a location x, y
  p.allBacteria = append(p.allBacteria, b.InitializeBacterium(x, y))
}

func RandomTheta() float64 {
  rand.Seed(time.Now().UnixNano())
  return rand.Float64() * 2 * math.Pi
}

// Does this function exist already?
func (b *Bacteria) InitializeBacterium(x, y float64) Bacteria {
  var newBact Bacteria
  newBact.attackRange = b.attackRange
  newBact.ABenzyme = b.ABenzyme
  newBact.resistEnzyme = b.resistEnzyme
  newBact.currentEnergy = 50
  newBact.energyCapacity = b.energyCapacity
  newBact.repEnergy = b.repEnergy
  newBact.position.coorX = x
  newBact.position.coorY = y
  return newBact
}
