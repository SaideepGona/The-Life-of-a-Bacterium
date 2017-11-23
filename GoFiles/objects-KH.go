// Kwanho Kim
// 10.21.2017

type Bacteria struct {
  size Size
  location Location
  ABenzyme ABenzyme
  AttackRange float64
  ResistEnzyme ResistEnzyme
  linage int
  energy float64
}

type Petri struct {
  size Size
  allBacteria []*Bacteria
}


type Location struct {
  Petri Petri
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
