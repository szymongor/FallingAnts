package main

import "fmt"
import "./model"

func main() {
  model.Init();

  stick := model.NewRandomStick(10)

  ants := stick.Ants();
  for i := 1; i < len(ants); i++ {
		fmt.Println(ants[i])
	}

}
