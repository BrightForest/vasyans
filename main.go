package main

import "fmt"
import "./gameobjects"

func main()  {
	vasyan1 := new(gameobjects.Character)
	fmt.Println(vasyan1.CharName)
}