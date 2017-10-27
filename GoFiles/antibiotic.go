// Kwanho Kim
// 10.21.2017

package main

import (

)


func (b *Bacteria) Attack() {
  // scan bacteria in attack range and list them under targets
  var targets []*Bacteria
  targets = b.OthersInRange(b.location.Petri.allBacteria)
  for target := range targets {
    if b.ABenzyme.lock != target.ResistEnzyme.key {
      b.InflictDamage(target, b.ABenzyme.potency)
    } else if b.ABenzyme.lock == target.ResistEnzyme.key {
      if b.ABenzyme.potency > target.ResistEnzyme.potency {
        attackDamage = b.ABenzyme.potency - target.ResistEnzyme.potency
        b.InflictDamage(target, attackDamage)
      }
    }
  }
}

// for a given bacterium, SenseOther function determines wheter
// there are other bacteria near by to attack
func (b *Bacteria) OthersInRange(all []*Bacteria) []*Bacteria {
  var inRange []*Bacteria
  for bacterium := range all {
    attackRange = b.AttackRange + bacterium.size.radius
    if b.DistToTarget(bacterium) <= attackRange {
      inRange = append(inRange, bacterium)
    }
  }
  return inRange
}

func (b *Bacteria) DistToTarget(target *Bacteria) float64 {
  deltaX := b.location.coorX - target.location.coorX
  deltaY := b.location.coorY - target.location.coorY
  dist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
  return dist
}

func (b *Bacteria) InflictDamage(t *bacteria, damage float64) {

}
