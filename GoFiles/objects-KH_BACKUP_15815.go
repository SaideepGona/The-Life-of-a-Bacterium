// Kwanho Kim
// 10.21.2017

type Bacteria struct {
  size Size
  location Location
  ABenzyme ABenzyme
  AttackRange float64
  ResistEnzyme ResistEnzyme
  linage int
}

type Petri struct {
  size Size
<<<<<<< HEAD
  allBacteria [][]*Bacteria
=======
  allBacteria []*Bacteria
>>>>>>> acbf4197a7bd5cd2ecbab004f1325802e18773b2
}


type Location struct {
<<<<<<< HEAD
  petri Petri
=======
  Petri Petri
>>>>>>> acbf4197a7bd5cd2ecbab004f1325802e18773b2
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
