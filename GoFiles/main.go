package main

import (

)

func main() {
  func main()  {
  //---------- Initialize Petri, Bacteria and food packages--
     var p Petri
     p.InitializeBact(10,100.0) // The two parameters are number of bact and radius of Petri
     p.InitializeFoodpackage(200,100.0) // Two parameters are number of foodpackages and radius of distribution range

  //-------------------------Update Phenotypes-----------------

   //--------------------------Replication---------------------------

   //----------------------------Movement--------------------------
  for i := range p.allBacteria{
   if IsRandomMove(i) == true {
   RandomWalk(stepsize, X, Y, p)
 } else {
   p.MovetoToFood()
 }
}
   //------------------------------Fight-------------------------------

   //--------------------------Metabolism--------------------------

   //----------------------------Mutation-----------------------------

    //--------------------------Drawing-------------------------------

}

}
