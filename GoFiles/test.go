package main

import "fmt"


func main() {

	bacteriaSlice := make([]*Bacteria, 0)

    for i := 0; i < 10; i++ {
		var newBact Bacteria
		bacteriaSlice = append(bacteriaSlice, &newBact)
	}

	fmt.Println(bacteriaSlice[0])
	for i := 0; i <5; i++{
		ModBact(bacteriaSlice[i])
	}
	fmt.Println(bacteriaSlice)
}

func ModBact(bact *Bacteria) {
	bact.AttackRange = 5.5
}



type Bacteria struct {
  size Size
  location Location
  ABenzyme ABenzyme
  AttackRange float64
  ResistEnzyme ResistEnzyme
  linage int
  id int
}

type Location struct {
  petri Petri
  coorX, coorY float64
}

type Size struct {
  centerX, centerY float64
  radius float64
}

type ABenzyme struct {
  lock int
  potency int
}

type ResistEnzyme struct {
  key int
  potency int
}

type Petri struct {
  size Size
  allBacteria []*Bacteria
  counter int
}