//Zhenyu Yang 2017/10/27

package main

import (
//  "fmt"
  "image"
  "image/color"
  "gonum.org/v1/plot/plotter"
  "math"
)

func (p *Petri) DrawPetri() Canvas {
  width := p.radius
  white:= MakeColor(255,255,255)
  black := MakeColor(0, 0, 0)
  pic := CreateNewCanvas(int(width * 2.0),int(width * 2.0))
  pic.SetFillColor(black)
  pic.Clear()
  pic.SetFillColor(white)
  pic.Circle(width,width,width)
  pic.Fill()

  // This boolean returns true when there are more than one strain of bacteria
  // exist in the Petri dish.

  var competition bool
  for k:=1; k<len(p.allBacteria); k++ {
    if p.allBacteria[0].strain != p.allBacteria[k].strain {
      competition = true
    }
  }
  //fmt.Println(competition)

  for j := range p.allFoodpack {
    if p.allFoodpack[j].energy > 0.0 {
      pic.DrawFoodpack(p.allFoodpack[j].position.coorX,p.allFoodpack[j].position.coorY,1.0)
    }
  }

  for index := range p.allPredator {
    if p.IsLivePredator(index) == true {
      pic.DrawLivePredator(p,index,p.allPredator[index].position.coorX , p.allPredator[index].position.coorY,p.allPredator[index].sizeRadius)
    }
  }

  for k:=range p.allDrugpack {
    pic.DrawDrugpack(p.allDrugpack[k].position.coorX, p.allDrugpack[k].position.coorY, 1.0)
  }

  for k:=range p.allPredKill {
    pic.DrawPredatorKiller(p.allPredKill[k].position.coorX, p.allPredKill[k].position.coorY, 1.0)
  }

/*
  for i:=0; i<len(p.allBacteria); i++ {
    if competition == true { // p.IsAlive(i) == true && competition == true {
      pic.DrawBacteriaMultiColor(p.allBacteria[i].strain, p.allBacteria[i].position.coorX , p.allBacteria[i].position.coorY,p.allBacteria[i].sizeRadius)
    } else if competition == false { // p.IsAlive(i) == true && competition == false {
      pic.DrawBacteriaEnergy(p,i,p.allBacteria[i].position.coorX , p.allBacteria[i].position.coorY,p.allBacteria[i].sizeRadius)
    }
  }
*/

  // Drawing bacteria without switching over to energy showing mode
  for i:=0; i<len(p.allBacteria); i++ {
    pic.DrawBacteriaMultiColor(p.allBacteria[i].strain, p.allBacteria[i].position.coorX , p.allBacteria[i].position.coorY,p.allBacteria[i].sizeRadius)
    competition = competition
  }

  return pic
}

func (pic *Canvas) DrawFoodpack(x,y,r float64) {
    // Set Agar to be small yellow dot with a circle of radius 1
      foodColor:= MakeColor(240,200,200)
      pic.SetFillColor(foodColor)
      pic.Circle(x,y,r)
      pic.Fill()
}

func (pic *Canvas) DrawBacteriaMultiColor(strain int,x,y,r float64) {
  var color color.Color

  // Color scheme for each strain of bacteria
  switch {
  case strain == 0:
    color = MakeColor(255, 0, 0)
  case strain == 1:
    color = MakeColor(0, 0, 255)
  case strain == 2:
    color = MakeColor(0, 255, 0)
  case strain == 3:
    color = MakeColor(255, 100, 255)
  case strain == 4:
    color = MakeColor(255, 255, 100)
  }

  pic.SetFillColor(color)
  pic.Circle(x,y,r)
  pic.Fill()
}

