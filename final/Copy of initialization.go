package main

import (
  "math"
  "math/rand"
)

func (b Bacteria) PickInitialLocation(petriRadius float64) (float64, float64) {
  r := rand.Float64() * (petriRadius - b.sizeRadius)
  theta := rand.Float64() * 2 * math.Pi
  x := petriRadius + r*math.Sin(theta)
  y := petriRadius + r*math.Cos(theta)
  return x, y
}

func (pred Predator) PickInitialLocation(petriRadius float64) (float64, float64) {
  r := rand.Float64() * (petriRadius - pred.sizeRadius)
  theta := rand.Float64() * 2 * math.Pi
  x := petriRadius + r*math.Sin(theta)
  y := petriRadius + r*math.Cos(theta)
  return x, y
}

func (bac Bacteria) TooClose(x, a, y, b float64) bool {
  deltaX := x - a
  deltaY := y - b
  dist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
  if dist < 2*bac.sizeRadius {
    return true
  }
  return false
}

func (pred Predator) TooClose(x, a, y, b float64) bool {
  deltaX := x - a
  deltaY := y - b
  dist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
  if dist < 2*pred.sizeRadius {
    return true
  }
  return false
}

//----------------------- Initialize Bacteria -----------------------------

func (dish *Petri) InitializeBact(numBact int) {
  bactSlice := make([]Bacteria,0)
  var a, b float64
  for i := 0; i < numBact; i++ {
    var newBact Bacteria
    newBact.sizeRadius = 3
    x, y := newBact.PickInitialLocation(dish.radius)
    ABlock := rand.Intn(10)
    resistLock := rand.Intn(10)
    potency := rand.Float64() * 10
    for (x == a && y == b) || newBact.TooClose(x, a, y, b) {
      x, y = newBact.PickInitialLocation(dish.radius)
    }
    a = x
    b = y
    newBact.energyEfficiency = 0.3
    newBact.attackRange = 5
    newBact.currentEnergy = 100
    newBact.energyCapacity = 300
    newBact.position.coorX = x
    newBact.position.coorY = y
    newBact.ABenzyme.lock = ABlock
    newBact.ABenzyme.potency = potency
    newBact.resistEnzyme.key = resistLock
    newBact.resistEnzyme.potency = potency
    newBact.stepSize = 5
    newBact.repEnergy = 280
    newBact.strain = i
    newdna := MakeNewDNA()
    newBact.dna = newdna
    bactSlice = append(bactSlice, newBact)
  }
  dish.allBacteria = bactSlice
}
//------------------------------------------------------------------

//----------------------- Initialize food --------------------------

func (p *Petri) InitializeFoodpackage(numFood int, value float64) {
  p.allFoodpack = make([]Foodpackage, numFood)
  for i := 0; i < numFood; i++ {
    p.allFoodpack[i].position.coorX, p.allFoodpack[i].position.coorY = p.allBacteria[0].PickInitialLocation(p.radius)
    p.allFoodpack[i].energy = value
  }
}
//--------------------------------------------------------------------

//---------------------- Initialize predator -------------------------

func (dish *Petri) InitializePredator(n int) {
  predSlice := make([]Predator,0)
  var a, b float64
  for i := 0; i < n; i++ {
    var pred Predator
    pred.sizeRadius = 5
    x, y := pred.PickInitialLocation(dish.radius)
    for (x == a && y == b) || pred.TooClose(x, a, y, b) {
      x, y = pred.PickInitialLocation(dish.radius)
    }
    a = x
    b = y
    pred.energyEfficiency = 0.3
    pred.currentEnergy = 500
    pred.energyCapacity = 2000
    pred.position.coorX = x
    pred.position.coorY = y
    pred.stepSize = 10
    pred.repEnergy = 1700
    predSlice = append(predSlice, pred)
  }
  dish.allPredator = predSlice
}

//------------------------------------------------------------------------

func (p *Petri) InitializeDrugpackage(numDrug int, potency float64) {
  // make a slice of Drugpackage
  p.allDrugpack = make([]Drugpackage, numDrug)
  lock := rand.Intn(10)

  for i := 0; i < numDrug; i++ {
    r := rand.Float64() * (p.radius - p.allBacteria[0].sizeRadius)
    theta := RandomTheta()
    p.allDrugpack[i].position.coorY = p.radius + r * math.Sin(theta)
    p.allDrugpack[i].position.coorX = p.radius + r * math.Cos(theta)
    p.allDrugpack[i].lock = lock
    p.allDrugpack[i].potency = potency
  }
}

func (p *Petri) InitializePredatorKiller(numDrug int, potency float64) {
  // make a slice of Drugpackage
  p.allPredKill = make([]PredatorKiller, numDrug)

  for i := 0; i < numDrug; i++ {
    r := rand.Float64() * (p.radius - p.allPredator[0].sizeRadius)
    theta := RandomTheta()
    p.allPredKill[i].position.coorY = p.radius + r * math.Sin(theta)
    p.allPredKill[i].position.coorX = p.radius + r * math.Cos(theta)
    p.allPredKill[i].potency = potency
  }
}
