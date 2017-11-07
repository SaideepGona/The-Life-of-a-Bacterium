package main

import (
  "strconv"
  "os"
  "fmt"
  "bufio"
  "strings"
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
  allBacteria []*Bacteria
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

func (p *Petri)ReadToInitialize(filename string) {
  // read the petrifile and modify p Petri
    lines := ReadFile("pertriBact.txt")
  // Convert what's on a file to a slice of string
    width,_ := strconv.Atoi(lines[0])
  // The first line in the file should be the width of a petri dish
    for i := 1 ; i < len(lines) ; i++ {
      for j := 0 ; j < width ; j ++ {
    var items []string = strings.Split(lines[i], "")
    if items[j] == "X" {
    p.allBacteria[j].location.coorX = j
    p.allBacteria[j].location.coorX = i-1
  }
}
}
}

func ReadEnergyResourceFile(filename string)  {

}
// Several Circumstances happens with bacteria: increasing/decreasing sunlight,
// increasing/decreasing heat,  add agar, add oxygen, add minerals, add protein source,
// add
func main(){

}
