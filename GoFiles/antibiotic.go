// Kwanho Kim
// 10.21.2017

package main

import (

)

//--------------------------------------------------------------------------
// This is a model bacteria to test antibiotic.go
type Bacteria struct {
  currentEnergy float64
  energyCapacity float64
  ABenzyme ABenzyme
  resistEnzyme ResistEnzyme
}

type ABenzyme struct {
  lock int
  potency int
}

type ResistEnzyme struct {
  key int
  potency int
}

func Test() {
  
}
//--------------------------------------------------------------------------

func (b *Bacteria) Attack(p Petri) {
  // scan bacteria in attack range and list them under targets
  var targets []*Bacteria
  targets = b.OthersInRange(Petri.allBacteria)
  for target := range targets {
    if b.ABenzyme.lock != target.ResistEnzyme.key {
      b.InflictDamage(target, b.ABenzyme.potency)
    } else if b.ABenzyme.lock == target.ResistEnzyme.key && b.ABenzyme.potency > target.ResistEnzyme.potency {
      attackDamage = b.ABenzyme.potency - target.ResistEnzyme.potency
      b.InflictDamage(target, attackDamage)
    }
  }
}

// for a given bacterium, SenseOther function determines wheter
// there are other bacteria near by to attack
func (b *Bacteria) OthersInRange(all []Bacteria) []*Bacteria {
  var inRange []*Bacteria
  for bacterium := range all {
    range = b.attackRange + bacterium.size.radius
    if b.DistToTarget(bacterium) <= range {
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

// Damage can range from 0 to 9
func (b *Bacteria) InflictDamage(t *Bacteria, damage float64) {
  t.energyCap -= damage*10
  if t.energyCap < 0 {
    // Die!
  }
}
