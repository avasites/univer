package entity

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/avasites/univer/internal/app/models"
)

//entity ...
type entity struct {
	ID   int            `json:"id"`
	Name string         `json:"name"`
	Type sql.NullString `json:"type"`
	TID  sql.NullString `json:"tid"`
	GID  sql.NullString `json:"gid"`
}

//ReadMany ...
func ReadMany(params models.QueryParametr) (string, error) {
	db := models.GetConnection()
	defer db.Close()

	params.DB = "entity"

	if params.Select != "" {

		arrSelect := strings.Split(params.Select, ",")

		if len(arrSelect) > 0 {
			countAdd := 5 - len(arrSelect)

			if countAdd > 0 {

				count := len(arrSelect)

				for count < 5 {
					params.Select += ", null"
					count++
				}

			}
		}
	}

	selectQuery := models.QueryGenerator(params)

	rows, err := db.Query(selectQuery)

	if err != nil {
		return "Error", err
	}

	defer rows.Close()
	entities := []entity{}

	for rows.Next() {

		e := entity{}

		err := rows.Scan(&e.ID, &e.Name, &e.Type, &e.TID, &e.GID)

		if err != nil {
			return "Error", err
		}

		entities = append(entities, e)
	}

	content, err := json.Marshal(entities)

	return string(content), err
}
