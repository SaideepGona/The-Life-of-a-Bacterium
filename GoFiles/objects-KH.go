// Kwanho Kim
// 10.21.2017

type Bacteria struct {
  radius float64
  location Location
  ABenzyme ABenzyme
  attackRange float64
  ResistEnzyme ResistEnzyme
  
  linage int
  id int
  energy float64
  dna DNA
}

type Petri struct {
  radius float64
  allBacteria []Bacteria
}


type Location struct {
  petri Petri
  coorX, coorY float64
}



type ABenzyme struct {
  lock int
  potency int
}

type ResistEnzyme struct {
  key int
  potency int
}
