// Zhenyu Yang 2017/11/7
package main

import (
  "strconv"
  "os"
  "fmt"
  "bufio"
  "strings"
  "math"
)
type Bacteria struct {
  size Size
  location Location
  //ABenzyme ABenzyme
  AttackRange float64
  //ResistEnzyme ResistEnzyme
  linage int
  energy float64
}

type Petri struct {
  size Size
  allBacteria []Bacteria
}


type Location struct {
  Petri Petri
  coorX, coorY int
}

type Size struct {
  centerX, centerY float64
  radius float64
}

func ReadFile(filename string) []string {
    // open the file and make sure all went well
    in, err := os.Open(filename)
    if err != nil {
      fmt.Println("Error: couldn’t open the file")
      os.Exit(1)
    }
    // create the variable to hold the lines
    var lines []string = make([]string, 0)
    // for every line in the file
    scanner := bufio.NewScanner(in)
    for scanner.Scan() {
      // append it to the lines slice
      lines = append(lines, scanner.Text())
    }
    // check that all went ok
    if scanner.Err() != nil {
      fmt.Println("Sorry: there was some kind of error during the file reading")
      os.Exit(1)
    }
    // close the file and return the lines
    in.Close()
    return lines
}

func IsIn(x,y,width int) bool{ // to check whether a point is in side a circle in a square
     distanceToCenter :=math.Sqrt(float64((x-width/2)*(x-width/2)+(y-width/2)*(y-width/2)))
     if distanceToCenter > float64(width/2){
       return false
     }
     return true
}

func (p *Petri) ReadToInitialize(filename string) {
  // read the petrifile and modify p Petri
    lines := ReadFile(filename)
  // Convert what's on a file to a slice of string
    width,_ := strconv.Atoi(lines[0])
    k := 0
    p.allBacteria = make([]Bacteria, 10000)
  // The first line in the file should be the width of a petri dish
    for i := 1 ; i < len(lines) ; i ++ {
  // loop through length of lines, should be y coordinate
    for j := 0 ; j < width ; j ++ {
  // loop through length of widthm should be x coodinate
    var items []string = strings.Split(lines[i],"")
    if items[j] == "X" && IsIn(i-1,j,width) {
  // assign values to the bacteria[] slice
    p.allBacteria[k].location.coorX = j
    p.allBacteria[k].location.coorY = i-1
    k++

  }
}
}
}

func (p *Petri) ReadEnergyResourceFile(filename string)  {
     lines :=  ReadFile(filename)
// There are 7 lines in the file, they are agar 100 minerals, pentose
// sunlight, temperature, oxygen, carbon
     var items1 []string = strings.Split(lines[0], " ")
     if items1[0] != "agar" {
       fmt.Println("Error happens with the file, line 1")
       os.Exit(1)
     }
     agar,_ := strconv.ParseFloat(items1[1],64)
// Continue with the second line in the file
     var items2 []string = strings.Split(lines[1], " ")
     if items2[0] != "minerals" {
       fmt.Println("Error happens with the file, line 2")
       os.Exit(1)
     }
     minerals,_ := strconv.ParseFloat(items2[1],64)
// continue with the third line
     var items3 []string = strings.Split(lines[2], " ")
     if items3[0] != "pentose" {
       fmt.Println("Error happens with the file, line 3")
       os.Exit(1)
     }
     pentose,_ := strconv.ParseFloat(items3[1],64)
// 4th line
    var items4 []string = strings.Split(lines[3], " ")
    if items4[0] != "sunlight" {
        fmt.Println("Error happens with the file, line 4")
        os.Exit(1)
      }
     sunlight,_ := strconv.ParseFloat(items4[1],64)
// 5 th line
    var items5 []string = strings.Split(lines[4], " ")
    if items5[0] != "temperature" {
        fmt.Println("Error happens with the file, line 5")
        os.Exit(1)
      }
     temperature,_ := strconv.ParseFloat(items5[1],64)
// 6th line
    var items6 []string = strings.Split(lines[5], " ")
    if items6[0] != "oxygen" {
        fmt.Println("Error happens with the file, line 6")
        os.Exit(1)
      }
     oxygen,_ := strconv.ParseFloat(items6[1],64)
// 7 th line
    var items7 []string = strings.Split(lines[6], " ")
    if items7[0] != "carbon" {
        fmt.Println("Error happens with the file, line 7")
        os.Exit(1)
      }
     carbon,_ := strconv.ParseFloat(items7[1],64)
     for i := range p.allBacteria{
     p.allBacteria[i].energy=(agar+minerals+pentose+sunlight+temperature+oxygen+carbon)/7.0
   }
}

func main(){
  var p Petri
  p.ReadToInitialize("petriBact.txt")
  p.ReadEnergyResourceFile("nutrientFile.txt")
  fmt.Println(p.allBacteria[0].location.coorX,p.allBacteria[0].location.coorY)
  fmt.Println(p.allBacteria[1].location.coorX,p.allBacteria[1].location.coorY)
  fmt.Println(p.allBacteria[0].energy)
}
