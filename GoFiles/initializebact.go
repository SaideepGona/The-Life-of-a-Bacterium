package main

import (
  "math/rand"
  "math"
  "strconv"
  "os"
  "time"
)

func InitialBact(dish Petri, number int) dish{
  dish.size.centerX := strconv.Atoi(os.Args[1])
  dish.size.centerY := strconv.Atoi(os.Args[2])
  rand.Seed(time.Now().UTC().UnixNano())
  dish.size.radius := math.Sqrt(dish.size.centerX*dish.size.centerX+dish.size.centerY*dish.size.centerY)
  for i:=0; i < number; i++{
    x:= rand.Intn(2 * dish.size.radius)
    y:= rand.Intn(2 * dish.size.radius)
  dish.allBacteria[i].size.centerX = x
  dish.allBacteria[i].size.centerY = y
}
  return dish
}
