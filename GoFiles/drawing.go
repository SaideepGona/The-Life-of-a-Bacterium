//Zhenyu Yang 2017/10/27
// microbals competition: Chthamalus/Balanus,
package main



func DrawPetri(width int) Canvas{
   white:= MakeColor(255,255,255)
   black := MakeColor(0, 0, 0)
   pic:= CreateNewCanvas(width,width)
   pic.SetFillColor(black)
   pic.Clear()
   pic.SetFillColor(white)
   pic.Circle(width/2,width/2,width/2)
   return pic
}

func (pic *Canvas) DrawBacteria(x,y,r float64, color color.Color){
   pic.SetFillColor(color)
   pic.Circle(x,y,r)
}

// Control with temperature
func (pic *Canvas) DrawHotDeadCell(x,y,r float64) { // if the temperature is too hot
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

func (pic *Canvas) DrawAgarRandom(numAgars int,  width int) Canvas{
    // Set Agar to be small yellow dot with a circle of radius 1
    yellow := MakeColor(255,0,255)
    pic.SetFillColor(yellow)
    for i:= 0; i < numAgars ; i++{
      x := rand.Intn(width)
      y := rand.Intn(width)
      pic.MoveTo(x,y)
      pic.Circle(x,y,1)
      pic.Stoke()
  }
}

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
