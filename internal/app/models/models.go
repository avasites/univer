package models

import (
	"database/sql"
	"fmt"

	//pq ...
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

//DbData ...
type DbData struct {
	NameDB   string
	NameUSER string
	Passw    string
}

//QueryParametr ...
type QueryParametr struct {
	DB     string
	Select string
	Where  string
}

//QueryGenerator ...
func QueryGenerator(params QueryParametr) string {

	selectQuery := "SELECT "

	if params.DB == "" {
		return "Not passed required parameteres: Name DB"
	}

	if params.Select != "" {
		selectQuery += params.Select
	} else {
		selectQuery += "*"
	}

	selectQuery += " FROM " + params.DB

	if params.Where != "" {
		selectQuery += " WHERE " + params.Where
	}

	return selectQuery

}

//GetConnection ...
func GetConnection() *sql.DB {

	var config DbData

	if _, err := toml.DecodeFile("/go/src/github.com/avasites/univer/config.toml", &config); err != nil {
		fmt.Println(err)
	}

	dbinfo := "postgres://" + config.NameUSER + ":" + config.Passw + "@psql/" + config.NameDB + "?sslmode=disable"

	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println(err)
	}

	return db
}
