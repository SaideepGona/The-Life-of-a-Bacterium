//Zhenyu Yang 2017/10/27

package main

import (
// "fmt"
"image"
"math/rand"
"time"
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

  for i := range p.allBacteria {
    if p.IsLive(i) == true {
      pic.DrawLiveBacteria(p.allBacteria[i].position.coorX , p.allBacteria[i].position.coorY,p.allBacteria[i].sizeRadius)
  } /*else {
      pic.DrawDeadBacteria(p.allBacteria[i].position.coorX , p.allBacteria[i].position.coorY,rBact)
    }*/

  for j := range p.allFoodpack {
    if p.allFoodpack[j].energy > 0.0 {
      pic.DrawFoodpack(p.allFoodpack[j].position.coorX,p.allFoodpack[j].position.coorY,1.0)
  }
}
}
  return pic
}

func (pic *Canvas) DrawLiveBacteria(x,y,r float64) {
   purple := MakeColor(198,177,177)
   pic.SetFillColor(purple)
   pic.Circle(x,y,r)
   pic.Fill()
}

func (pic *Canvas) DrawDeadBacteria(x,y,r float64) {
   deepPurple := MakeColor(0,0,0)
   pic.SetFillColor(deepPurple)
   pic.Circle(x,y,r)
   pic.Fill()
}


func (pic *Canvas) DrawFoodpack(x,y,r float64) {
    // Set Agar to be small yellow dot with a circle of radius 1
      yellow := MakeColor(120,20,120)
      pic.SetFillColor(yellow)
      pic.Circle(x,y,r)
      pic.Fill()
}

func (p *Petri) AnimationPetri() []image.Image {

gifImages := make([]image.Image,0)
 //count := 0
 //p.ChecktoDeleteBact()
 //p.ChecktoDeleteFood()
 count := 0
  pic := p.DrawPetri()
  for p.IsEnd() == false && count < 80 {
    p.ChecktoDeleteBact()
    p.ChecktoDeleteFood()
    p.MoveToFood()
    p.RandomStep(p.radius,p.radius)
    pic = p.DrawPetri()
    gifImages = append(gifImages,pic.img )
    count ++
    //p.ChecktoDeleteBact()
    //p.ChecktoDeleteFood()
  }
  //gifImages[count] = pic.img
  //count ++
  //p.ChecktoDeleteBact()
  //p.ChecktoDeleteFood()
//}

  return gifImages
}

func main() {
  var p Petri
  rand.Seed(time.Now().UTC().UnixNano())
// Call function menetioned before to initialize bacteria and food package
  p.InitializeBact(200,200.0)

  p.InitializeFoodpackage(100,200.0)
  p.InitializeEnergyEfficiency()

  pic := p.DrawPetri()
  pic.SaveToPNG("PetriOrigin.png")

  gifImages := p.AnimationPetri()
   Process(gifImages, "Petri")

/*
 Print coordinates before the bacteria's moving
for j := range p.allBacteria {
  fmt.Println("The position of this bacteria is" , p.allBacteria[j].position)
  fmt.Println("The energy of this bacteria is" , p.allBacteria[j].currentEnergy)
}
for i := range p.allFoodpack {
  fmt.Println("The position of this foodpackage is" , p.allFoodpack[i].position)
  fmt.Println("The energy of this foodpackage is" , p.allFoodpack[i].energy)
}*/

/*
// Set number of moves, this example showed 5.
  p.MovetoToFood()
// Print coordinates after the bacteria's moving
for i := range p.allBacteria {
  fmt.Println("The position of this bacteria is after moving" , p.allBacteria[i].position)
  fmt.Println("The energy of this bacteria is" , p.allBacteria[i].currentEnergy)
}
for i := range p.allFoodpack {
  fmt.Println("The position of this foodpackage is" , p.allFoodpack[i].position)
  fmt.Println("The energy of this foodpackage is" , p.allFoodpack[i].energy)
}
*/

}
// Control with temperature
/*func (pic *Canvas) DrawHotDeadCell(x,y,r float64) { // if the temperature is too hot
  // cell will dissolve its self
    m := 0
    n := 0
    for i:=0; i < 5; i++{
    color := MakeColor(m, n, 255)
    DrawBacteria(x,y,r,pic,color)
    m = m + 51
    n = n + 51
  } // hot to death cell will turn into white
}

func (pic *Canvas) DrawFreezeDeadCell(x,y,r float64) Canvas{ // if the temperature is too cold
  // cell will dissolve its self
    m := 0
    n := 255
    for i := 1; i <= 2 ; i++{
    color := MakeColor(m, 0, n)
    DrawBacteria(x,y,r,pic,color)
    m = m + 64
    n = n - 64
  }
}

func (pic *Canvas) DrawAboveBestTemp(x,y,r int, speed,currenttemp, besttemp,hotdeadtemp float64){
  // when the environment's temperature we above best temp

   if currenttemp > besttemp && currenttemp < hotdeadtemp {
     speed = speed + speed * (currenttemp - besttemp)/10
   }
   if currenttemp > hotdeadtemp {
      speed = 0
      DrawHotDeadCell(x,y,r,pic)
   }
}

func (pic *Canvas) DrawBelowBestTemp(x,y,r int, maxspeed,speed,currenttemp,besttemp, freezedeadtemp float64) Canvas{
    // when the environment's temperature we below best temp
    blue := MakeColor(0,0,255)
    if currentTemp < bestTemp && currentTemp > 0 {
      speed = maxSpeed * currenttemp / bestTemp
      DrawBacteria(x,y,r,blue)
    }
    if currentTemp <= 0 {
      speed = 0
      DrawColdDeadCell(x,y,r,pic)
    }
}

// Control with oxygen
func (pic *Canvas) DrawAbnormalOxy(x,y,r int, oxy,bestOxy,speed,destroyOxy float64){
  white := MakeColor(255,255,255)
  if oxy < destroyOxy {
    speed = speed * oxy/bestOxy
  } else if oxy >= destroyOxy{
    speed = 0
    pic.SetFillColor(white)
    pic.Stroke()
    pic.Circle(x,y,r)
  }
}
*/
/*
func (pic *Canvas) DrawDehydrationBact(x,y,r int, c color.Color, concerntation float64){
  // When a cell is dehydrationing,  the cell is smaller and smaller till vanished.
    white := MakeColor(255,255,255)
    pic.SetFillColor(c)
    r0 := r/5
    for i:= 0; i < 5; i ++ {
      pic.SetFillColor(white)
      pic.Circle(x,y,r)
      pic.SetFillColor(c)
      r = r - r0/5
      pic.Circle(x,y,r)
    }
}

func (pic *Canvas) DrawOsmosisBact(x,y,r int, c color.Color) {
  // When a osmosis cell , the cell become bigger
    white := MakeColor(255,255,255)
    pic.SetFillColor(c)
    r0 := r/20
    for i := 0; i < 4; i ++ {
      pic.SetFillColor(white)
      pic.Circle(x,y,r)
      pic.SetFillColor(c)
      r = r + r0/20
      pic.Circle(x,y,r)
    }
}
*/
