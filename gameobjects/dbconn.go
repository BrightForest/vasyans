package gameobjects

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	DATA_SOURCE_LINK	=	"postgres://testUser:8IK1q23e@creatorain.com:7100/GameServer?sslmode=disable"
	PSQLDRIVER			=	"postgres"
	GET_ACCOUNT_DATA	=	"SELECT id, accountpassword, charid FROM public.\"Accounts\" where accountlogin = $1"
)

func LoginInGame(accountLogin string) (int, string, int){
	var idI int
	var accountpasswordI string
	var charidI int

	dbconn, err := sql.Open(PSQLDRIVER, DATA_SOURCE_LINK)
	checkErr(err)
	defer dbconn.Close()

	stmt, err:= dbconn.Prepare(GET_ACCOUNT_DATA)
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
