package main

import (
//  "fmt"
  "math"
  "math/rand"
  //"time"
)

//r is the radius of bacteria, R is the radius of petridish, x is the x coordiante of the center of circle and y is the y coordinate of the center of circle
func InField(x, y, X, Y float64, p *Petri) bool {
  total:=math.Pow(X-x, 2)+math.Pow(Y-y, 2)
  distance:=math.Sqrt(total)
  difference:=p.radius-p.allBacteria[0].sizeRadius
  /*fmt.Println("distance", distance)
  fmt.Println("difference", difference)*/
	if distance>difference{
		return false
	}else{
	  return true
  }
}

func InFieldPredator(x, y, X, Y float64, p *Petri) bool {
  total:=math.Pow(X-x, 2)+math.Pow(Y-y, 2)
  distance:=math.Sqrt(total)
  difference:=p.radius-p.allPredator[0].sizeRadius
	if distance>difference{
		return false
	}else{
	  return true
  }
}

func IsOccupied(x, y float64, i int, p *Petri)bool{
  for index:=0; index<len(p.allBacteria); index++{
    if index!=i{
      if p.allBacteria[index].position.coorX==x && p.allBacteria[index].position.coorY==y{
        return true
      }
    }
  }
  return false
}

/*This function checkes if there is complete overlap among (x,y) and the location of all the predators in the slice*/
func IsOccupiedPredator(x, y float64, i int, p Petri)bool{
  for index:=0; index< len(p.allPredator); index++{
    if index!=i{
      //If (x,y) is the same as the the location of bacteria in the slice, return true.Otherwise, return false
      if p.allPredator[index].position.coorX==x && p.allPredator[index].position.coorY==y{
        return true
      }
    }
  }
  return false
}

func IsOverlap(x, y float64, i int, p *Petri)bool{
  for index:=0; index<len(p.allBacteria); index++{
    if index!=i{
      limit:=p.allBacteria[i].sizeRadius+p.allBacteria[index].sizeRadius
      total:=math.Pow(p.allBacteria[index].position.coorX-x, 2)+math.Pow(p.allBacteria[index].position.coorY-y, 2)
      distance:=math.Sqrt(total)
      /*fmt.Println("x", x)
      fmt.Println("y", y)
      fmt.Println("p.allBacteria[index].position.coorX", p.allBacteria[index].position.coorX)
      fmt.Println("p.allBacteria[index].position.coorY", p.allBacteria[index].position.coorY)
      fmt.Println("this is distance", distance)*/
      if distance<limit{
        return true
      }
    }
  }
  return false
}

/*This function checkes if there is partial overlap among (x,y) and the location of all the bacteria in the slice*/
func IsOverlapPredator(x, y float64, i int, p Petri)bool{
  for index:=0; index< len(p.allPredator); index++{
    if index!=i{
      limit:=p.allPredator[i].sizeRadius+p.allPredator[index].sizeRadius
      total:=math.Pow(p.allPredator[index].position.coorX-x, 2)+math.Pow(p.allPredator[index].position.coorY-y, 2)
      distance:=math.Sqrt(total)
      if distance<limit{
        return true
      }
    }
  }
  return false
}

func PermuteList(p *Petri)[]int{
  list := rand.Perm(len(p.allBacteria))
  for i, _ := range list {
    list[i]++
  }
  //fmt.Println("list", list)
  return list
}

