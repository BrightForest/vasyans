package main

import "fmt"
import "./gameobjects"

func main()  {
	vasyan1 := gameobjects.NewChar(1)
	fmt.Println(vasyan1.CharName)
}