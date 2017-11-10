// Kwanho Kim
// 10.21.2017

type Bacteria struct {
  size Size
  location Location
  ABenzyme ABenzyme
  attackRange float64
  ResistEnzyme ResistEnzyme
  linage int
  id int
  energyCap int
  senseRange float64
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
