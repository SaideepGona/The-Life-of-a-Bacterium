package main
import(
  "fmt"
  "math"
  "math/rand"
)

//r is the radius of bacteria, R is the radius of petridish, x is the x coordiante of the center of circle and y is the y coordinate of the center of circle
func InField(x, y, r, X, Y, R float64) bool {
  total:=math.Pow(X-x, 2)+math.Pow(Y-y, 2)
  distance:=math.Sqrt(total)
  //fmt.Println("this is distance", distance)
  difference:=R-r
  //fmt.Println("this is difference", difference)
	if distance>difference{
		return false
	}else{
	  return true
  }
}

/*This function generate a random radian and return it*/
func RandomDelta() float64 {
	randomnumber := rand.Float64()
	randomradian := 2 * randomnumber * math.Pi //the range of theta is between 0 to 2pi
  //fmt.Println("randomradian", randomradian)
	return randomradian
}

func IsOccupied(XOccupied, YOccupied []float64, x, y float64)bool{
  for index:=0; index<len(XOccupied); index++{
    if XOccupied[index]==x || YOccupied[index]==y{
      return true
    }
  }
  return false
}

func IsOverlap(XOccupied, YOccupied []float64, x, y, limit float64)bool{
  for index:=0; index<len(XOccupied); index++{
    total:=math.Pow(XOccupied[index]-x, 2)+math.Pow(XOccupied[index]-y, 2)
    distance:=math.Sqrt(total)
    //fmt.Println("distance", distance)
    if distance<limit{
      return true
    }
  }
  return false
}

/*This function generates a random step that is within the petridish*/
func RandomStep(x, y, r, d, X, Y, R float64, XOccupied, YOccupied []float64) (float64, float64, float64, float64) {
	a, b := x, y
  limit:=2*r
  //count:=0
	for (a == x && b == y) || !InField(a, b, r, X, Y, R) || IsOccupied(XOccupied, YOccupied, a, b)==true || IsOverlap(XOccupied, YOccupied, a, b, limit)==true{
		randomTheta := RandomDelta()
		a = x + math.Cos(randomTheta)*d           //a and b are updated and they are the new coordinates
		b = y + math.Sin(randomTheta)*d
    //count=count+1
    //fmt.Println("a", a)
    //fmt.Println("b", b)
    //fmt.Println("check", a==x && b==y)
    //result:=InField(a, b, r, X, Y, R)
    //fmt.Println("result", result)
    //result2:=IsOccupied(XOccupied, YOccupied, a, b)
    //fmt.Println("result2", result2)
    //result3:=IsOverlap(XOccupied, YOccupied, a, b, limit)
    //fmt.Println("result3", result3)
    //fmt.Println("This is count", count)
	}
  //fmt.Println("XOccupied", XOccupied)
  //fmt.Println("YOccupied", YOccupied)
	return a, b, x, y
}

func UpdateXYOccupied(XOccupied, YOccupied []float64, x, y, a, b float64)([]float64, []float64){
  XOccupied=append(XOccupied, a)
  YOccupied=append(YOccupied, b)
  //fmt.Println("this is XOccupied", XOccupied)
  //fmt.Println("this is YOccupied", YOccupied)
  for index:=0; index<len(XOccupied); index++{
    if XOccupied[index]==x{
      XOccupied=append(XOccupied[0:index],XOccupied[index+1:]...)                 //delete x and y in XOccupied and YOccupied
      YOccupied=append(YOccupied[0:index],YOccupied[index+1:]...)
    }
  }
  //fmt.Println("updated XOccupied", XOccupied)
  //fmt.Println("updated YOccupied", YOccupied)
  return XOccupied, YOccupied
}

func RandomWalk(r, stepSize, R float64, XOccupied, YOccupied []float64) {
  x_0 := r    //The starting position of bacteria is (r,r) which is the center of petridish
  y_0 := r
  x := x_0
	y := y_0
  prex := 0.0
  prey := 0.0
  i:=0
  XOccupied=append(XOccupied, x)
  YOccupied=append(YOccupied, y)
  fmt.Println("XOccupied", XOccupied)
  fmt.Println("YOccupied", YOccupied)
  for i<3{
    x,y,prex, prey=RandomStep(x, y, r, stepSize, x_0, y_0, R, XOccupied, YOccupied)
    XOccupied, YOccupied=UpdateXYOccupied(XOccupied, YOccupied, prex, prey, x, y)
    fmt.Println("XOccupied", XOccupied)
    fmt.Println("YOccupied", YOccupied)
    fmt.Println(x, y)
    i=i+1
  }
}

func main(){
  var XOccupied []float64          //xbool is the slice which records all the x index of the coordinate which is occupied
  var YOccupied []float64          //ybool is the slice which records all the y index of the coordinate which is occupied
                                   //together (XOccupied[index] and YOccupied[index]) represents coordinate of bacteria in the petridish
  RandomWalk(2.0, 5.0, 50.0, XOccupied, YOccupied)
}
