package main

import (
    "fmt"
    "math"
    "math/rand"
)

// REPLICATION AND FIGHT --- BURN ENERGY

type Bacteria struct {
  sizeRadius float64
  position Coords
  ABenzyme ABenzyme
  attackRange float64
  resistEnzyme ResistEnzyme
  currentEnergy float64
  energyCapacity float64
  repEnergy float64
}

type Coords struct {
  coorX float64
	coorY float64
}

type ABenzyme struct {
  lock int
  potency float64
}

type ResistEnzyme struct {
  key int
  potency float64
}

type Petri struct {
  radius float64
  allBacteria []Bacteria

}

func PickInitialLocation(p Petri) (float64, float64) {
  theta := RandomTheta()
  dist := rand.Intn(int(math.Floor(p.radius)))
  x := float64(dist)*math.Cos(theta)
  y := float64(dist)*math.Sin(theta)
  return x, y
}

func TooClose(x, a, y, b float64, bac Bacteria) bool {
  deltaX := x - a
  deltaY := y - b
  dist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
  if dist < 2*bac.sizeRadius {
    return true
  }
  return false
}

func main() {
  bactSlice := make([]Bacteria,0)
  var dish Petri
  dish.radius = 10
  var a, b float64
  for i := 0; i < 1; i++ {
    var newBact Bacteria
    newBact.sizeRadius = 2
    x, y := PickInitialLocation(dish)
    for (x == a && y == b) || TooClose(x, a, y, b, newBact) {
      x, y = PickInitialLocation(dish)
    }
    a = x
    b = y
    newBact.attackRange = 5
    newBact.currentEnergy = 90
    newBact.energyCapacity = 100
    newBact.position.coorX = x
    newBact.position.coorY = y
    newBact.ABenzyme.lock = i+4
    newBact.ABenzyme.potency = float64(i+4)
    newBact.resistEnzyme.key = 7-i
    newBact.resistEnzyme.potency = float64(9-i)
    newBact.repEnergy = 80
    fmt.Println(newBact)
    bactSlice = append(bactSlice, newBact)
  }

  dish.allBacteria = bactSlice

  //fmt.Println(dish)
//  var additionalBact Bacteria
//  dish.allBacteria = append(dish.allBacteria, additionalBact)

/*
  var newDish Petri
  newBactSlice := make([]Bacteria,0)
  newDish.allBacteria = dish.allBacteria[0].Attack(dish, newBactSlice)
  newDish.radius = dish.radius

  for i:=0; i<len(newDish.allBacteria); i++ {
    for j:=0; j<len(dish.allBacteria); j++ {
      if newDish.allBacteria[i].position == dish.allBacteria[j].position {
        dish.allBacteria[j].currentEnergy = newDish.allBacteria[i].currentEnergy
      }
    }
  }
  fmt.Println("result")
  fmt.Println(dish)
  */

  var pointerToDish *Petri
  pointerToDish = &dish
  pointerToDish.Replication()

  fmt.Println("result")
  fmt.Println(dish)
}
