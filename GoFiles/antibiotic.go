// Kwanho Kim
// 10.21.2017

package main

import (
 "math"
 "fmt"
)

// REMEMBER TO BURN ENERGY

//--------------------------------------------------------------------------

func (b *Bacteria) Attack(p Petri, new []Bacteria) []Bacteria {
  // scan bacteria in attack range and list them under targets
  var targets []Bacteria
  targets = b.OthersInRange(p.allBacteria)
  for _, target := range targets {
    if b.ABenzyme.lock != target.resistEnzyme.key {
      target.currentEnergy = b.InflictDamage(target, b.ABenzyme.potency)
    } else if b.ABenzyme.lock == target.resistEnzyme.key && b.ABenzyme.potency > target.resistEnzyme.potency {
      attackDamage := b.ABenzyme.potency - target.resistEnzyme.potency
      target.currentEnergy = b.InflictDamage(target, attackDamage)
    }
    new = append(new, target)
  }
  fmt.Println("new")
  fmt.Println(new)
  return new
}

// for a given bacterium, SenseOther function determines wheter
// there are other bacteria near by to attack
func (b *Bacteria) OthersInRange(all []Bacteria) []Bacteria {
  var inRange []Bacteria
  for _, bacterium := range all {
    r := b.attackRange // + bacterium.size.radius
    self := b.position
    if b.DistToTarget(bacterium) <= r && bacterium.position != self {
      inRange = append(inRange, bacterium)
    }
  }
  fmt.Println("targets")
  fmt.Println(inRange)
  return inRange
}

func (b *Bacteria) DistToTarget(target Bacteria) float64 {
  deltaX := b.position.coorX - target.position.coorX
  deltaY := b.position.coorY - target.position.coorY
  dist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
  return dist
}

// Damage can range from 0 to 9
func (b *Bacteria) InflictDamage(t Bacteria, damage float64) float64 {
  inflictedDamage := damage*10
  t.currentEnergy -= inflictedDamage
  fmt.Println("damage")
  fmt.Println(inflictedDamage)
  fmt.Println("energy left")
  fmt.Println(t.currentEnergy)
  if t.currentEnergy < 0 {
    // Die!
  }
  return t.currentEnergy
}
