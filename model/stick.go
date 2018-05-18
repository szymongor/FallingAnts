package model

import "math/rand"
import "time"
import "fmt"

type Stick struct {
  ants []Ant
}

func (s *Stick) Ants() []Ant {
  return s.ants
}

func (s *Stick) AntsPostions() []*float32 {
  antsPostions := []*float32{}
  for i := 0; i < len(s.ants); i++ {
    antsPostions = append(antsPostions, &s.ants[i].position)
	}
  return antsPostions
}

func (s *Stick) AntsDirections() []*string {
  antsDirections := []*string{}
  for i := 0; i < len(s.ants); i++ {
    antsDirections = append(antsDirections, &s.ants[i].direction)
	}
  return antsDirections
}

func Init(){
  rand.Seed(time.Now().Unix())
}

func NewRandomStick(antsNumber int) Stick {
  stick := Stick {make([]Ant,0,antsNumber)}
  for i := 0; i < antsNumber; i++ {
		stick.ants = append(stick.ants , NewAnt())
    fmt.Println(stick.ants[i])
	}
  antsPostions := stick.AntsPostions()
  antsDirections := stick.AntsDirections()
  *antsPostions[0]=2
  *antsDirections[0]="r"
  return stick
}
