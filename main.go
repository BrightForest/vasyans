package main

import "fmt"
import "./gameobjects"

func main()  {
	vasyan1, success := gameobjects.GetAccountCharacter("testFour", "testPass")
	if success == true {
		item1 := gameobjects.NewItem(1)
		vasyan1.CharInventory["head"] = item1
		fmt.Println(vasyan1.CharInventory["head"].ItemName)
		fmt.Println("--------------------------------")
		fmt.Println(vasyan1.CharacterClassID)
	}
}