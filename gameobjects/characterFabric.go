package gameobjects

func GetAccountCharacter(accountLogin string, accountPassword string) (*Character, bool){
	idI, accountpasswordI, charidI := LoginInGame(accountLogin)
	isTruePassword := checkAccountPassword(accountPassword, accountpasswordI)
	if idI == 0 {
		log.Info("Account", accountLogin, "not found in database.")
		return noCharacter(), false
	}
	if isTruePassword == false {
		log.Info("Account password does not match for login:", accountLogin)
		return noCharacter(), false
	}
	name, win, classId := getCharacterParamsById(idI)
	getCharacter := makeCharacter(idI, charidI, name, win, classId)
	return getCharacter, true
}

func checkAccountPassword(accountPassword string, accountPasswordI string)  bool{
	if accountPassword != accountPasswordI {
		return false
	}
	return true
}

func noCharacter()  *Character{
	noCharacter := makeCharacter(0,0, "", 0, 0)
	return noCharacter
}

func makeCharacter(Id int, charId int, charName string, winCount int, charClassId int)  *Character{
	charstats := make(map[string]*float32)
	charInventory := make(map[string]*Item)
	getChar := new(Character)
	getChar.Id = Id
	getChar.CharName = charName
	getChar.WinCount = winCount
	getChar.CharacterClassID = charClassId
	getChar.ClassName = "Warrior"
	getChar.AccountID = Id
	getChar.CharStats = charstats
	getChar.CharInventory = charInventory
	return getChar
}