/*This function generates a random step that is within the petridish*/
func (p *Petri) RandomStep(X, Y float64) {
  //fmt.Println("index", index)
  randomlist:=PermuteList(p)
  for index:=0; index< len(randomlist); index++ {
     count:=0
	   a, b := (p).allBacteria[index].position.coorX, (p).allBacteria[index].position.coorY
     energyConsumed:=p.allBacteria[index].stepSize*p.allBacteria[index].energyEfficiency*.7
     //fmt.Println("energyconsumed", energyconsumed)
     if energyConsumed<=(p).allBacteria[index].currentEnergy && p.IsFoodAround(index) == false && p.allBacteria[index].hasRep == false {
	      for (a == (p).allBacteria[index].position.coorX && b == (p).allBacteria[index].position.coorY) || !InField(a, b, X, Y, p) || IsOccupied(a, b, index, p)==true || IsOverlap(a, b, index, p)==true{
          /*fmt.Println("x", a)
          fmt.Println("y", b)
          fmt.Println("in field", InField(a,b,X,Y,p))
          fmt.Println("IsOccupied", IsOccupied(a,b,index, p))
          fmt.Println("IsOverlap", IsOverlap(a, b, index, p))*/
		      randomTheta := RandomTheta()
		      a = p.allBacteria[index].position.coorX + math.Cos(randomTheta)*p.allBacteria[index].stepSize           //a and b are updated and they are the new coordinates
		      b = p.allBacteria[index].position.coorY + math.Sin(randomTheta)*p.allBacteria[index].stepSize
          count=count+1
          if count==100{
            break
          }
        }
	   }
     if count<100 {
       p.allBacteria[index].position.coorX = a
       p.allBacteria[index].position.coorY = b
       p.allBacteria[index].currentEnergy -= energyConsumed*.5
     }
       /*fmt.Println("updateda", a)
      fmt.Println("updatedb", b)*/
  //fmt.Println(p)
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

func (p *Petri) IsFull(i int) bool {
  // The bacteria will randomly move if it is full. If not, it should look for
  //food itsself
   if p.allBacteria[i].currentEnergy < p.allBacteria[i].energyCapacity {
     return false
   }
   return true
}

// THIS IS FOR BACTERIA MOVEMENT
func (p *Petri) MoveToFood() {
  // first, range through all the bacteria i a petri dish
  for i := 0; i < len(p.allBacteria); i ++ {
    if p.IsFoodAround(i) == true && p.allBacteria[i].hasRep == false {
    //fmt.Println(len(p.allBacteria))
    bactMaxEnergy := p.allBacteria[i].energyCapacity
    xBact := p.allBacteria[i].position.coorX
    yBact := p.allBacteria[i].position.coorY
  // k and minDistance stands for which foodpackge were detected and how much
  // distance it has.
    k,minDistance :=  MinDisFood(p.allFoodpack,xBact,yBact,p.radius)
    energyConsumption := minDistance*p.allBacteria[i].energyEfficiency*.7
    p.allBacteria[i].currentEnergy -= energyConsumption*.5
  // To check the distance between foodpackage and bacteria are in the detection range
    if p.IsAlive(i) == true && p.IsFull(i) == false && minDistance < p.allBacteria[i].stepSize {
  // Move the coordinate of bacteria to the nearest foodpackge that has energy and
  // inside its detection range
      p.allBacteria[i].position.coorX = p.allFoodpack[k].position.coorX
      p.allBacteria[i].position.coorY = p.allFoodpack[k].position.coorY
  // If the bacteria need more than the food pakage contained, it just simply take
  // all the energy.
      if p.allBacteria[i].currentEnergy < bactMaxEnergy - p.allFoodpack[k].energy {
      p.allBacteria[i].currentEnergy = p.allBacteria[i].currentEnergy + p.allFoodpack[k].energy
      //p.allBacteria[i].sizeRadius += 0.01
      p.allFoodpack[k].energy = 0.0
      } else {
  // If the bacteria need less energy that food pakage contained, it only takes
  // the amount it want to be full.
        p.allBacteria[i].currentEnergy = bactMaxEnergy
        //p.allBacteria[i].sizeRadius += 0.005
        p.allFoodpack[k].energy = p.allFoodpack[k].energy - (bactMaxEnergy - p.allBacteria[i].currentEnergy)
      }
    }
   } else if p.IsFoodAround(i) == false && p.allBacteria[i].hasRep == false {
     p.RandomStep(p.radius,p.radius)
   }
/*
   // After taking one step, is it lands on a drug molecule, it takes damage.
   for index:=0; index< len(p.allDrugpack); index++ {
     if DistanceToDrugPackage(p.allBacteria[i].position.coorX, p.allBacteria[i].position.coorY, p.allDrugpack[index].position.coorX, p.allDrugpack[index].position.coorY)<=p.allBacteria[i].sizeRadius{
       if p.allDrugpack[index].lock == p.allBacteria[i].resistEnzyme.key{
         if p.allDrugpack[index].potency>p.allBacteria[i].resistEnzyme.potency{
           difference:=p.allDrugpack[index].potency-p.allBacteria[i].resistEnzyme.potency
           p.allBacteria[i].currentEnergy=p.allBacteria[i].currentEnergy-difference
         }
       } else {
         // if the drug's lock is not the same as bacteria's key, bacteria lose the same amount of energy as the potency of drug package.
         p.allBacteria[i].currentEnergy=p.allBacteria[i].currentEnergy-p.allDrugpack[index].potency
       }
       p.allDrugpack = append(p.allDrugpack[:index], p.allDrugpack[index+1:]...)
     }
   } // end of for loop (drug effect)
*/
  } // end of for loop going through the slice of bacteria
}

// The predator will randomly move if it is full. If not, it should look for bacteria.
func (p *Petri) IsRandomMovePredator(i int) bool {
   if p.allPredator[i].currentEnergy < p.allPredator[i].energyCapacity {
     return false
   }
   return true
}

/*calculate the distance between the current predator and its closest prey. The longest distance in a circle is 2 * radius
return the minimum distance and the index of the bacteria*/
 func MinDisPredator(allBacteria []Bacteria,xPred,yPred,radius float64) (int,float64){
   minDistance := 2 * radius
   var j int         // j is used to catch which foodpackge to get
   for k := 0; k < len(allBacteria); k ++ {
       xBact := allBacteria[k].position.coorX
       yBact := allBacteria[k].position.coorY
       distance := math.Sqrt((xPred-xBact)*(xPred-xBact) + (yPred-yBact)*(yPred-yBact))
       if distance < minDistance {
        minDistance  =  distance
        j = k
     }
 }
    return j,minDistance
 }

func (p *Petri) MoveToBacteria() {
   // first, range through all the bacteria i a petri dish
   for i := 0; i < len(p.allPredator); i ++ {
     if p.IsBacteriaAround(i) == true && p.allPredator[i].hasRep == false {
     //fmt.Println(len(p.allBacteria))
     predMaxEnergy := p.allPredator[i].energyCapacity
     xPred := p.allPredator[i].position.coorX
     yPred := p.allPredator[i].position.coorY
   // k and minDistance stands for which bacterias were detected and how much
   // distance it has.
     k,minDistance :=  MinDisPredator(p.allBacteria,xPred,yPred,p.radius)
     energyConsumption := minDistance*p.allPredator[i].energyEfficiency
     p.allPredator[i].currentEnergy -= energyConsumption
   //  fmt.Println("p.allPredator[i].currentEnergybeforeupdate", p.allPredator[i].currentEnergy)
   // To check the distance between predator and bacteria are in the detection range
     if p.IsLivePredator(i) == true && p.IsRandomMovePredator(i) == false && minDistance < p.allPredator[i].stepSize {
   // Move the coordinate of bacteria to the nearest foodpackge that has energy and
   // inside its detection range
       p.allPredator[i].position.coorX = p.allBacteria[k].position.coorX
       p.allPredator[i].position.coorY = p.allBacteria[k].position.coorY
   // If the predator need more than the energy the bacteria has, it just simply take
   // all the energy of bacteria.
       if p.allPredator[i].currentEnergy < predMaxEnergy - p.allBacteria[k].currentEnergy {
       p.allPredator[i].currentEnergy = p.allPredator[i].currentEnergy + p.allBacteria[k].currentEnergy
       p.allBacteria[k].currentEnergy = 0.0
       p.ChecktoDeleteBact()
       } else {
   // If the predator need less energy than the energy the bacteria has, it only takes
   // the amount it want to be full.
         p.allPredator[i].currentEnergy = predMaxEnergy
         p.allBacteria[k].currentEnergy = 0.0
         p.ChecktoDeleteBact()
       }
     }
    } else if p.IsBacteriaAround(i) == false && p.allPredator[i].hasRep == false {
    count2:=0
    a, b := (*p).allPredator[i].position.coorX, (*p).allPredator[i].position.coorY
    energyconsumed:=p.allPredator[i].stepSize*p.allPredator[i].energyEfficiency
    if energyconsumed<=(*p).allPredator[i].currentEnergy{
       for (a == (*p).allPredator[i].position.coorX && b == (*p).allPredator[i].position.coorY) || !InFieldPredator(a, b, p.radius, p.radius, p) || IsOccupiedPredator(a, b, i, *p)==true || IsOverlapPredator(a, b, i, *p)==true{
         randomTheta := RandomTheta()
         //a and b are updated and they are the new coordinates
         a = p.allPredator[i].position.coorX + math.Cos(randomTheta)*p.allPredator[i].stepSize
         b = p.allPredator[i].position.coorY + math.Sin(randomTheta)*p.allPredator[i].stepSize
         /*count keeps track of the number of time the loop is iterated. if the loop has been iterated for 100 times and the generated coordinate are still not valid,
         exit the loop*/
         count2=count2+1
         if count2==100{
           break
         }
       }
    }

    //pic.MoveTo(p.allBacteria[index].position.coorX, p.allBacteria[index].position.coorY)
    /*if count is smaller than 100, update the location of predator. Otherwise, keep the original location of the predator*/
    if count2<100{
      p.allPredator[i].position.coorX = a
      p.allPredator[i].position.coorY = b
      p.allPredator[i].currentEnergy=p.allPredator[i].currentEnergy-energyconsumed
    }
    //fmt.Println("p.allPredator[i].currentEnergythere", p.allPredator[i].currentEnergy)
   } // end of else statement

   // *** PREDATOR SPECIFIC DRUG ***
   for index:=0; index< len(p.allPredKill); index++ {
     if DistanceToDrugPackage(p.allPredator[i].position.coorX, p.allPredator[i].position.coorY, p.allPredKill[index].position.coorX, p.allPredKill[index].position.coorY)<=p.allPredator[i].sizeRadius{
       p.allPredator[i].currentEnergy -= p.allPredKill[index].potency
       p.allPredKill = append(p.allPredKill[:index], p.allPredKill[index+1:]...)
     }
   } // end of for loop (drug effect)

  } // end of for loop going through the slice of predators
}

func (p *Petri) IsBacteriaAround(index int) bool {
  for i := 0; i < len(p.allBacteria); i ++{
    xf := p.allBacteria[i].position.coorX
    yf := p.allBacteria[i].position.coorY
    xb := p.allPredator[index].position.coorX
    yb := p.allPredator[index].position.coorY
    if math.Sqrt((yb-yf)*(yb-yf)+(xb-xf)*(xb-xf)) < p.allPredator[index].stepSize{
      return true
    }
  }
  return false
}

func DistanceToDrugPackage(bactX, bactY, drugPackX, drugPackY float64) float64 {
   distance := math.Sqrt((bactX-drugPackX)*(bactX-drugPackX) + (bactY-drugPackY)*(bactY-drugPackY))
   //fmt.Println("distance", distance)
   return distance
}

func (bact *Bacteria) UpdateMovement() {

	// Updates energy efficiency
	newMovementRaw := bact.dna.PhenotypeAverage("Movement")
	newMovement := Logistic(newMovementRaw, bact.dna.phenotypes["Movement"].aggFuncArgs)
	bact.stepSize = newMovement

}

func (p *Petri) UpdateAllMovement() {
	//fmt.Println("UpdateAllMovement")
	// Updates the energy efficiency for all bacteria in petri dish
	for i := 0; i < len(p.allBacteria); i++{
		p.allBacteria[i].UpdateMovement()
	}
}