func (pic *Canvas) DrawBacteriaEnergy(p *Petri, i int, x,y,r float64) {
  n := p.allBacteria[i].currentEnergy
  maxE := p.allBacteria[i].energyCapacity
  strain := p.allBacteria[i].strain
  color := MakeColor(0,0,0)

  // Color scheme with energy consideration
  switch {
  case strain == 0:
    color = MakeColor(uint8(255 - (255-n*(255/maxE))), 0, 0)
  case strain == 1:
    color = MakeColor(0, 0, uint8(255 - (255-n*(255/maxE))))
  case strain == 2:
    color = MakeColor(0, uint8(255 - (255-n*(255/maxE))), 0)
  case strain == 3:
    color = MakeColor(uint8(255 - (255-n*(255/maxE))), uint8(100 - (100-n*(100/maxE))), uint8(255 - (255-n*(255/maxE))))
  case strain == 4:
    color = MakeColor(uint8(255 - (255-n*(255/maxE))), uint8(255 - (255-n*(255/maxE))), uint8(100 - (100-n*(100/maxE))))
  }

  pic.SetFillColor(color)
  pic.Circle(x,y,r)
  pic.Fill()
}

/*draw the live predator. The color of predator depends on its energy. Orinally, the color of the predator is blue.
If it has less and less energy, it turns black*/
func (pic *Canvas) DrawLivePredator(p *Petri,i int,x,y,r float64) {
   n := p.allPredator[i].currentEnergy
   maxE := p.allPredator[i].energyCapacity
   /*if the current energy of predator is greater than 255, set the color to blue*/
   color := MakeColor(uint8(128 - (128-n*(128/maxE))), uint8(128 - (128-n*(128/maxE))), uint8(128 - (128-n*(128/maxE))))
   pic.SetFillColor(color)
   pic.Circle(x,y,r)
   pic.Fill()
}

func (pic *Canvas) DrawDrugpack(x,y,r float64) {
    // Set Agar to be small yellow dot with a circle of radius 1
      color := MakeColor(100,150,200)
      pic.SetFillColor(color)
      pic.Circle(x,y,r)
      pic.Fill()
}

func (pic *Canvas) DrawPredatorKiller(x,y,r float64) {
    // Set Agar to be small yellow dot with a circle of radius 1
      color := MakeColor(200,100,200)
      pic.SetFillColor(color)
      pic.Circle(x,y,r)
      pic.Fill()
}

func (p *Petri) IsFoodAround(index int) bool{
  for i := 0; i < len(p.allFoodpack); i ++{
    xf := p.allFoodpack[i].position.coorX
    yf := p.allFoodpack[i].position.coorY
    xb := p.allBacteria[index].position.coorX
    yb := p.allBacteria[index].position.coorY
    if math.Sqrt((yb-yf)*(yb-yf)+(xb-xf)*(xb-xf)) < p.allBacteria[index].stepSize{
      return true
    }
  }
  return false
}

// Simulation ends when all bacteria are dead
func (p *Petri) IsEnd() bool{
  count := 0
  for i := range p.allBacteria {
    if p.IsAlive(i) == false {
      count ++
    }
  }
  if count == len(p.allBacteria) {
    return true
  }
  return false
}

