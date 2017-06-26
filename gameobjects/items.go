package gameobjects

type Item struct {
	ItemInGameID	int
	ItemBaseID	int
	ItemName	string
	ItemTypeID	int
	WearAttribute	string
	ItemLocation	string
	ItemParameters	map[string]float32
}
