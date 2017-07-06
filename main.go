package main

import (
	"fmt"
	"./gameobjects"
)

func main()  {
	vasyan1, success := gameobjects.GetAccountCharacter("testTwo", "testPass")
	if success == true {
		item1 := gameobjects.NewItem(1)
		vasyan1.CharInventory["head"] = item1
		fmt.Println(vasyan1.CharInventory["head"].ItemName)
		fmt.Println("--------------------------------")
		fmt.Println(vasyan1.CharacterClassID)
	}
}