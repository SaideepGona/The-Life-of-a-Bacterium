// Kwanho Kim
// 10.21.2017

type Bacteria struct {
  radius float64
  location Location
  ABenzyme ABenzyme
  AttackRange float64
  ResistEnzyme ResistEnzyme
  
  linage int
  id int
  energy float64
  dna DNA
}

type Petri struct {
  size Size
  allBacteria []*Bacteria
  counter int
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
