package main
import(
  "fmt"
)

/* This function draw the lines in the board*/
func DrawGridLines(pic Canvas, array []float64) {
  //find the width and height of the canvas
  w, h := pic.width, pic.height
  //divide the width of the canvas by length of the array
  tempWidth := w/len(array)
  fmt.Println("tempWidth", tempWidth)
  largest:=0.0
  //find the largest number in the array
  for i:=0; i< len(array); i++{
    if array[i]>largest{
      largest=array[i]
    }
  }
  templargest:=largest+2
  //divide height of the canvas by largest number+2.
  tempWidth2 := float64(h)/templargest
  fmt.Println("tempWidth2", tempWidth2)
  black := MakeColor(0, 0, 0)
  //set the current line color to black
  pic.SetStrokeColor(black)
  //move the current point to position(float64(0), float64(float(h)-array[0]*tempWidth2))
  pic.MoveTo(float64(0), float64(float64(h)-array[0]*tempWidth2))
  for index:=0; index< len(array); index++{
    fmt.Println("index*tempWidth", index*tempWidth)
    fmt.Println("array[index]*tempWidth2", array[index]*tempWidth2)
    fmt.Println("")
    //logically draw a line from the current point to (float64(index*tempWidth), float64(float64(h)-array[index]*tempWidth2))
    //and make (float64(index*tempWidth), float64(float64(h)-array[index]*tempWidth2)) the current point using
    //the current line color and current line width
    pic.LineTo(float64(index*tempWidth), float64(float64(h)-array[index]*tempWidth2))
    pic.MoveTo(float64(index*tempWidth), float64(float64(h)-array[index]*tempWidth2))
  }
  //actually draw the lines specified by LineTo calls and clear the pending lines.
  pic.Stroke()
  red := MakeColor(255, 0, 0)
  //set the current fill color to red
  pic.SetFillColor(red)
  //draw te dots 
  for j:=0; j< len(array); j++{
    pic.Circle(float64(j*tempWidth), float64(float64(h)-array[j]*tempWidth2), 5)
  }
  pic.Fill()
}

func main(){
  array:=[]float64{1, 2, 3, 4, 5, 5, 5}
  pic := CreateNewCanvas(300, 300)
  DrawGridLines(pic, array)
  pic.SaveToPNG("test.png")
}
