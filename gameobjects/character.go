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

func NewChar(Id int)  *Character{
	charstats := make(map[string]*float32)
	charInventory := make(map[string]*Item)
	getChar := new(Character)
	getChar.Id = Id
	getChar.CharName = "Vasya1"
	getChar.WinCount = 0
	getChar.CharacterClassID = 1
	getChar.ClassName = "Warrior"
	getChar.AccountID = 1
	getChar.CharStats = charstats
	getChar.CharInventory = charInventory
	return getChar
}