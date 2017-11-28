package main
import(
  "fmt"
  "math"
  "math/rand"
  "time"
)

/*type Bacteria struct {
  position Coords
  sizeRadius float64
  currentEnergy float64
  energyEfficiency float64
  stepSize float64
}

type Petri struct {
  radius float64
  allBacteria []Bacteria
}

type Coords struct {
  coorX float64
	coorY float64
}*/

//r is the radius of bacteria, x and y are the x and y coordiante of the center of circle, and X and Y are the x and y coordinate of the petridish
func InField(x, y, X, Y float64, p Petri) bool {
  total:=math.Pow(X-x, 2)+math.Pow(Y-y, 2)
  distance:=math.Sqrt(total)
  difference:=p.radius-p.allBacteria[0].sizeRadius
	if distance>difference{
		return false
	}else{
	  return true
  }
}

/*This function generate a random radian and return it*/
func RandomDelta() float64 {
	randomnumber := rand.Float64()
  //the range of theta is between 0 to 2pi
	randomradian := 2 * randomnumber * math.Pi
	return randomradian
}

/*This function checkes if there is complete overlap among (x,y) and the location of all the bacteira in the slice*/
func IsOccupied(x, y float64, i int, p Petri)bool{
  for index:=0; index< len(p.allBacteria); index++{
    if index!=i{
      //If (x,y) is the same as the the location of bacteria in the slice, return true.Otherwise, return false
      if p.allBacteria[index].position.coorX==x && p.allBacteria[index].position.coorY==y{
        return true
      }
    }
  }
  return false
}

/*This function checkes if there is partial overlap among (x,y) and the location of all the bacteria in the slice*/
func IsOverlap(x, y float64, i int, p Petri)bool{
  for index:=0; index<len(p.allBacteria); index++{
    if index!=i{
      limit:=p.allBacteria[i].sizeRadius+p.allBacteria[index].sizeRadius
      total:=math.Pow(p.allBacteria[index].position.coorX-x, 2)+math.Pow(p.allBacteria[index].position.coorY-y, 2)
      distance:=math.Sqrt(total)
      if distance<limit{
        return true
      }
    }
  }
  return false
}

/*This function randomize the index of allbacteria list, which suggests the order of movement for each bacteria*/
func PermuteList(p Petri)[]int{
  list := rand.Perm(len(p.allBacteria))
  for i, _ := range list {
    list[i]++
  }
  fmt.Println("list", list)
  return list
}

/*This function generates a random step that is within the petridish*/
func (p *Petri) RandomStep(X, Y float64) {
  randomlist:=PermuteList(*p)
  for index:=0; index< len(randomlist); index++{
     count:=0
	   a, b := (*p).allBacteria[index].position.coorX, (*p).allBacteria[index].position.coorY
     energyconsumed:=EnergyBurnMovement(*p, index)
     if energyconsumed<=(*p).allBacteria[index].currentEnergy{
	      for (a == (*p).allBacteria[index].position.coorX && b == (*p).allBacteria[index].position.coorY) || !InField(a, b, X, Y, *p) || IsOccupied(a, b, index, *p)==true || IsOverlap(a, b, index, *p)==true{
		      randomTheta := RandomDelta()
          //a and b are updated and they are the new coordinates
		      a = p.allBacteria[index].position.coorX + math.Cos(randomTheta)*p.allBacteria[index].stepSize
		      b = p.allBacteria[index].position.coorY + math.Sin(randomTheta)*p.allBacteria[index].stepSize
          /*count keeps track of the number of time the loop is iterated. if the loop has been iterated for 100 times and the generated coordinate are still not valid,
          exit the loop*/
          count=count+1
          if count==100{
            break
          }
        }
	   }

     pic.MoveTo(p.allBacteria[index].position.coorX, p.allBacteria[index].position.coorY)
     /*if count is smaller than 100, update the location of bacteira. Otherwise, keep the original location of the bacteira*/
     if count<100{
       p.allBacteria[index].position.coorX = a
       p.allBacteria[index].position.coorY = b
       p.allBacteria[index].currentEnergy=p.allBacteria[index].currentEnergy-energyconsumed
     }
  }
  fmt.Println(p)
}

/*calculate the energy consumed for the movement*/
func EnergyBurnMovement(p Petri, index int) float64{
  energyConsumption:=p.allBacteria[index].stepSize*1*p.allBacteria[index].energyEfficiency
  return energyConsumption
}

/*func main(){
  R:=50.0
  X:=R
  Y:=R
  bacteriaSlice:=make([]Bacteria, 0)
  for i:=0; i<2; i++{
    var newBact Bacteria
    fmt.Println("newBact", newBact)
    bacteriaSlice=append(bacteriaSlice,newBact)
  }
  var p Petri
  p.allBacteria=bacteriaSlice
  rand.Seed(time.Now().UTC().UnixNano())
  p.allBacteria[0].position.coorX=49
  p.allBacteria[0].position.coorY=49
  p.allBacteria[0].sizeRadius=2.0
  p.allBacteria[0].stepSize=5.0
  p.allBacteria[0].energyEfficiency=0.6
  p.allBacteria[0].currentEnergy=20
  p.allBacteria[1].position.coorX=45
  p.allBacteria[1].position.coorY=45
  p.allBacteria[1].sizeRadius=2.0
  p.allBacteria[1].stepSize=6.0
  p.allBacteria[1].energyEfficiency=0.8
  p.allBacteria[1].currentEnergy=30
  p.radius=50.0
  fmt.Println("p", p)
  p.RandomStep(X, Y)
  p.RandomStep(X, Y)
  p.RandomStep(X, Y)*/
}
