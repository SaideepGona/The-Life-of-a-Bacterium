// 11/10/2017
// DO NOT CHANGE THESE OBJECTS!!
// WE ONLY MODIFY THIS AS A GROUP

type Bacteria struct {

  linage int

  sizeRadius float64
  position Coords
  detectRadius float64

  currentEnergy float64
  energyCapacity float64
  energyEfficiancy float64

  ABenzyme ABenzyme
  attackRange float64
  resistEnzyme ResistEnzyme

  movement float64

  repEnergy float64

  dna DNA

}

type Coords struct {
  coorX float64
	coorY float64
}

type Petri struct {
  radius float64
  allBacteria []Bacteria
}

type ABenzyme struct {
  lock int
  potency float64
}

type ResistEnzyme struct {
  key int
  potency float64
}

type FoodPackage struct {
  position Coords
  theta float64
}
