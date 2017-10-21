// Kwanho Kim
// 10.21.2017

type Bacteria struct {
  size size
  location location
}

type Petri struct {
  size size
}

type location struct {
  Petri Petri
  coorX, coorY float64
}

type size struct {
  centerX, centerY float64
  radius float64
}
