package main
import(
  "fmt"
  "math"
  "math/rand"
  "time"
)

type Bacteria struct {
  position Coords
  sizeRadius float64
  //id int
}

type Petri struct {
  radius float64
  allBacteria []Bacteria
}

type Coords struct {
  coorX float64
	coorY float64
}

//r is the radius of bacteria, R is the radius of petridish, x is the x coordiante of the center of circle and y is the y coordinate of the center of circle
func InField(x, y, X, Y float64, p Petri) bool {
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

/*This function generate a random radian and return it*/
func RandomDelta() float64 {
	randomnumber := rand.Float64()
	randomradian := 2 * randomnumber * math.Pi //the range of theta is between 0 to 2pi
	return randomradian
}

func IsOccupied(x, y float64, i int, p Petri)bool{
  for index:=0; index<len(p.allBacteria); index++{
    if index!=i{
      if p.allBacteria[index].position.coorX==x && p.allBacteria[index].position.coorY==y{
        return true
      }
    }
  }
  return false
}

func IsOverlap(x, y float64, i int, p Petri)bool{
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

/*This function generates a random step that is within the petridish*/
func RandomStep(d, X, Y float64, p Petri) {
  index := rand.Intn(len(p.allBacteria))
  //fmt.Println("index", index)
	a, b := p.allBacteria[index].position.coorX, p.allBacteria[index].position.coorY
	for (a == p.allBacteria[index].position.coorX && b == p.allBacteria[index].position.coorY) || !InField(a, b, X, Y, p) || IsOccupied(a, b, index, p)==true || IsOverlap(a, b, index, p)==true{
    /*fmt.Println("x", a)
    fmt.Println("y", b)
    fmt.Println("in field", InField(a,b,X,Y,p))
    fmt.Println("IsOccupied", IsOccupied(a,b,index, p))
    fmt.Println("IsOverlap", IsOverlap(a, b, index, p))*/
		randomTheta := RandomDelta()
		a = p.allBacteria[index].position.coorX + math.Cos(randomTheta)*d           //a and b are updated and they are the new coordinates
		b = p.allBacteria[index].position.coorY + math.Sin(randomTheta)*d
	}
  p.allBacteria[index].position.coorX = a
  p.allBacteria[index].position.coorY = b
  /*fmt.Println("updateda", a)
  fmt.Println("updatedb", b)*/
  fmt.Println(p)
}

func RandomWalk(d, X, Y float64, p Petri){
  count:=0
  for count<2{
    RandomStep(d, X, Y, p)
    count=count+1
  }
  //fmt.Println("this is count", count)
}

func main(){
  R:=50.0
  X:=R
  Y:=R
  bacteriaSlice:=make([]Bacteria, 0)
  for i:=0; i<2; i++{
    var newBact Bacteria
    fmt.Println(newBact)
    bacteriaSlice=append(bacteriaSlice,newBact)
  }
  var p Petri
  p.allBacteria=bacteriaSlice
  rand.Seed(time.Now().UTC().UnixNano())
  p.allBacteria[0].position.coorX=49
  p.allBacteria[0].position.coorY=49
  p.allBacteria[0].sizeRadius=2.0
  p.allBacteria[1].position.coorX=45
  p.allBacteria[1].position.coorY=45
  p.allBacteria[1].sizeRadius=2.0
  p.radius=50.0
  fmt.Println("p", p)
  //RandomStep(1.0, X, Y, p)
  RandomWalk(5.0, X, Y, p)
}
