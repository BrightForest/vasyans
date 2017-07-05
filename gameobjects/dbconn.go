package gameobjects

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DATA_SOURCE_LINK	=	"postgres://user:pass@host:port/GameServer?sslmode=disable"
	PSQLDRIVER			=	"postgres"
	GET_ACCOUNT_DATA	=	"SELECT id, accountpassword, charid FROM public.\"Accounts\" where accountlogin = $1"
)

func LoginInGame(accountLogin string) {
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
		fmt.Printf("%3v | %8v | %6v\n", id, accountpassword, charid)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
