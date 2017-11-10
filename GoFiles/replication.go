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
  if b.energyCap > 80 {
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

// How big is our Petri dish?????
func (b *bacteria) CreateDaughterBac() {
  // Daughter cell is created at an arbitrary location next to parent
  distToDaughter := 5
  theta := RandomTheta()
  x := b.location.coorX + distToDaughter*math.Cos(theta)
  y := b.location.coorY + distToDaughter*math.Sin(theta)
}

func RandomTheta() float64 {
  rand.Seed(time.Now().UnixNano())
  return rand.Float64() * 2 * math.Pi
}

// Does this function exist already?
func InitializeBacterium() {

}
