package apiserver

import (
	"net/http"

	"github.com/avasites/univer/internal/app/crud/entity"
	"github.com/avasites/univer/internal/app/crud/listtable"
	"github.com/avasites/univer/internal/app/models"
)

//searchHandler ...
func searchHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		r.ParseForm()

		nameEntity := r.FormValue("name")

		fileds := r.FormValue("fields")

		res, err := entity.ReadMany(models.QueryParametr{Select: fileds, Where: "name LIKE '%" + nameEntity + "%'"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(res))

	} else {

		http.Error(w, "Not supported method of GET", http.StatusInternalServerError)

	}

}

//listtableHandler...

func listtableHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		r.ParseForm()

		typeEntity := r.FormValue("type")

		id := r.FormValue("id")

		if typeEntity == "" || id == "" {

			notParam := ""

			if typeEntity == "" {
				notParam += "type, "
			}

			if id == "" {
				notParam += "id"
			}

			http.Error(w, "Not passed required parametrs: "+notParam, http.StatusInternalServerError)
			return
		}

		where := typeEntity[:1] + "id = " + id

		res, err := listtable.ReadMany(models.QueryParametr{Select: "", Where: where})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(res))

	} else {

		http.Error(w, "Not supported method of POST", http.StatusInternalServerError)

	}

}

//New ...
func New() {

	http.HandleFunc("/search", searchHandler)

	http.HandleFunc("/listtable", listtableHandler)

	http.ListenAndServe(":9091", nil)
}
