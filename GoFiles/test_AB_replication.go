package main

import (
    "fmt"
    "math/rand"
    "time"
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

func main() {

  bactSlice := make([]Bacteria,0)
  for i := 0; i < 3; i++ {
    rand.Seed(time.Now().UnixNano())
    n := rand.Float64()
    fmt.Println("n")
    fmt.Println(n)
    var newBact Bacteria
    newBact.attackRange = 5
    newBact.currentEnergy = 90
    newBact.energyCapacity = 100
    newBact.position.coorX = n*10
    newBact.position.coorY = n*10
    newBact.ABenzyme.lock = i+4
    newBact.ABenzyme.potency = float64(i+4)
    newBact.resistEnzyme.key = 7-i
    newBact.resistEnzyme.potency = float64(9-i)
    newBact.repEnergy = 80
    fmt.Println(newBact)
    bactSlice = append(bactSlice, newBact)
  }

  var dish Petri
  dish.allBacteria = bactSlice
  dish.radius = 10

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

  for i:=0; i<len(dish.allBacteria); i++ {
    dish.allBacteria[i].Replication(dish)
  }
  fmt.Println("result")
  fmt.Println(dish)
}
