// Kwanho Kim
// 10.21.2017

type Bacteria struct {
  size size
  location location
  ABenzyme ABenzyme
  AttackRange float64
  ResistEnzyme ResistEnzyme
}

type Petri struct {
  size size
  allBacteria [][]*Bacteria
}


type location struct {
  Petri Petri
  coorX, coorY float64
}

type size struct {
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
