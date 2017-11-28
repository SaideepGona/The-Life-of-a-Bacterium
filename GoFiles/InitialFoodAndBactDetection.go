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
//   "time"
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
  stepSize float64
  energyEfficiency float64
  sizeRadius float64
  repEnergy float64
  attackRange float64
  ABenzyme ABenzyme
  resistEnzyme ResistEnzyme
  dna DNA
}

type ResistEnzyme struct {
  key int
  potency float64
}


type ABenzyme struct {
  lock int
  potency float64
}

type Petri struct {
  radius float64
  allBacteria []Bacteria
  allFoodpack []Foodpackage
  basalMetabolism float64
}

type Foodpackage struct {
  position Coords
  energy float64

}

func (p *Petri) InitializeFoodpackage(numFood int,radius float64) {
  //First of all, foodpackage were initially randomized in the petri dish.
  //rand.Seed(1000)
  // Rand.Seed between InitializeFoodpackage and InitialBact should be different
  // to prevent they were at the same place at the begining
  pi := math.Pi
  // make a slice of Foodpackage
  p.allFoodpack = make([]Foodpackage, numFood)
  for i := 0; i < numFood; i++ {
    r := rand.Float64() * (p.radius - p.allBacteria[0].sizeRadius)
    theta := rand.Float64() * 2 * pi
    // Both r and theta were randomized so that the foodpackage were randomly distrbuted.
    p.allFoodpack[i].position.coorY = p.radius + r * math.Sin(theta)
    p.allFoodpack[i].position.coorX = p.radius + r * math.Cos(theta)
    // Set every food package containing 5 energy scores
    p.allFoodpack[i].energy = 8.0
}
}

func (p *Petri) InitializeEnergyEfficiency() {
  for i := range p.allBacteria {
    p.allBacteria[i].energyEfficiency = 0.6
  }
}

func (p *Petri) InitializeBact(numBact int, radius float64) {
  // This process were quite similar to what we have in InitializeFoodpackage

    pi := math.Pi
    p.radius = radius
    p.allBacteria = make([]Bacteria, numBact)
    newdna := MakeNewDNA()
    fmt.Println(newdna)
    for i := 0; i < numBact ; i++ {
      r := rand.Float64()* p.radius
      theta := rand.Float64() * 2 * pi
      p.allBacteria[i].position.coorY = p.radius - p.allBacteria[i].sizeRadius + r*math.Sin(theta)
      p.allBacteria[i].position.coorX = p.radius - p.allBacteria[i].sizeRadius + r*math.Cos(theta)
      p.allBacteria[i].currentEnergy = 200.0
      p.allBacteria[i].energyCapacity = 300.0
      p.allBacteria[i].stepSize = 5.0
      p.allBacteria[i].sizeRadius = 3.0
      p.allBacteria[i].repEnergy = 100.0
      p.allBacteria[i].dna = newdna
    }
}


func (p *Petri) IsRandomMove(i int) bool {
  // The bacteria will randomly move if it is full. If not, it should look for
  //food itsself
   if p.allBacteria[i].currentEnergy < p.allBacteria[i].energyCapacity {
     return false
   }
   return true
}

func (p *Petri) IsLive(index int) bool {
  if p.allBacteria[index].currentEnergy <= 0.0 {
    return false
  }
    return true
}

func (p *Petri) IsEnd() bool{
  count := 0
  for i := range p.allBacteria {
    if p.IsLive(i) == false {
      count ++
    }
    }
    if count == len(p.allBacteria) {
      return true
      }
      return false
}

func (p *Petri)CostBasicEnergy() {
    for i := 0; i < len(p.allBacteria) ; i ++ {
      p.allBacteria[i].currentEnergy -= 3.0
    }
  }

func (p *Petri) ChecktoDeleteBact(){
  for index := 0; index < len(p.allBacteria); index ++ {
    if p.IsLive(index) == false {
       p.allBacteria = append(p.allBacteria[:index], p.allBacteria[index+1:]...)
    }
  }
}

func (p *Petri) ChecktoDeleteFood(){
  for index := 0; index < len(p.allFoodpack); index ++ {
  if p.allFoodpack[index].energy <= 0.0 {
     p.allFoodpack = append(p.allFoodpack[:index], p.allFoodpack[index+1:]...)
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

func (p *Petri) MoveToFood() {
  // first, range through all the bacteria i a petri dish
  for i := 0; i < len(p.allBacteria); i ++ {
    if p.IsFoodAround(i) == true {
    //fmt.Println(len(p.allBacteria))
    bactMaxEnergy := p.allBacteria[i].energyCapacity
    xBact := p.allBacteria[i].position.coorX
    yBact := p.allBacteria[i].position.coorY
  // k and minDistance stands for which foodpackge were detected and how much
  // distance it has.
    k,minDistance :=  MinDisFood(p.allFoodpack,xBact,yBact,p.radius)
    energyConsumption := minDistance*0.6*p.allBacteria[i].energyEfficiency
    p.allBacteria[i].currentEnergy -= energyConsumption
  // To check the distance between foodpackage and bacteria are in the detection range
    if p.IsLive(i) == true && p.IsRandomMove(i) == false && minDistance < p.allBacteria[i].stepSize {
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
    }/*else if p.IsLive(i) == true && p.IsRandomMove(i) == true {
      pi := math.Pi
        theta := rand.Float64() * 2 * pi
        p.allBacteria[i].position.coorX += p.allBacteria[i].stepSize * math.Cos(theta)
        p.allBacteria[i].position.coorY += p.allBacteria[i].stepSize * math.Sin(theta)
        p.allBacteria[i].currentEnergy -= p.allBacteria[i].stepSize*1*p.allBacteria[i].energyEfficiency
    }
}*/
   }
  }
}
