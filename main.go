package main

import "fmt"
import "./gameobjects"

func main()  {
	vasyan1 := gameobjects.NewChar(1)
	item1 := gameobjects.NewItem(1)
	vasyan1.CharInventory["head"] = item1
	fmt.Println(vasyan1.CharInventory["head"].ItemName)
}