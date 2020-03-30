package teachers

import (
	"encoding/json"
	"errors"

	"github.com/avasites/univer/internal/app/models"
)

//entity ...
type entity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

//QueryParametr ...
type QueryParametr struct {
	DB     string
	Select string
	Where  string
}

//ReadMany...
func ReadMany(params QueryParametr) (string, error) {
	db := models.GetConnection()
	defer db.Close()

	params.DB = "entity"

	selectQuery := "SELECT "

	if params.DB == "" {
		return "Error", errors.New("Not passed required parametrs")
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

	rows, err := db.Query(selectQuery)
	if err != nil {
		return "Error", err
	}

	defer rows.Close()

	entities := []entity{}

	for rows.Next() {

		e := entity{}

		err := rows.Scan(&e.ID, &e.Name, &e.Type)

		if err != nil {
			return "Error", err
		}

		entities = append(entities, e)
	}

	content, err := json.Marshal(entities)

	return string(content), err
}
