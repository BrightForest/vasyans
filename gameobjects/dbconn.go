package gameobjects

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/op/go-logging"
)

const (
	DATA_SOURCE_LINK	=	"postgres://*"
	PSQLDRIVER			=	"postgres"
	GET_ACCOUNT_DATA	=	"SELECT id, accountpassword, charid FROM public.\"Accounts\" where accountlogin = $1"
	GET_CHARACTER_PARAMS_BY_ID	=	"SELECT name, win, \"classId\" FROM public.\"Characters\" where accountid = $1"
)

var log = logging.MustGetLogger("DBConnection")

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