func (p *Petri) AnimationPetri(numStrain, numIteration, drugIntro, predatorIntro, numPred int, basalMetabolicRate float64) []image.Image {
  gifImages := make([]image.Image,0)
   count := 0
    pic := p.DrawPetri()
    pic.SaveToPNG("data/Original.png")


//------------------- Initiate data collection----------------------
    // We will keep track of the following five data:
    // 1. Energy Efficiency
    EEs := make(plotter.XYs, numIteration)

    // 2. Movement
    movements := make(plotter.XYs, numIteration)

    // 3. Antibiotic Enzyme
    AEs := make(plotter.XYs, numIteration)

    // 4. Resistance Enzyme
    REs := make(plotter.XYs, numIteration)

    // 5. Number of cells per strain of bacteria
    cellCount := make([]plotter.XYs, numStrain+1)
    for i:=0; i<numStrain+1; i++ {
      cellCount[i] = make(plotter.XYs, numIteration)
    }
//-------------------------------------------------------------------

// *** THE MAIN SIMULATION LOOP STARTS HERE ***
    for p.IsEnd() == false && count < numIteration {
      for i:=0; i<len(p.allBacteria); i++ {
        p.allBacteria[i].hasRep = false
      }

      for i:=0; i<len(p.allPredator); i++ {
        p.allPredator[i].hasRep = false
      }

      // Initiate antibiotic drug when there are more than certain number of cells
      if len(p.allBacteria) >= drugIntro && len(p.allDrugpack) == 0 {
        p.InitializeDrugpackage(500, 50)
      }

      // Initiate predators when antibiotic drug doesn't work
      if len(p.allBacteria) >= predatorIntro && len(p.allPredator) == 0 {
        p.InitializePredator(numPred)
      }

      // To prevent predators to dominate simulation, introduce predator killer drug
      // into the Petri dish.
      if len(p.allBacteria) < 10 && len(p.allPredator) != 0 && len(p.allPredKill) == 0 {
        p.InitializePredatorKiller(500, 1000)
        var deleteDrug []Drugpackage
        p.allDrugpack = deleteDrug
      }

      if len(p.allPredator) == 0 && len(p.allPredKill) != 0 {
        var deletePredKill []PredatorKiller
        p.allPredKill = deletePredKill
      }

/*
      // Replenish food
      if len(p.allFoodpack) < 1000 {
        for i:=0; i<1000; i++ {
          var moreFood Foodpackage
          moreFood.position.coorX, moreFood.position.coorY = p.allBacteria[0].PickInitialLocation(p.radius)
          moreFood.energy = 15
          p.allFoodpack = append(p.allFoodpack, moreFood)
        }
      }
*/

//------- Keep track of data change in each iteration --------
      EEs[count].Y = p.AllPhenotypeExpectation("EE")
      EEs[count].X = float64(count)

      movements[count].Y = p.AllPhenotypeExpectation("Movement")
      movements[count].X = float64(count)

      AEs[count].Y = p.AllPhenotypeExpectation("AE")
      AEs[count].X = float64(count)

      REs[count].Y = p.AllPhenotypeExpectation("RE")
      REs[count].X = float64(count)

      for i:=0; i<len(p.allBacteria); i++ {
        if p.IsAlive(i) {
          cellCount[p.allBacteria[i].strain+1][count].Y += 1
        }
      }

      cellCount[0][count].Y = float64(len(p.allPredator))

      for i:=0; i<numStrain+1; i++ {
        cellCount[i][count].X = float64(count)
      }
//---------------------------------------------------------------------

      p.CostBasicEnergy(basalMetabolicRate) // metabolism.go

      // Sample from the DNA to update the appropriate field of bacteria
      p.UpdateAllEE()
      p.UpdateAllMovement()
      p.UpdateAllAE()
      p.UpdateAllRE()

      // Delete dead cells and depleted food package from Petri dish
      p.ChecktoDeleteBact()
      p.ChecktoDeleteFood()
      p.ChecktoDeletePred()

      // *** THREE MAIN BACTERIA ABILITIES ***
      p.Replication()
      p.MoveToFood()
      p.Attack()
      p.MoveToBacteria()

      // Last step: DNA mutation may or may not occur
      p.MutateAll() // DNA.go

      pic = p.DrawPetri()
      gifImages = append(gifImages,pic.img)
      count ++
    }
// *** END OF THE MAIN SIMULATION LOOP

// ------------------------------ Generate Plot ------------------------------
    Plot(EEs, "Expected Energy Efficiency Value", "Iteration", "Average Energy Efficiency")

    Plot(movements, "Expected Movement Value", "Iteration", "Average Step Size")

    Plot(AEs, "Expected Antibiotic Enzyme Potency", "Iteration", "Average Potency")

    Plot(REs, "Expected Resistance Enzyme Potency", "Iteration", "Average Potency")

    MultiPlot(cellCount, "Population of Cells per Bacteria Strain", "Iteration", "Number of Cells")

// ----------------------------------------------------------------------------

  return gifImages
}
