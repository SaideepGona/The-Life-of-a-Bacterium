// Kwanho Kim
// 10.21.2017

package main

import (

)


func (b *Bacteria) Attack() {
  // scan bacteria in attack range and list them under targets
  var targets []*Bacteria
  targets = b.OthersInRange()
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
func (b *Bacteria) OthersInRange() []*Bacteria {

}

func (b *Bacteria) InflictDamage(t *bacteria, damage float64) {

}
