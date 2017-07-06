package gameobjects

type Character struct {
	Id	int
	CharName	string
	WinCount	int
	CharacterClassID	int
	ClassName	string
	AccountID	int
	CharStats	map[string]*float32
	CharInventory	map[string]*Item
}