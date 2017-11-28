// Kwanho Kim
// 10.21.2017

package main

import (
 "math"
// "fmt"
)

// REMEMBER TO BURN ENERGY

//--------------------------------------------------------------------------

func (p *Petri) Attack() {
  for i:=0; i<len(p.allBacteria); i++ {
  // scan bacteria in attack range and list them under targets
    b := p.allBacteria[i]
    pointerToB := &b
    var targets []Bacteria
    for j:=0; j<len(p.allBacteria); j++ {
      target := p.allBacteria[j]
      if i != j && pointerToB.IsTarget(target) {
        targets = append(targets, target)
      }
    }
//    fmt.Println("initial targets")
//    fmt.Println(targets)

    targets = UpdateTarget(b, targets)
//    fmt.Println("after update")
//    fmt.Println(targets)

    for k:=0; k<len(p.allBacteria); k++ {
      for x:=0; x<len(targets); x++ {
        if p.allBacteria[k].position == targets[x].position {
          p.allBacteria[k] = targets[x]
        }
      }
    }

  }
}

func UpdateTarget(b Bacteria, targets []Bacteria) []Bacteria {
  for index:=0; index<len(targets); index++ {
    if b.ABenzyme.lock != targets[index].resistEnzyme.key {
      targets[index] = InflictDamage(targets[index], b.ABenzyme.potency)
    } else if b.ABenzyme.lock == targets[index].resistEnzyme.key && b.ABenzyme.potency > targets[index].resistEnzyme.potency {
        attackDamage := b.ABenzyme.potency - targets[index].resistEnzyme.potency
        targets[index] = InflictDamage(targets[index], attackDamage)
    }
  }
  return targets
}

/*
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
*/
// for a given bacterium, SenseOther function determines wheter
// there are other bacteria near by to attack

func (b *Bacteria) IsTarget(t Bacteria) bool {
  r := b.attackRange
  if b.DistToTarget(t) <= r {
    return true
  }
  return false
}
/*
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
*/
func (b *Bacteria) DistToTarget(target Bacteria) float64 {
  deltaX := b.position.coorX - target.position.coorX
  deltaY := b.position.coorY - target.position.coorY
  dist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
  return dist
}

// Damage can range from 0 to 9
func InflictDamage(t Bacteria, damage float64) Bacteria {
  inflictedDamage := damage*10
  t.currentEnergy -= inflictedDamage
  // if currentEnergy <= 0, its dead
  return t
}
