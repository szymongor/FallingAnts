package model

import "math/rand"
import "time"
// import "fmt"

var DIRECTIONS = [...]string{"l","r"}

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

func NewRandomStick(antsNumber int) Stick {
  stick := Stick {make([]Ant,0,antsNumber)}
  for i := 0; i < antsNumber; i++ {
		stick.ants = append(stick.ants , NewAnt())
	}
  antsPostions := stick.AntsPostions()
  antsDirections := stick.AntsDirections()
  randomAntsPositions(antsPostions)
  randomAntsDirections(antsDirections)
  return stick
}

func randomAntsPositions(antsPostions []*float32) {
  sum := rand.Float32()
  *antsPostions[0] = sum
  for i := 1; i < len(antsPostions); i++ {
    sum += rand.Float32()
		*antsPostions[i] = sum
	}
  sum += rand.Float32()
  for i := 0; i < len(antsPostions); i++ {
		*antsPostions[i] /= sum
	}
}

func randomAntsDirections(antsDirections []*string) {
  for i := 0; i < len(antsDirections); i++ {
		*antsDirections[i] = DIRECTIONS[rand.Intn(len(DIRECTIONS))]
	}
}

func Init(){
  rand.Seed(time.Now().Unix())
}
