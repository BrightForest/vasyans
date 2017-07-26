package gameobjects

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/op/go-logging"
)

const (
	DATA_SOURCE_LINK	=	"@"
	PSQLDRIVER			=	"postgres"
	GET_ACCOUNT_DATA	=	"SELECT id, accountpassword, charid FROM public.\"Accounts\" where accountlogin = $1"
	GET_CHARACTER_PARAMS_BY_ID	=	"SELECT name, win, \"classId\" FROM public.\"Characters\" where accountid = $1"
	GET_ITEMS_TO_CACHE = "SELECT * FROM public.\"Items\""
)

var log = logging.MustGetLogger("DBConnection")

func LoadItemsFromBase() {

	dbconn, err := sql.Open(PSQLDRIVER, DATA_SOURCE_LINK)
	checkErr(err)
	defer dbconn.Close()

	stmt, err := dbconn.Prepare(GET_ITEMS_TO_CACHE)
	checkErr(err)
	rows, err := stmt.Query();
	checkErr(err)
	for rows.Next(){
		var (itemid int
			itemname string
			itemtypeid int
			strenght float32
			lifepoints float32
			liferegenerationspeed float32
			physicalattackpower float32
			physicalstunchance float32
			fireresist float32
			waterresist float32
			airresist float32
			earthresist float32
			energyresist float32
			darkresist float32
			mentalresist float32
			fireabsorb float32
			waterabsorb float32
			airabsorb float32
			earthabsorb float32
			energyabsorb float32
			darkabsorb float32
			mentalabsorb float32
			blocknegativeeffectchance float32
			dexterity float32
			castspeed float32
			attackspeed float32
			movementspeed float32
			criticalhitchance float32
			criticalhitmultiplier float32
			intelligence float32
			energysheildpoints float32
			energysheildrechargespeed float32
			magicattack float32
			magiceffectchance float32
			magiceffectduration float32
			manapoints float32
			manaregenerationspeed float32
			lifepointspercent float32
			energysheildpercent float32
			physicalresist float32)

		err = rows.Scan(
			&itemid,
			&itemname,
			&itemtypeid,
			&strenght,
			&lifepoints,
			&liferegenerationspeed,
			&physicalattackpower,
			&physicalstunchance,
			&fireresist,
			&waterresist,
			&airresist,
			&earthresist,
			&energyresist,
			&darkresist,
			&mentalresist,
			&fireabsorb,
			&waterabsorb,
			&airabsorb,
			&earthabsorb,
			&energyabsorb,
			&darkabsorb,
			&mentalabsorb,
			&blocknegativeeffectchance,
			&dexterity,
			&castspeed,
			&attackspeed,
			&movementspeed,
			&criticalhitchance,
			&criticalhitmultiplier,
			&intelligence,
			&energysheildpoints,
			&energysheildrechargespeed,
			&magicattack,
			&magiceffectchance,
			&magiceffectduration,
			&manapoints,
			&manaregenerationspeed,
			&lifepointspercent,
			&energysheildpercent,
			&physicalresist)
		checkErr(err)

		getItem := new(Item)
		getItem.ItemInGameID = 0
		getItem.ItemBaseID = itemid
		getItem.ItemName = itemname
		getItem.ItemTypeID = itemtypeid
		getItem.ItemLocation = "cache"
		switch itemtypeid {
		case 1:
			getItem.WearAttribute = "Amulet"
		case 2:
			getItem.WearAttribute = "Body"
		case 3:
			getItem.WearAttribute = "Head"
		case 4:
			getItem.WearAttribute = "Belt"
		case 5:
			getItem.WearAttribute = "Gloves"
		case 6:
			getItem.WearAttribute = "Boots"
		case 7:
			getItem.WearAttribute = "Right Hand"
		case 8:
			getItem.WearAttribute = "Left Hand"
		case 9:
			getItem.WearAttribute = "Right Ring"
		case 10:
			getItem.WearAttribute = "Left Ring"
		}
		getItem.ItemParameters["strenght"] = strenght
		getItem.ItemParameters["lifepoints"] = lifepoints
		getItem.ItemParameters["liferegenerationspeed"] = liferegenerationspeed
		getItem.ItemParameters["physicalattackpower"] = physicalattackpower
		getItem.ItemParameters["physicalstunchance"] = physicalstunchance
		getItem.ItemParameters["fireresist"] = fireresist
		getItem.ItemParameters["waterresist"] = waterresist
		getItem.ItemParameters["airresist"] = airresist
		getItem.ItemParameters["earthresist"] = earthresist
		getItem.ItemParameters["energyresist"] = energyresist
		getItem.ItemParameters["darkresist"] = darkresist
		getItem.ItemParameters["mentalresist"] = mentalresist
		getItem.ItemParameters["fireabsorb"] = fireabsorb
		getItem.ItemParameters["waterabsorb"] = waterabsorb
		getItem.ItemParameters["airabsorb"] = airabsorb
		getItem.ItemParameters["earthabsorb"] = earthabsorb
		getItem.ItemParameters["energyabsorb"] = energyabsorb
		getItem.ItemParameters["darkabsorb"] = darkabsorb
		getItem.ItemParameters["mentalabsorb"] = mentalabsorb
		getItem.ItemParameters["blocknegativeeffectchance"] = blocknegativeeffectchance
		getItem.ItemParameters["dexterity"] = dexterity
		getItem.ItemParameters["castspeed"] = castspeed
		getItem.ItemParameters["attackspeed"] = attackspeed
		getItem.ItemParameters["movementspeed"] = movementspeed
		getItem.ItemParameters["criticalhitchance"] = criticalhitchance
		getItem.ItemParameters["criticalhitmultiplier"] = criticalhitmultiplier
		getItem.ItemParameters["intelligence"] = intelligence
		getItem.ItemParameters["energysheildpoints"] = energysheildpoints
		getItem.ItemParameters["energysheildrechargespeed"] = energysheildrechargespeed
		getItem.ItemParameters["magicattack"] = magicattack
		getItem.ItemParameters["magiceffectchance"] = magiceffectchance
		getItem.ItemParameters["magiceffectduration"] = magiceffectduration
		getItem.ItemParameters["manapoints"] = manapoints
		getItem.ItemParameters["manaregenerationspeed"] = manaregenerationspeed
		getItem.ItemParameters["lifepointspercent"] = lifepointspercent
		getItem.ItemParameters["energysheildpercent"] = energysheildpercent
		getItem.ItemParameters["physicalresist"] = physicalresist

		ItemsCacheMap[itemid] = getItem
	}
}

