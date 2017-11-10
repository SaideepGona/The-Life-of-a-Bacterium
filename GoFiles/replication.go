// Kwanho Kim
// 11.05.2017

package main

import (
  "math"
  "math/rand"
  "time"
)

func (b *bacteria) CanReplicate() bool {
  // This function tests whether a bacterium has enough energy to replicate.
  // If it has energyCap greater than certain point, the function returns true.
  if b.energyCap > b.repEnergy {
    return true
  }
  return false
}

func (b *bacteria) Replication() {
  // if CanReplicate returns true, bacterium can undergo replication
  if b.CanReplicate() == true {
    b.CreateDaughterBac()
  }
}

func (b *bacteria) CreateDaughterBac(p Petri) {
  // Daughter cell is created at an arbitrary location next to parent
  distToDaughter := b.sizeRadius*2
  theta := RandomTheta()
  x := b.position.coorX + distToDaughter*math.Cos(theta)
  y := b.position.coorY + distToDaughter*math.Sin(theta)
  // create a daughter bacterium at a location x, y
  InitializeBacterium(x, y)
}

func RandomTheta() float64 {
  rand.Seed(time.Now().UnixNano())
  return rand.Float64() * 2 * math.Pi
}

// Does this function exist already?
func InitializeBacterium(x, y float64) {

}
