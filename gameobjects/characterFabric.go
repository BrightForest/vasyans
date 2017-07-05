package gameobjects

import "fmt"

func GetAccountCharacter(accountLogin string, accountPassword string) (*Character, bool){
	idI, accountpasswordI, charidI := LoginInGame(accountLogin)
	isTruePassword := checkAccountPassword(accountPassword, accountpasswordI)
	if isTruePassword == false {
		fmt.Println("Account password does not match.")
		noCharacter := NewChar(0,0)
		return noCharacter, false
	}
	getCharacter := NewChar(idI, charidI)
	return getCharacter, true
}

func checkAccountPassword(accountPassword string, accountPasswordI string)  bool{
	if accountPassword != accountPasswordI {
		return false
	}
	return true
}