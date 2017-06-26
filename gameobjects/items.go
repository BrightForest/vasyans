package gameobjects

type Item struct {
	ItemInGameID	int
	ItemBaseID	int
	ItemName	string
	ItemTypeID	int
	WearAttribute	string
	ItemLocation	string
	ItemParameters	map[string]*float32
}

func NewItem(ItemBaseID int) *Item{
	getItem := new(Item)
	getItem.ItemBaseID = ItemBaseID
	getItem.ItemInGameID = 1
	getItem.ItemName = "Test Item"
	getItem.ItemTypeID = 1
	getItem.WearAttribute = "head"
	getItem.ItemLocation = "inventory"
	itemParametersMap := make(map[string]*float32)
	getItem.ItemParameters = itemParametersMap
	return getItem
}