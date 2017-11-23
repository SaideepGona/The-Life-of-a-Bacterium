// Zhenyu Yang 2017/11
// Foodpackage is the only energy source that could be visuallized. When an food package
// is absorbed by a bacteria, energy of the bacteria should have changed.
//This file is to initialize position of foodpackage and to set up the rule for
//detecting bacteria.
package main

import (
   "math"
   "math/rand"
   "fmt"
)

type Coords struct {
  coorX float64
	coorY float64
}

type Bacteria struct {
  position Coords
  theta float64
  currentEnergy float64
  energyCapacity float64
  movement float64
}

type Petri struct {
  radius float64
  allBacteria []Bacteria
  allFoodpack []Foodpackage
}

type Foodpackage struct {
  position Coords
  energy float64
}

func InitializeFoodpackage(p Petri, numFood int,radius float64) Petri{
  //First of all, foodpackage were initially randomized in the petri dish.
  rand.Seed(1000)
  // Rand.Seed between InitializeFoodpackage and InitialBact should be different
  // to prevent they were at the same place at the begining
  pi := math.Pi
  p.radius = radius
  // make a slice of Foodpackage
  p.allFoodpack = make([]Foodpackage, numFood)
  for i := 0; i < numFood; i++ {
    r := rand.Float64() * p.radius
    theta := rand.Float64() * 2 * pi
    // Both r and theta were randomized so that the foodpackage were randomly distrbuted.
    p.allFoodpack[i].position.coorY = p.radius + r*math.Sin(theta)
    p.allFoodpack[i].position.coorX = p.radius + r*math.Cos(theta)
    // Set every food package containing 5 energy scores
    p.allFoodpack[i].energy = 5.0
  }
    return p
}

func InitializeBact(p Petri, numBact int, radius float64) Petri{
  // This process were quite similar to what we have in InitializeFoodpackage
    rand.Seed(20)
    pi := math.Pi
    p.radius = radius
    p.allBacteria = make([]Bacteria, numBact)
    for i := 0; i < numBact ; i++ {
    r := rand.Float64() * p.radius
    theta := rand.Float64() * 2 * pi
    p.allBacteria[i].position.coorY = p.radius + r*math.Sin(theta)
    p.allBacteria[i].position.coorX = p.radius + r*math.Cos(theta)
    p.allBacteria[i].currentEnergy = 60.0
    p.allBacteria[i].energyCapacity = 100.0
    }
    return p
}

func (p Petri) IsRandomMove(i int) bool {
  // The bacteria will randomly move if it is full. If not, it should look for
  //food itsself
   if p.allBacteria[i].currentEnergy < p.allBacteria[i].energyCapacity {
     return false
   }
   return true
}

func MovetoToFood(p Petri, detectRadius float64) {
  // first, range through all the bacteria i a petri dish
  for i := 0; i < len(p.allBacteria); i ++ {
    //fmt.Println(len(p.allBacteria))
    bactMaxEnergy := p.allBacteria[i].energyCapacity
    xBact := p.allBacteria[i].position.coorX
    yBact := p.allBacteria[i].position.coorY
  // k and minDistance stands for which foodpackge were detected and how much
  // distance it has.
    k,minDistance :=  MinDisFood(p.allFoodpack,xBact,yBact,p.radius)
  // To check the distance between foodpackage and bacteria are in the detection range
    if p.IsRandomMove(i) == false && minDistance < detectRadius  {
  // Move the coordinate of bacteria to the nearest foodpackge that has energy and
  // inside its detection range
      p.allBacteria[i].position.coorX = p.allFoodpack[k].position.coorX
      p.allBacteria[i].position.coorY = p.allFoodpack[k].position.coorY
  // If the bacteria need more than the food pakage contained, it just simply take
  // all the energy.
      if p.allBacteria[i].currentEnergy < bactMaxEnergy - p.allFoodpack[k].energy {
      p.allBacteria[i].currentEnergy = p.allBacteria[i].currentEnergy + p.allFoodpack[k].energy
      p.allFoodpack[k].energy = 0.0
      } else {
  // If the bacteria need less energy that food pakage contained, it only takes
  // the amount it want to be full.
        p.allBacteria[i].currentEnergy = bactMaxEnergy
        p.allFoodpack[k].energy = p.allFoodpack[k].energy - (bactMaxEnergy - p.allBacteria[i].currentEnergy)
      }
    }
}

}

func MinDisFood(foodBoard []Foodpackage,xBact,yBact,radius float64) (int,float64){
// The longest distance in a circle is 2 * radius
// return the minimum distance and the number of the foodpackage
// The bacteria only move to the foodpackage that is not 0 energy
  minDistance := 2 * radius
  var j int         // j is used to catch which foodpackge to get
  for k := 0; k < len(foodBoard); k ++ {
      xFood := foodBoard[k].position.coorX
      yFood := foodBoard[k].position.coorY
      distance := math.Sqrt((xBact-xFood)*(xBact-xFood) + (yBact-yFood)*(yBact-yFood))
      if distance < minDistance && foodBoard[k].energy != 0 {
       minDistance  =  distance
       j = k
    }
}
   return j,minDistance
}

func main(){
  var p Petri
  // Call function menetioned before to initialize bacteria and food package
  p = InitializeBact(p,10,100.0)
  p = InitializeFoodpackage(p,100,100.0)
  // Print coordinates before the bacteria's moving
  for j := range p.allBacteria {
    fmt.Println("The position of this bacteria is" , p.allBacteria[j].position)
    fmt.Println("The energy of this bacteria is" , p.allBacteria[j].currentEnergy)
  }
  for i := range p.allFoodpack {
    fmt.Println("The position of this foodpackage is" , p.allFoodpack[i].position)
    fmt.Println("The energy of this foodpackage is" , p.allFoodpack[i].energy)
  }
  // Set number of moves, this example showed 5.
  for i := 0; i < 5; i ++{
    MovetoToFood(p,5.0)
  }
 // Print coordinates after the bacteria's moving
  for i := range p.allBacteria {
    fmt.Println("The position of this bacteria is after moving" , p.allBacteria[i].position)
    fmt.Println("The energy of this bacteria is" , p.allBacteria[i].currentEnergy)
  }
  for i := range p.allFoodpack {
    fmt.Println("The position of this foodpackage is" , p.allFoodpack[i].position)
    fmt.Println("The energy of this foodpackage is" , p.allFoodpack[i].energy)
  }
}
