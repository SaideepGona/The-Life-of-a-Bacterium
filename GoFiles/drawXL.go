package main
import(
  "fmt"
  //"math"
  //"math/rand"
  //"time"
)

func DrawGridLines(pic Canvas, array []float64) {
  w, h := pic.width, pic.height
  tempWidth := w/len(array)
  fmt.Println("tempWidth", tempWidth)
  largest:=0.0
  for i:=0; i< len(array); i++{
    if array[i]>largest{
      largest=array[i]
    }
  }
  tempWidth2 := float64(h)/largest
  fmt.Println("tempWidth2", tempWidth2)
  black := MakeColor(0, 0, 0)
  pic.SetStrokeColor(black)
  pic.MoveTo(float64(0), float64(float64(h)-array[0]*tempWidth2))
  for index:=0; index< len(array); index++{
    fmt.Println("index*tempWidth", index*tempWidth)
    fmt.Println("array[index]*tempWidth2", array[index]*tempWidth2)
    fmt.Println("")
    pic.LineTo(float64(index*tempWidth), float64(float64(h)-array[index]*tempWidth2))
    pic.MoveTo(float64(index*tempWidth), float64(float64(h)-array[index]*tempWidth2))
  }
  pic.Stroke()
  red := MakeColor(255, 0, 0)
  pic.SetFillColor(red)
  for j:=0; j< len(array); j++{
    pic.Circle(float64(j*tempWidth), float64(float64(h)-array[j]*tempWidth2), 5)
  }
  /*pic.Circle(float64(0), float64(h), float64(5))
  pic.Circle(float64(0), float64(h-30), float64(5))
  pic.Circle(float64(50), float64(h-60), float64(5))
  pic.Circle(float64(100), float64(h-150), float64(5))
  pic.Circle(float64(150), float64(h-270), float64(5))
  pic.Circle(float64(200), float64(h-210), float64(5))
  pic.Circle(float64(250), float64(h-300), float64(5))*/
  //pic.Circle(float64())
  pic.Fill()
}

func main(){
  array:=[]float64{10, 20, 50, 90, 70, 100}
  pic := CreateNewCanvas(300, 300)
  DrawGridLines(pic, array)
  pic.SaveToPNG("test.png")
}
