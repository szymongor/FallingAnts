package model


var directions = [...]string{"l","r"}

type Ant struct {
  direction string
  position float32
  velocity float32
}

func (a *Ant) Direction() string {
  return a.direction
}

func (a *Ant) Move(time float32) {
  shift := a.velocity*time
  if(a.direction == "l") {
    shift*=-1
  }
  a.position+=shift
}

func NewAnt() Ant {
  ant := Ant{"l",0,1}
  return ant
}