func LoginInGame(accountLogin string) (int, string, int){
	var idI int
	var accountpasswordI string
	var charidI int

	dbconn, err := sql.Open(PSQLDRIVER, DATA_SOURCE_LINK)
	checkErr(err)
	defer dbconn.Close()

	stmt, err := dbconn.Prepare(GET_ACCOUNT_DATA)
	checkErr(err)
	rows, err := stmt.Query(accountLogin)
	checkErr(err)
	for rows.Next() {
		var id int
		var accountpassword string
		var charid int
		err = rows.Scan(&id, &accountpassword, &charid)
		checkErr(err)
		idI = id
		accountpasswordI = accountpassword
		charidI = charid
	}
	return idI, accountpasswordI, charidI
}

func getCharacterParamsById(Id int)  (string, int, int){
	var nameI string
	var winI int
	var classIdI int

	dbconn, err := sql.Open(PSQLDRIVER, DATA_SOURCE_LINK)
	checkErr(err)
	defer dbconn.Close()

	stmt, err := dbconn.Prepare(GET_CHARACTER_PARAMS_BY_ID)
	checkErr(err)
	rows, err := stmt.Query(Id)
	checkErr(err)
	for rows.Next(){
		var name string
		var win int
		var classId int
		err = rows.Scan(&name, &win, &classId)
		checkErr(err)
		nameI = name
		winI = win
		classIdI = classId
	}
	return nameI, winI, classIdI
}

func checkErr(err error) {
	if err != nil {
		log.Info(err)
	}
}
