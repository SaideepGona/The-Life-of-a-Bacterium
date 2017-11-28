package main
// This file is to test how long will the bacteria exist and draw a graph on
// the relationship between density of foodpackage and the duration of bacteria
func (p *Petri)CountSteps(){
  count := 0
 for p.IsEnd() == false && count < 100 {
     p.UpdateAllEE()
     p.ChecktoDeleteBact()
     p.ChecktoDeleteFood()
   //  p.Replication()
    // p.RandomStep(p.radius,p.radius)
     //p.Attack()
     p.MoveToFood()
}
}
