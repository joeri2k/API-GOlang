package config

import (
	"database/sql"
)

func GetMySQLDB() (db *sql.DB, err error){
	dbhost := "localhost";
	dbuser :="root";
	dbpass := "";
	dbname :="facebookdb";
	db, err = sql.Open(dbhost, dbuser+":"+dbpass+"@/"+dbname)
	return
}