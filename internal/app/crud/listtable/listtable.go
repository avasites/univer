package listtable

import (
	"encoding/json"

	"github.com/avasites/univer/internal/app/models"
)

//listtable ...
type listtable struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Time    string `json:"time"`
	Week    string `json:"week"`
	Day     string `json:"day"`
	Cabinet string `json:"cabinet"`
	GID     string `json:"gid"`
	TID     string `json:"tid"`
}

//ReadMany...
func ReadMany(params models.QueryParametr) (string, error) {
	db := models.GetConnection()
	defer db.Close()

	params.DB = "listtable"

	selectQuery := models.QueryGenerator(params)

	rows, err := db.Query(selectQuery)
	if err != nil {
		return "Error", err
	}

	defer rows.Close()
	listtables := []listtable{}

	for rows.Next() {

		l := listtable{}

		err := rows.Scan(&l.ID, &l.TID, &l.GID, &l.Name, &l.Time, &l.Cabinet, &l.Day, &l.Week)

		if err != nil {
			return "Error", err
		}

		listtables = append(listtables, l)
	}

	content, err := json.Marshal(listtables)

	return string(content), err
}
