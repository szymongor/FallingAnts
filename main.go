package main

import "fmt"
import "./model"

func main() {
  model.Init();

  stick := model.NewRandomStick(4)
  fmt.Println(stick)
}
