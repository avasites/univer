package models

import (
	"database/sql"

	//pq ...
	_ "github.com/lib/pq"
)

//DdUser ...
const (
	DdUser     = "postgres"
	DbPassword = "example"
	DbName     = "test_go"
)

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

	dbinfo := "postgres://" + DdUser + ":" + DbPassword + "@psql/" + DbName + "?sslmode=disable"

	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(err)
	}

	return db
}
