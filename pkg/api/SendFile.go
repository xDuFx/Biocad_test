package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



func  (api *api) files(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		pglm := mux.Vars(r)
		pg, ok := pglm["page"]
		if !ok {
			http.Error(w, "No parameter", http.StatusInternalServerError)
			return
		}
		lm, ok :=  pglm["limit"]
		if !ok {
			http.Error(w, "No parameter", http.StatusInternalServerError)
			return
		}
		page, err := strconv.Atoi(pg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		limit, err := strconv.Atoi(lm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := api.db.Pagin(page, limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}